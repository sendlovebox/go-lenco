package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/mitchellh/mapstructure"
	"github.com/sendlovebox/go-lenco/model"
)

// MakeRequest is a super function that handles all requests to the lenco SDK
func (c *Call) makeRequest(ctx context.Context, method, path string, body, expectedRes interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")
	log.Info().Interface(model.LogStrRequest, body).Msg("request")
	defer log.Info().Msg("done...")

	var (
		err  error
		res  *resty.Response
		resp = &model.Response{}
	)

	client := c.client.R().
		SetContext(ctx).
		SetResult(&resp).
		SetError(&resp)

	if body != nil {
		client.SetBody(body)
	}

	switch method {
	case http.MethodGet:
		res, err = client.Get(endpoint)
	case http.MethodPost:
		res, err = client.Post(endpoint)
	case http.MethodPut:
		res, err = client.Put(endpoint)
	case http.MethodPatch:
		res, err = client.Patch(endpoint)
	default:
		err = errors.New("invalid method")
		log.Err(err).Str("method", method).Msg("invalid method passed")
		return err
	}

	if err != nil {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return err
	}

	if !resp.Status {
		err = errors.New(resp.Message)
		log.Err(err).Msg("error while making request")
		return err
	}

	log.Info().Interface(model.LogStrResponse, resp).Msg("response")

	return mapstruct(resp.Data, expectedRes)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}
