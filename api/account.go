package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sendlovebox/go-lenco/model"
)

// CreateVirtualAccount makes the request to create a virtual account on Lenco
func (c *Call) CreateVirtualAccount(ctx context.Context, request model.VirtualAccountRequest) (*model.VirtualAccount, error) {
	response := &model.VirtualAccount{}

	err := c.makeRequest(ctx, http.MethodPost, "/virtual-accounts", request, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetVirtualAccountByRef makes the request to get a virtual account by reference
func (c *Call) GetVirtualAccountByRef(ctx context.Context, reference string) (*model.VirtualAccount, error) {
	response := &model.VirtualAccount{}

	path := fmt.Sprintf("/virtual-accounts/%s", reference)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetVirtualAccountByBVN makes the request to get a virtual account by user's BVN
func (c *Call) GetVirtualAccountByBVN(ctx context.Context, bvn string) (*model.VirtualAccount, error) {
	response := &model.VirtualAccount{}

	path := fmt.Sprintf("/virtual-account-by-bvn/%s", bvn)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetVirtualAccountTransaction makes the request to get a virtual account transaction on by its ID
func (c *Call) GetVirtualAccountTransaction(ctx context.Context, txID string) (*model.Transaction, error) {
	response := &model.Transaction{}

	path := fmt.Sprintf("/virtual-accounts/transactions/%s", txID)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetAccunts makes the request to get a virtual account transaction on by its ID
func (c *Call) GetAccounts(ctx context.Context, txID string) (*model.Transaction, error) {
	response := &model.Transaction{}

	//path := fmt.Sprintf("/virtual-accounts/transactions/%s", txID)
	err := c.makeRequest(ctx, http.MethodGet, "/accounts\n", nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}
