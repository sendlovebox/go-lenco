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
		ID                   string `mapstructure:"id,omitempty" json:"id,omitempty"`
		Name                 string `mapstructure:"name,omitempty" json:"name,omitempty"`
		Vendor               Vendor `mapstructure:"vendor,omitempty" json:"vendor,omitempty"`
		Amount               Amount `mapstructure:"amount,omitempty" json:"amount,omitempty"`
		CustomerIDLabel      string `mapstructure:"customerIdLabel,omitempty" json:"customerIdLabel,omitempty"`
		BillCategory         string `mapstructure:"billCategory,omitempty" json:"billCategory,omitempty"`
		CommissionPercentage string `mapstructure:"commissionPercentage,omitempty" json:"commissionPercentage,omitempty"`
	}

	// Amount schema
	Amount struct {
		Type    AmountType
		Fixed   *string
		Minimum *string
		Maximum *string
	}

	// GetBillVendorsResponse schema

	// VirtualAccountType string representation of virtual account type
	VirtualAccountType string
	AmountType         string
	BillCategory       string
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
