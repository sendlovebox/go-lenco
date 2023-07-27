package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sendlovebox/go-lenco/model"
)

// GetBillVendors makes the request to get bill vendors
func (c *Call) GetBillVendors(ctx context.Context, category model.BillCategory) ([]*model.Vendor, error) {
	var response []*model.Vendor

	path := "/bills/vendors"

	params := url.Values{}
	params.Set("categories[]", string(category))

	requestPath := fmt.Sprintf("%s?%s", path, params.Encode())

	err := c.makeRequest(ctx, http.MethodGet, requestPath, nil, response)
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

// GetBillProducts makes the request to get products
func (c *Call) GetBillProducts(ctx context.Context, vendorID *string, category *model.BillCategory) ([]*model.Product, error) {
	var response []*model.Product

	path := "/bills/products"
	params := url.Values{}

	if vendorID != nil {
		params.Set("vendorIds[]", *vendorID)
	}

	if category != nil {
		params.Set("categories[]", string(*category))
	}

	requestPath := fmt.Sprintf("%s?%s", path, params.Encode())

	err := c.makeRequest(ctx, http.MethodGet, requestPath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetBillProductByID makes the request to get a product by its ID
func (c *Call) GetBillProductByID(ctx context.Context, productID string) (*model.Product, error) {
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
func (c *Call) GetAllBills(ctx context.Context, request model.GetAllBillsRequest) ([]*model.BillData, error) {

	var response []*model.BillData

	path := "/bills"

	params := url.Values{}

	if request.VendorID != nil {
		params.Set("vendorIds[]", *request.VendorID)
	}

	if request.ProductID != nil {
		params.Set("vendorIds[]", *request.VendorID)
	}

	if request.CustomerID != nil {
		params.Set("customerId", *request.CustomerID)
	}

	if request.Status != nil {
		params.Set("status", *request.Status)
	}

	if request.Category != nil {
		params.Set("categories[]", string(*request.Category))
	}

	if request.Page != nil {
		params.Set("page", *request.Page)
	}

	if request.Start != nil {
		params.Set("start", *request.Start)
	}

	if request.End != nil {
		params.Set("end", *request.End)
	}

	requestPath := fmt.Sprintf("%s?%s", path, params.Encode())

	err := c.makeRequest(ctx, http.MethodGet, requestPath, nil, response)
	if err != nil {
		return nil, err
	}

	return response, err
}
