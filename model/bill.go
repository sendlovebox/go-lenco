package model

type (
	// Vendor schema
	Vendor struct {
		ID       string    `mapstructure:"id,omitempty" json:"id,omitempty"`
		Name     string    `mapstructure:"name,omitempty" json:"name,omitempty"`
		Products []Product `mapstructure:"products,omitempty" json:"products,omitempty"`
	}

	// Product schema
	Product struct {
		ID                   string       `mapstructure:"id,omitempty" json:"id,omitempty"`
		Name                 string       `mapstructure:"name,omitempty" json:"name,omitempty"`
		Vendor               Vendor       `mapstructure:"vendor,omitempty" json:"vendor,omitempty"`
		Amount               Amount       `mapstructure:"amount,omitempty" json:"amount,omitempty"`
		CustomerIDLabel      string       `mapstructure:"customerIdLabel,omitempty" json:"customerIdLabel,omitempty"`
		BillCategory         BillCategory `mapstructure:"billCategory,omitempty" json:"billCategory,omitempty"`
		CommissionPercentage string       `mapstructure:"commissionPercentage,omitempty" json:"commissionPercentage,omitempty"`
	}

	// Amount schema
	Amount struct {
		Type    AmountType
		Fixed   *string
		Minimum *string
		Maximum *string
	}

	// CustomerDetails schema
	CustomerDetails struct {
		CustomerID   string `mapstructure:"customerId,omitempty" json:"customerId,omitempty"`
		CustomerName string `mapstructure:"customerName,omitempty" json:"customerName,omitempty"`
	}

	// CustomerLookUpResponse schema
	CustomerLookUpResponse struct {
		Status  bool            `mapstructure:"status,omitempty" json:"status,omitempty"`
		Message string          `mapstructure:"message,omitempty" json:"message,omitempty"`
		Data    CustomerDetails `mapstructure:"data,omitempty" json:"data,omitempty"`
	}

	// CreateBillRequest payload schema for making a bill request
	CreateBillRequest struct {
		CustomerID     string `json:"customerId,omitempty"`
		DebitAccountID string `json:"debitAccountId,omitempty"`
		ProductID      string `json:"productId,omitempty"`
		Amount         string `json:"amount,omitempty"`
		Reference      string `json:"reference,omitempty"`
	}

	// GetAllBillsRequest schema
	GetAllBillsRequest struct {
		CustomerID string       `url:"customerId,omitempty"`
		ProductID  string       `url:"productId,omitempty"`
		VendorID   string       `url:"vendorIds[],omitempty"`
		Status     string       `url:"status,omitempty"`
		Category   BillCategory `url:"categories[],omitempty"`
		Start      string       `url:"start,omitempty"`
		End        string       `url:"end,omitempty"`
		Page       string       `url:"page,omitempty"`
	}

	// CreateBillResponse schema
	CreateBillResponse struct {
		Status  interface{} `mapstructure:"status,omitempty" json:"status,omitempty"`
		Message string      `mapstructure:"message,omitempty" json:"message,omitempty"`
		Data    BillData    `mapstructure:"data,omitempty" json:"data,omitempty"`
	}

	// BillData schema
	BillData struct {
		ID                   string          `mapstructure:"id,omitempty" json:"id,omitempty"`
		Amount               string          `mapstructure:"amount,omitempty" json:"amount,omitempty"`
		Vendor               Vendor          `mapstructure:"vendor,omitempty" json:"vendor,omitempty"`
		Category             BillCategory    `mapstructure:"category,omitempty" json:"category,omitempty"`
		Product              Product         `mapstructure:"product,omitempty" json:"product,omitempty"`
		Details              CustomerDetails `mapstructure:"details,omitempty" json:"details,omitempty"`
		DebitAccountID       string          `mapstructure:"debitAccountId,omitempty" json:"debitAccountId,omitempty"`
		Instructions         *string         `mapstructure:"instructions,omitempty" json:"instructions,omitempty"`
		InitiatedAt          string          `mapstructure:"initiatedAt,omitempty" json:"initiatedAt,omitempty"`
		Status               interface{}     `mapstructure:"status,omitempty" json:"status,omitempty"`
		CompletedAt          *string         `mapstructure:"completedAt,omitempty" json:"completedAt,omitempty"`
		FailedAt             *string         `mapstructure:"failedAt,omitempty" json:"failedAt,omitempty"`
		ReasonForFailure     *string         `mapstructure:"reasonForFailure,omitempty" json:"reasonForFailure,omitempty"`
		ClientReference      *string         `mapstructure:"clientReference,omitempty" json:"clientReference,omitempty"`
		TransactionReference string          `mapstructure:"transactionReference,omitempty" json:"transactionReference,omitempty"`
		Commission           *string         `mapstructure:"commission,omitempty" json:"commission,omitempty"`
	}

	// VirtualAccountType string representation of virtual account type
	VirtualAccountType string
	// AmountType string representation of amount type
	AmountType string
	// BillCategory string representation of bill category
	BillCategory string
)

const (
	// AmountTypeFixed string representation of fixed amount type
	AmountTypeFixed AmountType = "fixed"
	// AmountTypeRange string representation of range amount type
	AmountTypeRange AmountType = "range"
	// BillCategoryAirtime string representation of airtime bill category
	BillCategoryAirtime BillCategory = "airtime"
	// BillCategoryMobileData string representation of mobile-data bill category
	BillCategoryMobileData BillCategory = "mobile-data"
	// BillCategoryCableTV string representation of cable-tv bill category
	BillCategoryCableTV BillCategory = "cable-tv"
	// BillCategoryElectricity string representation of electricity bill category
	BillCategoryElectricity BillCategory = "electricity"
)
