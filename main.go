// Package main is the entry point
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	apiCalls "github.com/sendlovebox/go-lenco/api"
	"github.com/sendlovebox/go-lenco/model"
)

func main() {
	fmt.Println("Running lenco SDK")
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := resty.New()
	logger.Info().Msg("app starting")
	defer logger.Info().Msg("stopped")
	api := apiCalls.New(&logger, client, model.BaseURL, model.APIKey)
	api.RunInSandboxMode() // to ensure it is running in sandbox mode
	ctx := context.Background()

	/*
		// ---------------CreateVirtualAccount---------------

		virtualAcc, err := api.CreateVirtualAccount(ctx, model.VirtualAccountRequest{
			AccountName: "The only Ninja",
			IsStatic:    true,
			//TrxReference:     "",
			//Amount:           0,
			//MinAmount:        0,
			BVN:              "123",
			CreateNewAccount: true,
		})
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", virtualAcc).Msg("success")
	*/

	/*
		// ---------------GetVirtualAccountByRef---------------

		virtualAcc, err := api.GetVirtualAccountByRef(ctx, "afede78d-c265-4080-949c-302c5c147002")
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", virtualAcc).Msg("success")
	*/

	/*
		// ---------------GetBillVendors---------------

		vendors, err := api.GetBillVendors(ctx, string(model.BillCategoryElectricity))
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", vendors).Msg("success")
	*/

	/*
		// ---------------GetBillVendorByID---------------
		vendor, err := api.GetBillVendorByID(ctx, "f46c9b9f-ddc1-4466-9cf2-dda59d18b84f")
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", vendor).Msg("success")

	*/

	/*
		// ---------------GetBillProducts---------------
		products, err := api.GetBillProducts(ctx, "", "")
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", products).Msg("success")

	*/

	/*

			// ---------------GetBillProductByID---------------
		product, err := api.GetBillProductByID(ctx, "c96432f5-c5b7-4eb4-a42d-9b873f8f3ff1")
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", product).Msg("success")

	*/

	/*


			// ---------------RunCustomerLookup---------------
		product, err := api.RunCustomerLookup(ctx, "10433749628", "96c548e4-02cb-4d25-a646-8d6489374a1e")
		if err != nil {
			logger.Err(err).Msg("failed")
			return
		}
		logger.Info().Interface("response", product).Msg("success")

	*/

	/*

			// ---------------CreateBill---------------
		//bill, err := api.CreateBill(ctx, model.CreateBillRequest{
		//	CustomerID:     "10433749628",
		//	DebitAccountID: "431827e5-8f9f-4340-ad37-d07185403b8b",
		//	ProductID:      "96c548e4-02cb-4d25-a646-8d6489374a1e",
		//	Amount:         "100",
		//	Reference:      "testReference",
		//})
		//if err != nil {
		//	logger.Err(err).Msg("failed")
		//	return
		//}
		//logger.Info().Interface("response", bill).Msg("success")

	*/

	// ---------------GetBillByID---------------

	//bill, err := api.GetBillByID(ctx, "99267671-6377-427d-9142-698663a6e77b")
	//if err != nil {
	//	logger.Err(err).Msg("failed")
	//	return
	//}
	//logger.Info().Interface("response", bill).Msg("success")

	// ---------------GetBillByReference---------------
	//bill, err := api.GetBillByReference(ctx, "testReference")
	//if err != nil {
	//	logger.Err(err).Msg("failed")
	//	return
	//}
	//logger.Info().Interface("response", bill).Msg("success")

	// ---------------GetAllBills---------------
	//YYYY-MM-DD
	bills, err := api.GetAllBills(ctx, model.GetAllBillsRequest{Status: "successful", VendorID: "bfa5206c-6122-448b-adfd-b8b764bfa13f", Category: "cable-tv", Start: "2021-01-01", End: "2023-07-23"})

	if err != nil {
		logger.Err(err).Msg("failed")
		return
	}
	logger.Info().Interface("response", bills).Msg("success")

}
