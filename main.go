// Package main is the entry point
package main

import (
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
	//ctx := context.Background()

	/* virtualAcc, err := api.CreateVirtualAccount(ctx, model.VirtualAccountRequest{
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
	logger.Info().Interface("response", virtualAcc).Msg("success") */

	/* virtualAcc, err := api.GetVirtualAccountByRef(ctx, "afede78d-c265-4080-949c-302c5c147002")
	if err != nil {
		logger.Err(err).Msg("failed")
		return
	}
	logger.Info().Interface("response", virtualAcc).Msg("success") */
}
