package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sendlovebox/go-lenco/model"
)

// GetVendors makes the request to get bill vendors
func (c *Call) GetVendors(ctx context.Context, category string) (*[]model.Vendor, error) {
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

// GetVendorByID makes the request to get a bill vendor by its ID
func (c *Call) GetVendorByID(ctx context.Context, vendorID string) (*model.Vendor, error) {
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

// GetProductByID makes the request to get a product by its ID
func (c *Call) GetProductByID(ctx context.Context, productID string) (*model.Product, error) {
	response := &model.Product{}
	path := fmt.Sprintf("/bills/products/%s", productID)

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// RunCustomerLookup makes the request to run customer look up for a bill
func (c *Call) RunCustomerLookup(ctx context.Context, customerID string, productID string) (*model.Product, error) {
	response := &model.Product{}

	path := fmt.Sprintf("/bills/lookup-account?customerId=%s&productId=%s", customerID, productID)

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// CreateBill makes the request to create a bill
func (c *Call) CreateBill(ctx context.Context, request model.CreateBillRequest) (*model.CreateBillResponse, error) {

	response := &model.CreateBillResponse{}

	err := c.makeRequest(ctx, http.MethodPost, "/bills", request, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetBillByID makes the request to get a specific bill
func (c *Call) GetBillByID(ctx context.Context, id string) (*model.BillData, error) {

	response := &model.BillData{}

	path := fmt.Sprintf("/bills/%s", id)

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetBillByReference makes the request to get a specific bill
func (c *Call) GetBillByReference(ctx context.Context, reference string) (*model.BillData, error) {

	response := &model.BillData{}

	path := fmt.Sprintf("/bills/by-reference/%s", reference)

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetAllBills makes the request to get all bills
func (c *Call) GetAllBills(ctx context.Context, request model.GetAllBillsRequest) (*[]model.BillData, error) {

	response := &[]model.BillData{}

	path := "/bills"

	if string(request.Category) != "" {
		path = fmt.Sprintf("%s?categories[]=%s", path, request.Category)
	}

	if request.VendorID != "" {
		if strings.Contains(path, "?") {
			path = fmt.Sprintf("%s&vendorIds[]=%s", path, request.VendorID)
		} else {
			path = fmt.Sprintf("%s?vendorIds[]=%s", path, request.VendorID)
		}
	}

	if request.CustomerID != "" {
		if strings.Contains(path, "?") {
			path = fmt.Sprintf("%s&customerId=%s", path, request.CustomerID)
		} else {
			path = fmt.Sprintf("%s?customerId=%s", path, request.CustomerID)
		}
	}

	if request.ProductID != "" {
		if strings.Contains(path, "?") {
			path = fmt.Sprintf("%s&productId=%s", path, request.ProductID)
		} else {
			path = fmt.Sprintf("%s?productId=%s", path, request.ProductID)
		}
	}

	if request.Status != "" {
		if strings.Contains(path, "?") {
			path = fmt.Sprintf("%s&status=%s", path, request.Status)
		} else {
			path = fmt.Sprintf("%s?status=%s", path, request.Status)
		}
	}

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}
