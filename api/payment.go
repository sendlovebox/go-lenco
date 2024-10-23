package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sendlovebox/go-lenco/model"
)

// CreateRecipient makes the request to create a recipient on Lenco
func (c *Call) CreateRecipient(ctx context.Context, request model.CreateRecipientRequest) (*model.RecipientResponse, error) {
	response := &model.RecipientResponse{}

	err := c.makeRequest(ctx, http.MethodPost, "/recipients", request, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetRecipient makes the request to get a recipient on Lenco
func (c *Call) GetRecipient(ctx context.Context, id string) (*model.RecipientResponse, error) {
	response := &model.RecipientResponse{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/recipient/%s", id), nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetRecipients makes the request to get recipients on Lenco
func (c *Call) GetRecipients(ctx context.Context) (*model.RecipientsResponse, error) {
	response := &model.RecipientsResponse{}

	err := c.makeRequest(ctx, http.MethodGet, "/recipients", nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// CreateTransaction makes the request to create a transaction on Lenco
func (c *Call) CreateTransaction(ctx context.Context, request model.CreateTransactionRequest) (*model.TransactionResponse, error) {
	response := &model.TransactionResponse{}

	err := c.makeRequest(ctx, http.MethodPost, "/transactions", request, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetTransactionByID makes the request to get a transaction on Lenco
func (c *Call) GetTransactionByID(ctx context.Context, id string) (*model.TransactionResponse, error) {
	response := &model.TransactionResponse{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/transaction/%s", id), nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetTransactionByReference makes the request to get a transaction by reference on Lenco
func (c *Call) GetTransactionByReference(ctx context.Context, reference string) (*model.TransactionResponse, error) {
	response := &model.TransactionResponse{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/transaction-by-reference/%s", reference), nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetTransactions makes the request to get transactions on Lenco
func (c *Call) GetTransactions(ctx context.Context) (*model.TransactionsResponse, error) {
	response := &model.TransactionsResponse{}

	err := c.makeRequest(ctx, http.MethodGet, "/transactions", nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ResolveAccount makes the request to resolve account
func (c *Call) ResolveAccount(ctx context.Context, bankCode, accountNumber string) (*model.ResolveAccountResponse, error) {
	response := &model.ResolveAccountResponse{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/resolve?accountNumber=%s&bankCode=%s", accountNumber, bankCode), nil, response)
	if err != nil {
		return nil, err
	}
	return response, err
}
