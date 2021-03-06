package watcher

import (
	"context"
	"encoding/json"
	"github.com/everstake/teztracker/services"
	"github.com/everstake/teztracker/services/watcher/tasks"
	"github.com/everstake/teztracker/ws"
	"github.com/everstake/teztracker/ws/models"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

const eventsChannel = "events"

type Watcher struct {
	ctx    context.Context
	cancel context.CancelFunc

	hub   *ws.Hub
	l     *pq.Listener
	tasks map[models.EventType]tasks.EventExecutor
}

func NewWatcher(connection string, hub *ws.Hub, provider services.Provider) *Watcher {

	ctx, cancel := context.WithCancel(context.Background())

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Errorf("ListenerEventType error: %s", err)
		}
	}

	listener := pq.NewListener(connection, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen(eventsChannel)
	if err != nil {
		panic(err)
	}

	return &Watcher{
		ctx:    ctx,
		cancel: cancel,
		hub:    hub,
		l:      listener,
		tasks: map[models.EventType]tasks.EventExecutor{
			//Todo Add factory
			models.EventTypeBlock:          tasks.NewBlockTask(provider),
			models.EventTypeOperation:      tasks.NewOperationTask(provider),
			models.EventTypeAccountCreated: tasks.NewAccountTask(provider),
		},
	}
}

type DBEvent struct {
	Table  string      //"table": "blocks",
	Action string      //"action": "INSERT",
	Data   interface{} //"data": {}
}

func (w Watcher) Start() {

	for {
		select {
		case <-w.ctx.Done():
			w.l.Close()
			return
		case n := <-w.l.Notify:

			var ev DBEvent

			err := json.Unmarshal([]byte(n.Extra), &ev)
			if err != nil {
				log.Error(err)
			}

			handler, ok := w.tasks[models.EventType(ev.Table)]
			if !ok {
				log.Errorf("Unknown event: %s", ev.Table)
				continue
			}

			channels, wsData, err := handler.GetEventData(ev.Data)
			if err != nil {
				log.Errorf("GetEventData error: %s", err)
				continue
			}

			//Publish to all channels
			for i := range channels {
				w.hub.Broadcast(models.BasicMessage{
					Event: channels[i],
					Data:  wsData,
				})
			}
		}
	}
}
