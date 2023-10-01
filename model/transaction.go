package model

type (
	// Transaction object schema
	Transaction struct {
		ID                   string         `json:"id" mapstructure:""`
		TransactionAmount    string         `json:"transactionAmount" mapstructure:""`
		Fee                  string         `json:"fee" mapstructure:""`
		StampDuty            string         `json:"stampDuty" mapstructure:""`
		SettledAmount        string         `json:"settledAmount" mapstructure:""`
		Currency             string         `json:"currency" mapstructure:""`
		Type                 string         `json:"type" mapstructure:""`
		Status               string         `json:"status" mapstructure:""`
		Narration            string         `json:"narration" mapstructure:""`
		Details              BankAccount    `json:"details" mapstructure:""`
		VirtualAccount       VirtualAccount `json:"virtualAccount" mapstructure:""`
		AccountReference     string         `json:"accountReference" mapstructure:""`
		SettledAccountID     string         `json:"settledAccountId" mapstructure:""`
		Datetime             string         `json:"datetime" mapstructure:""`
		NipSessionID         string         `json:"nipSessionId" mapstructure:""`
		TransactionReference string         `json:"transactionReference" mapstructure:""`
	}
)
