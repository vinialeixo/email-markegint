package campaign

import (
	"github.com/vinialeixo/email-markegint/internal/dto"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign dto.NewCampaignRequest) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)

	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}
