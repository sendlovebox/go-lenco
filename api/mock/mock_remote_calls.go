// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sendlovebox/go-lenco/model"
)

// MockRemoteCalls is a mock of RemoteCalls interface.
type MockRemoteCalls struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteCallsMockRecorder
}

// MockRemoteCallsMockRecorder is the mock recorder for MockRemoteCalls.
type MockRemoteCallsMockRecorder struct {
	mock *MockRemoteCalls
}

// NewMockRemoteCalls creates a new mock instance.
func NewMockRemoteCalls(ctrl *gomock.Controller) *MockRemoteCalls {
	mock := &MockRemoteCalls{ctrl: ctrl}
	mock.recorder = &MockRemoteCallsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteCalls) EXPECT() *MockRemoteCallsMockRecorder {
	return m.recorder
}

// CreateBill mocks base method.
func (m *MockRemoteCalls) CreateBill(ctx context.Context, request model.CreateBillRequest) (*model.BillData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBill", ctx, request)
	ret0, _ := ret[0].(*model.BillData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBill indicates an expected call of CreateBill.
func (mr *MockRemoteCallsMockRecorder) CreateBill(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBill", reflect.TypeOf((*MockRemoteCalls)(nil).CreateBill), ctx, request)
}

// CreateRecipient mocks base method.
func (m *MockRemoteCalls) CreateRecipient(ctx context.Context, request model.CreateRecipientRequest) (*model.RecipientResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecipient", ctx, request)
	ret0, _ := ret[0].(*model.RecipientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecipient indicates an expected call of CreateRecipient.
func (mr *MockRemoteCallsMockRecorder) CreateRecipient(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecipient", reflect.TypeOf((*MockRemoteCalls)(nil).CreateRecipient), ctx, request)
}

// CreateTransaction mocks base method.
func (m *MockRemoteCalls) CreateTransaction(ctx context.Context, request model.CreateTransactionRequest) (*model.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, request)
	ret0, _ := ret[0].(*model.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockRemoteCallsMockRecorder) CreateTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockRemoteCalls)(nil).CreateTransaction), ctx, request)
}

