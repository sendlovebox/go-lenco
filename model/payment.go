package model

type (
	// CreateRecipientRequest schema
	CreateRecipientRequest struct {
		AccountNumber string `json:"accountNumber"`
		BankCode      string `json:"bankCode"`
	}

	// RecipientResponse schema
	RecipientResponse struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		Currency    string      `json:"currency"`
		BankAccount BankAccount `json:"bankAccount"`
	}

	// RecipientsResponse schema
	RecipientsResponse struct {
		Recipients []RecipientResponse `json:"recipients"`
	}

	// CreateTransactionRequest schema
	CreateTransactionRequest struct {
		AccountID     string  `json:"accountId"`
		RecipientID   *string `json:"recipientId,omitempty"`
		AccountNumber *string `json:"accountNumber,omitempty"`
		BankCode      *string `json:"bankCode,omitempty"`
		Amount        string  `json:"amount"`
		Narration     string  `json:"narration"`
		Reference     string  `json:"reference"` // Unique client reference. Only -, ., _ and alphanumeric characters allowed.
	}

	// TransactionResponse schema
	TransactionResponse struct {
		ID                   string       `json:"id"`
		Amount               string       `json:"amount"`
		Fee                  string       `json:"fee"`
		Narration            string       `json:"narration"`
		Type                 string       `json:"type"`
		InitiatedAt          interface{}  `json:"initiatedAt"`
		CompletedAt          interface{}  `json:"completedAt"`
		AccountID            string       `json:"accountId"`
		Details              *BankAccount `json:"details,omitempty"`
		Status               string       `json:"status"`
		FailedAt             interface{}  `json:"failedAt"`
		ReasonForFailure     *string      `json:"reasonForFailure,omitempty"`
		ClientReference      *string      `json:"clientReference,omitempty"`
		TransactionReference string       `json:"transactionReference"`
		NipSessionID         interface{}  `json:"nipSessionId"`
	}

	// ResolveAccountResponse schema
	ResolveAccountResponse struct {
		AccountName   string `json:"accountName"`
		AccountNumber string `json:"accountNumber"`
		Bank          struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"bank"`
	}

	// TransactionsResponse schema
	TransactionsResponse struct {
		Transactions []TransactionResponse `json:"transactions"`
	}
)

const (
	// StatusPending when the status is pending
	StatusPending = "pending"
	// StatusFailed when the status is failed
	StatusFailed = "failed"
	// StatusSuccessful when status is successful
	StatusSuccessful = "successful"
	// StatusDeclined when a status request is declined
	StatusDeclined = "declined"
)
