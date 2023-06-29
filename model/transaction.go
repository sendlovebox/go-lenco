package model

type (
	// Transaction object schema
	Transaction struct {
		ID                   string          `json:"id"`
		Amount               string          `json:"amount"`
		Fee                  string          `json:"fee"`
		Narration            string          `json:"narration"`
		Type                 string          `json:"type"`
		InitiatedAt          *string         `json:"initiatedAt,omitempty"`
		CompletedAt          *string         `json:"completedAt,omitempty"`
		FailedAt             *string         `json:"failedAt,omitempty"`
		AccountID            string          `json:"accountId"`
		Details              *BankAccount    `json:"details,omitempty"`
		VirtualAccount       *VirtualAccount `json:"virtualAccount,omitempty"`
		Status               string          `json:"status"`
		ReasonForFailure     string          `json:"reasonForFailure"`
		ClientReference      string          `json:"clientReference"`
		TransactionReference string          `json:"transactionReference"`
		NipSessionID         *string         `json:"nipSessionId,omitempty"`
	}
)