// CreateVirtualAccount mocks base method.
func (m *MockRemoteCalls) CreateVirtualAccount(ctx context.Context, request model.VirtualAccountRequest) (*model.VirtualAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualAccount", ctx, request)
	ret0, _ := ret[0].(*model.VirtualAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVirtualAccount indicates an expected call of CreateVirtualAccount.
func (mr *MockRemoteCallsMockRecorder) CreateVirtualAccount(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualAccount", reflect.TypeOf((*MockRemoteCalls)(nil).CreateVirtualAccount), ctx, request)
}

// GetAllBills mocks base method.
func (m *MockRemoteCalls) GetAllBills(ctx context.Context, request model.GetAllBillsRequest) ([]model.BillData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBills", ctx, request)
	ret0, _ := ret[0].([]model.BillData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBills indicates an expected call of GetAllBills.
func (mr *MockRemoteCallsMockRecorder) GetAllBills(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBills", reflect.TypeOf((*MockRemoteCalls)(nil).GetAllBills), ctx, request)
}

// GetBillByID mocks base method.
func (m *MockRemoteCalls) GetBillByID(ctx context.Context, id string) (*model.BillData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillByID", ctx, id)
	ret0, _ := ret[0].(*model.BillData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillByID indicates an expected call of GetBillByID.
func (mr *MockRemoteCallsMockRecorder) GetBillByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetBillByID), ctx, id)
}

// GetBillByReference mocks base method.
func (m *MockRemoteCalls) GetBillByReference(ctx context.Context, reference string) (*model.BillData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillByReference", ctx, reference)
	ret0, _ := ret[0].(*model.BillData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillByReference indicates an expected call of GetBillByReference.
func (mr *MockRemoteCallsMockRecorder) GetBillByReference(ctx, reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillByReference", reflect.TypeOf((*MockRemoteCalls)(nil).GetBillByReference), ctx, reference)
}

// GetBillProductByID mocks base method.
func (m *MockRemoteCalls) GetBillProductByID(ctx context.Context, productID string) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillProductByID", ctx, productID)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillProductByID indicates an expected call of GetBillProductByID.
func (mr *MockRemoteCallsMockRecorder) GetBillProductByID(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillProductByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetBillProductByID), ctx, productID)
}

// GetBillProducts mocks base method.
func (m *MockRemoteCalls) GetBillProducts(ctx context.Context, vendorID *string, category *model.BillCategory) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillProducts", ctx, vendorID, category)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillProducts indicates an expected call of GetBillProducts.
func (mr *MockRemoteCallsMockRecorder) GetBillProducts(ctx, vendorID, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillProducts", reflect.TypeOf((*MockRemoteCalls)(nil).GetBillProducts), ctx, vendorID, category)
}

// GetBillVendorByID mocks base method.
func (m *MockRemoteCalls) GetBillVendorByID(ctx context.Context, vendorID string) (*model.Vendor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillVendorByID", ctx, vendorID)
	ret0, _ := ret[0].(*model.Vendor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillVendorByID indicates an expected call of GetBillVendorByID.
func (mr *MockRemoteCallsMockRecorder) GetBillVendorByID(ctx, vendorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillVendorByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetBillVendorByID), ctx, vendorID)
}

// GetBillVendors mocks base method.
func (m *MockRemoteCalls) GetBillVendors(ctx context.Context, category model.BillCategory) ([]model.Vendor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillVendors", ctx, category)
	ret0, _ := ret[0].([]model.Vendor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillVendors indicates an expected call of GetBillVendors.
func (mr *MockRemoteCallsMockRecorder) GetBillVendors(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillVendors", reflect.TypeOf((*MockRemoteCalls)(nil).GetBillVendors), ctx, category)
}

// GetRecipient mocks base method.
func (m *MockRemoteCalls) GetRecipient(ctx context.Context, id string) (*model.RecipientResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipient", ctx, id)
	ret0, _ := ret[0].(*model.RecipientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipient indicates an expected call of GetRecipient.
func (mr *MockRemoteCallsMockRecorder) GetRecipient(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipient", reflect.TypeOf((*MockRemoteCalls)(nil).GetRecipient), ctx, id)
}

// GetRecipients mocks base method.
func (m *MockRemoteCalls) GetRecipients(ctx context.Context) (*model.RecipientsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipients", ctx)
	ret0, _ := ret[0].(*model.RecipientsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipients indicates an expected call of GetRecipients.
func (mr *MockRemoteCallsMockRecorder) GetRecipients(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipients", reflect.TypeOf((*MockRemoteCalls)(nil).GetRecipients), ctx)
}

// GetTransactionByID mocks base method.
func (m *MockRemoteCalls) GetTransactionByID(ctx context.Context, id string) (*model.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByID", ctx, id)
	ret0, _ := ret[0].(*model.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByID indicates an expected call of GetTransactionByID.
func (mr *MockRemoteCallsMockRecorder) GetTransactionByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetTransactionByID), ctx, id)
}

// GetTransactionByReference mocks base method.
func (m *MockRemoteCalls) GetTransactionByReference(ctx context.Context, reference string) (*model.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByReference", ctx, reference)
	ret0, _ := ret[0].(*model.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByReference indicates an expected call of GetTransactionByReference.
func (mr *MockRemoteCallsMockRecorder) GetTransactionByReference(ctx, reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByReference", reflect.TypeOf((*MockRemoteCalls)(nil).GetTransactionByReference), ctx, reference)
}

// GetTransactions mocks base method.
func (m *MockRemoteCalls) GetTransactions(ctx context.Context) (*model.TransactionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", ctx)
	ret0, _ := ret[0].(*model.TransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockRemoteCallsMockRecorder) GetTransactions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockRemoteCalls)(nil).GetTransactions), ctx)
}

// GetVirtualAccountByBVN mocks base method.
func (m *MockRemoteCalls) GetVirtualAccountByBVN(ctx context.Context, bvn string) (*model.VirtualAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualAccountByBVN", ctx, bvn)
	ret0, _ := ret[0].(*model.VirtualAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualAccountByBVN indicates an expected call of GetVirtualAccountByBVN.
func (mr *MockRemoteCallsMockRecorder) GetVirtualAccountByBVN(ctx, bvn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualAccountByBVN", reflect.TypeOf((*MockRemoteCalls)(nil).GetVirtualAccountByBVN), ctx, bvn)
}

// GetVirtualAccountByRef mocks base method.
func (m *MockRemoteCalls) GetVirtualAccountByRef(ctx context.Context, reference string) (*model.VirtualAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualAccountByRef", ctx, reference)
	ret0, _ := ret[0].(*model.VirtualAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualAccountByRef indicates an expected call of GetVirtualAccountByRef.
func (mr *MockRemoteCallsMockRecorder) GetVirtualAccountByRef(ctx, reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualAccountByRef", reflect.TypeOf((*MockRemoteCalls)(nil).GetVirtualAccountByRef), ctx, reference)
}

// GetVirtualAccountTransaction mocks base method.
func (m *MockRemoteCalls) GetVirtualAccountTransaction(ctx context.Context, txID string) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualAccountTransaction", ctx, txID)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualAccountTransaction indicates an expected call of GetVirtualAccountTransaction.
func (mr *MockRemoteCallsMockRecorder) GetVirtualAccountTransaction(ctx, txID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualAccountTransaction", reflect.TypeOf((*MockRemoteCalls)(nil).GetVirtualAccountTransaction), ctx, txID)
}

// RunCustomerLookup mocks base method.
func (m *MockRemoteCalls) RunCustomerLookup(ctx context.Context, customerID, productID string) (*model.CustomerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCustomerLookup", ctx, customerID, productID)
	ret0, _ := ret[0].(*model.CustomerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCustomerLookup indicates an expected call of RunCustomerLookup.
func (mr *MockRemoteCallsMockRecorder) RunCustomerLookup(ctx, customerID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCustomerLookup", reflect.TypeOf((*MockRemoteCalls)(nil).RunCustomerLookup), ctx, customerID, productID)
}

// RunInSandboxMode mocks base method.
func (m *MockRemoteCalls) RunInSandboxMode() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunInSandboxMode")
}

// RunInSandboxMode indicates an expected call of RunInSandboxMode.
func (mr *MockRemoteCallsMockRecorder) RunInSandboxMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInSandboxMode", reflect.TypeOf((*MockRemoteCalls)(nil).RunInSandboxMode))
}
