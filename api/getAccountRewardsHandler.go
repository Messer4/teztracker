package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAccountRewardsListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account Rewards List request.
func (h *getAccountRewardsListHandler) Handle(params accounts.GetAccountRewardsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountRewardsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountRewardsListBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, accs, err := service.GetAccountRewardsList(params.AccountID, limiter)
	if err != nil {
		logrus.Errorf("failed to get account baking: %s", err.Error())
		return accounts.NewGetAccountRewardsListNotFound()
	}
	return accounts.NewGetAccountRewardsListOK().WithPayload(render.AccountRewardsList(accs)).WithXTotalCount(count)
}

type getAccountDelegatorsByCycleListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account Delegators List request.
func (h *getAccountDelegatorsByCycleListHandler) Handle(params accounts.GetAccountDelegatorsByCycleListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountDelegatorsByCycleListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountDelegatorsByCycleListBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, accs, err := service.GetAccountDelegatorsByCycle(params.AccountID, params.CycleID, limiter)
	if err != nil {
		logrus.Errorf("failed to get account delegators: %s", err.Error())
		return accounts.NewGetAccountDelegatorsByCycleListNotFound()
	}
	return accounts.NewGetAccountDelegatorsByCycleListOK().WithPayload(render.AccountDelegators(accs)).WithXTotalCount(count)
}
