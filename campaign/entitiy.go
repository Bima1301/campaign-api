package campaign

import (
	"campaign-api/user"
	"time"
)

type Campaign struct {
	ID               string
	UserID           string
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time

	CampaignImages []CampaignImage
	User           user.User
}

type CampaignImage struct {
	ID         string
	CampaignID string
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
