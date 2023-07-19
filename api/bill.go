package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sendlovebox/go-lenco/model"
)

// GetBillVendors makes the request to get bill vendors
func (c *Call) GetBillVendors(ctx context.Context, category string) (*[]model.Vendor, error) {
	response := &[]model.Vendor{}
	path := "/bills/vendors"

	if category != "" {
		path = fmt.Sprintf("%s?categories[]=%s", path, category)
	}

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetBillVendorByID makes the request to get a bill vendor by its ID
func (c *Call) GetBillVendorByID(ctx context.Context, vendorID string) (*model.Vendor, error) {
	response := &model.Vendor{}
	path := fmt.Sprintf("/bills/vendors/%s", vendorID)

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetProducts makes the request to get products
func (c *Call) GetProducts(ctx context.Context, vendorID string, category string) (*[]model.Product, error) {
	response := &[]model.Product{}
	path := "/bills/products"

	if category != "" {
		path = fmt.Sprintf("%s?categories[]=%s", path, category)
	}

	if vendorID != "" {
		if strings.Contains(path, "?") {
			path = fmt.Sprintf("%s&vendorIds[]=%s", path, vendorID)
		} else {
			path = fmt.Sprintf("%s?vendorIds[]=%s", path, vendorID)
		}
	}

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}
