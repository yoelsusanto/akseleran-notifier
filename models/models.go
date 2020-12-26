package models

import (
	"encoding/json"
)

type Campaign struct {
	LoanCreditRating         string  `json:"loan_credit_rating"`
	RoundedEffectiveInterest float64 `json:"rounded_effective_interest"`
	InstallmentPaymentFreq   string  `json:"installment_payment_freq"`
	InterestPaymentFreq      string  `json:"interest_payment_freq"`
	InstallmentLength        int     `json:"installment_length"`
	InstallmentLengthUnit    struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"installment_length_unit"`
	FlatInterest          float64 `json:"flat_interest"`
	EffectiveInterest     float64 `json:"effective_interest"`
	HaveCollateral        bool    `json:"have_collateral"`
	CampaignStart         string  `json:"campaign_start"`
	CampaignEnd           string  `json:"campaign_end"`
	CampaignTimeRemaining string  `json:"campaign_time_remaining"`
	ID                    int     `json:"id"`
	UUID                  string  `json:"uuid"`
	CampaignName          string  `json:"campaign_name"`
	CampaignType          string  `json:"campaign_type"`
	MinFunding            float64 `json:"min_funding"`
	MaxFunding            float64 `json:"max_funding"`
	Cover                 string  `json:"cover"`
	Status                string  `json:"status"`
	TotalInvestment       float64 `json:"total_investment"`
	TotalInvestors        int     `json:"total_investors"`
	FundedPercentage      float64 `json:"funded_percentage"`
	Created               string  `json:"created"`
}

func (c Campaign) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

type AkseleranResponse struct {
	Data  []Campaign `json:"data"`
	Total int        `json:"total"`
	Limit string     `json:"limit"`
}
