// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/sendlovebox/go-lenco/model"
)

// RemoteCalls abstracted definition of supported functions
//
//go:generate mockgen -source api.go -destination ./mock/mock_remote_calls.go -package mock RemoteCalls
type RemoteCalls interface {
	RunInSandboxMode()

	CreateVirtualAccount(ctx context.Context, request model.VirtualAccountRequest) (*model.VirtualAccount, error)
	GetVirtualAccountByRef(ctx context.Context, reference string) (*model.VirtualAccount, error)
	GetVirtualAccountByBVN(ctx context.Context, bvn string) (*model.VirtualAccount, error)
	GetVirtualAccountTransaction(ctx context.Context, txID string) (*model.Transaction, error)
}

// Call object
type Call struct {
	client      *resty.Client
	logger      zerolog.Logger
	sandboxMode bool
	baseURL     string
}

// New initialises the object Call
func New(z *zerolog.Logger, c *resty.Client, baseURL, token string) RemoteCalls {
	c.SetTimeout(10 * time.Second)
	call := &Call{
		client:  c,
		logger:  z.With().Str("sdk", "lenco").Logger(),
		baseURL: baseURL,
	}
	c.SetAuthToken(token).
		SetHeader("Content-type", "application/json").
		SetHeader("Accept", "application/json")

	return RemoteCalls(call)
}

// RunInSandboxMode this forces Call functionalities to run in sandbox mode for relevant logic/API consumption
func (c *Call) RunInSandboxMode() {
	c.client.SetDebug(true)
	c.client.EnableTrace()
	c.sandboxMode = true
}
