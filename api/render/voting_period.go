package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
	"github.com/go-openapi/strfmt"
)

// Snapshot renders an app level model to a gennerated OpenAPI model.
func PeriodInfo(p models.PeriodStats) *genModels.PeriodInfo {
	pi := &genModels.PeriodInfo{
		Period: Period(p.PeriodInfo),
		VoteStats: &genModels.VoteStats{
			NumVoters:      p.Bakers,
			NumVotersTotal: p.TotalBakers,
			VotesAvailable: p.TotalRolls,
			VotesCast:      p.Rolls,
		},
	}

	if p.BallotsStat != nil {
		pi.Ballots = &genModels.Ballots{
			Yay:           p.BallotsStat.Yay,
			Nay:           p.BallotsStat.Nay,
			Pass:          p.BallotsStat.Pass,
			Quorum:        p.BallotsStat.Quorum,
			Supermajority: p.BallotsStat.Supermajority,
		}
	}

	if p.Proposal != nil {
		pi.Proposal = &genModels.Proposal{
			Hash:         p.Proposal.Hash,
			ProposalFile: p.Proposal.ProposalFile,
			Proposer: &genModels.ProposalProposer{
				Name: p.Proposal.Name,
				Pkh:  p.Proposal.Pkh,
			},
			ShortDescription: p.Proposal.ShortDescription,
			Title:            p.Proposal.Title,
		}
	}

	return pi
}

func Periods(vp []models.PeriodInfo) []*genModels.Period {
	votes := make([]*genModels.Period, len(vp))
	for i := range vp {
		votes[i] = Period(vp[i])
	}
	return votes
}

func Period(p models.PeriodInfo) *genModels.Period {
	return &genModels.Period{
		ID:         &p.ID,
		PeriodType: p.Type,
		StartLevel: p.StartBlock,
		EndLevel:   p.EndBlock,
		StartTime:  strfmt.DateTime(p.StartTime),
		EndTime:    strfmt.DateTime(p.EndTime),
	}
}
