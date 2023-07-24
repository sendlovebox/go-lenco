package model

type (
	// VirtualAccountRequest schema
	VirtualAccountRequest struct {
		AccountName      string  `json:"accountName,omitempty"`
		IsStatic         bool    `json:"isStatic,omitempty"`
		TrxReference     string  `json:"transactionReference,omitempty"`
		Amount           float64 `json:"amount,omitempty"`
		MinAmount        float64 `json:"minAmount,omitempty"`
		BVN              string  `json:"bvn,omitempty"`
		CreateNewAccount bool    `json:"createNewAccount,omitempty"`
	}

	// VirtualAccount schema
	VirtualAccount struct {
		ID               string             `mapstructure:"id,omitempty" json:"id,omitempty"`
		AccountReference string             `mapstructure:"accountReference,omitempty" json:"accountReference"`
		BankAccount      BankAccount        `mapstructure:"bankAccount,omitempty" json:"bankAccount"`
		Type             VirtualAccountType `mapstructure:"type,omitempty" json:"type"`
		Status           string             `mapstructure:"status,omitempty" json:"status"`
		CreatedAt        string             `mapstructure:"createdAt,omitempty" json:"createdAt"`
		ExpiresAt        *string            `mapstructure:"expiresAt,omitempty" json:"expiresAt"`
		Currency         string             `mapstructure:"currency,omitempty" json:"currency"`
		Meta             VirtualAccountMeta `mapstructure:"meta,omitempty" json:"meta"`
	}

	// BankAccount schema
	BankAccount struct {
		AccountName   string `mapstructure:"accountName,omitempty" json:"accountName"`
		AccountNumber string `mapstructure:"accountNumber,omitempty" json:"accountNumber"`
		Bank          struct {
			Code string `mapstructure:"code,omitempty" json:"code,omitempty"`
			Name string `mapstructure:"name,omitempty" json:"name,omitempty"`
		} `mapstructure:"bank,omitempty" json:"bank"`
	}

	// VirtualAccountMeta schema
	VirtualAccountMeta struct {
		Amount               string  `mapstructure:"amount,omitempty" json:"amount"`
		MinAmount            *string `mapstructure:"minAmount,omitempty" json:"minAmount"`
		TransactionReference *string `mapstructure:"minAmount,omitempty" json:"transactionReference"`
		Bvn                  *string `mapstructure:"bvn,omitempty" json:"bvn"`
		IsStatic             bool    `mapstructure:"isStatic,omitempty" json:"isStatic"`
	}
)

const (
	// VirtualAccountTypeStatic string representation of static virtual account
	VirtualAccountTypeStatic VirtualAccountType = "Static Virtual Account"
	// DynamicAccountTypeStatic string representation of dynamic virtual account
	DynamicAccountTypeStatic VirtualAccountType = "Dynamic Virtual Account"

	// StatusActive string
	StatusActive = "active"
	// StatusExpired string
	StatusExpired = "expired"
	// StatusBlacklisted string
	StatusBlacklisted = "blacklisted"
)
