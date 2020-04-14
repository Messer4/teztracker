package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) GetAccountRewardsList(accountID string, limits Limiter) (count int64, rewards []models.AccountReward, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, rewards, err
	}

	count, rewards, err = t.repoProvider.GetAccount().RewardsList(accountID, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	var blockReward, endorsementReward int64
	for i := range rewards {
		//Use FutureBakingRewards for future cycles
		if rewards[i].FutureBakingRewards > 0 {
			rewards[i].BakingRewards = rewards[i].FutureBakingRewards
		}

		rewards[i].Status = getRewardStatus(rewards[i].Cycle, lastBlock.MetaCycle)

		blockReward = BlockReward
		endorsementReward = EndorsementReward

		if rewards[i].Cycle < CarthageCycle {
			blockReward = BabylonBlockReward
			endorsementReward = BabylonEndorsementRewards
		}

		rewards[i].Losses += rewards[i].MissedBaking*blockReward + rewards[i].MissedEndorsements*endorsementReward
	}

	return count, rewards, nil
}

func getRewardStatus(cycle, currentCycle int64) (status models.RewardStatus) {
	switch {
	case cycle > currentCycle:
		status = models.StatusPending
	case cycle == currentCycle:
		status = models.StatusActive
	case cycle >= currentCycle-PreservedCycles:
		status = models.StatusFrozen
	default:
		status = models.StatusUnfrozen
	}

	return status
}