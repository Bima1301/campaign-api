package campaign

import "campaign-api/user"

type GetCampaignDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required,max=255"`
	ShortDescription string `json:"short_description" binding:"required,max=1500"`
	Description      string `json:"description" binding:"required,max=1500"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}
