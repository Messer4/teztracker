package api

import (
	"fmt"
	"github.com/everstake/teztracker/api/render"
	info "github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
	"github.com/sirupsen/logrus"
	"strings"
)

type getChartsInfoHandler struct {
	provider DbProvider
}

func (h *getChartsInfoHandler) Handle(params info.GetChartsInfoParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return info.NewGetChartsInfoBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return info.NewGetChartsInfoBadRequest()
	}

	service := services.New(repos.New(db), net)

	var columns []string

	for key := range params.Columns {
		columns = append(columns, strings.Split(params.Columns[key], ",")...)
	}

	for key := range columns {
		if err := validate.Enum(fmt.Sprintf("%s.%v", "column", key), "query", columns[key], []interface{}{"blocks", "volume", "operations", "avg_block_delay", "fees", "activations", "delegation_volume", "bakers", "blocks_priority", "whale_accounts"}); err != nil {
			return info.NewGetChartsInfoBadRequest()
		}
	}

	chartData, err := service.GetChartsInfo(params.From, params.To, params.Period, params.Columns)
	if err != nil {
		logrus.Errorf("failed to get charts data: %s", err.Error())
	}

	return info.NewGetChartsInfoOK().WithPayload(render.ChartData(chartData))
}

type getBakerChartHandler struct {
	provider DbProvider
}

func (h *getBakerChartHandler) Handle(params info.GetBakerChartInfoParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return info.NewGetBakerChartInfoBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return info.NewGetBakerChartInfoBadRequest()
	}

	limiter := NewLimiter(params.Limit, nil)
	service := services.New(repos.New(db), net)

	data, err := service.GetBakerChartInfo(limiter)
	if err != nil {
		return info.NewGetBakerChartInfoInternalServerError()
	}

	return info.NewGetBakerChartInfoOK().WithPayload(render.BakerChartData(data))
}

type getBlocksPriorityHandler struct {
	provider DbProvider
}

func (h *getBlocksPriorityHandler) Handle(params info.GetBlocksPriorityChartInfoParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return info.NewGetBlocksPriorityChartInfoBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return info.NewGetBlocksPriorityChartInfoBadRequest()
	}

	limiter := NewLimiter(params.Limit, nil)
	service := services.New(repos.New(db), net)

	data, err := service.GetBlocksPriorityByCycle(limiter)
	if err != nil {
		return info.NewGetBlocksPriorityChartInfoInternalServerError()
	}

	return info.NewGetBlocksPriorityChartInfoOK().WithPayload(render.BlocksPriorityChartData(data))
}
