// Code generated by MockGen. DO NOT EDIT.
// Source: device/spec/spec.go
//
// Generated by this command:
//
//	mockgen -source=device/spec/spec.go -destination=../../internal/agent/device/spec/mock_spec.go -package=spec
//

// Package spec is a generated GoMock package.
package spec

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/flightctl/flightctl/api/v1alpha1"
	client "github.com/flightctl/flightctl/internal/agent/client"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CheckOsReconciliation mocks base method.
func (m *MockManager) CheckOsReconciliation(ctx context.Context) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOsReconciliation", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckOsReconciliation indicates an expected call of CheckOsReconciliation.
func (mr *MockManagerMockRecorder) CheckOsReconciliation(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOsReconciliation", reflect.TypeOf((*MockManager)(nil).CheckOsReconciliation), ctx)
}

// ClearRollback mocks base method.
func (m *MockManager) ClearRollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearRollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearRollback indicates an expected call of ClearRollback.
func (mr *MockManagerMockRecorder) ClearRollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearRollback", reflect.TypeOf((*MockManager)(nil).ClearRollback))
}

// CreateRollback mocks base method.
func (m *MockManager) CreateRollback(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRollback", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRollback indicates an expected call of CreateRollback.
func (mr *MockManagerMockRecorder) CreateRollback(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRollback", reflect.TypeOf((*MockManager)(nil).CreateRollback), ctx)
}

// Ensure mocks base method.
func (m *MockManager) Ensure() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ensure")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ensure indicates an expected call of Ensure.
func (mr *MockManagerMockRecorder) Ensure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ensure", reflect.TypeOf((*MockManager)(nil).Ensure))
}

// GetDesired mocks base method.
func (m *MockManager) GetDesired(ctx context.Context) (*v1alpha1.RenderedDeviceSpec, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDesired", ctx)
	ret0, _ := ret[0].(*v1alpha1.RenderedDeviceSpec)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDesired indicates an expected call of GetDesired.
func (mr *MockManagerMockRecorder) GetDesired(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDesired", reflect.TypeOf((*MockManager)(nil).GetDesired), ctx)
}

// Initialize mocks base method.
func (m *MockManager) Initialize() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockManagerMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockManager)(nil).Initialize))
}

// IsOSUpdate mocks base method.
func (m *MockManager) IsOSUpdate() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOSUpdate")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOSUpdate indicates an expected call of IsOSUpdate.
func (mr *MockManagerMockRecorder) IsOSUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOSUpdate", reflect.TypeOf((*MockManager)(nil).IsOSUpdate))
}

// IsRollingBack mocks base method.
func (m *MockManager) IsRollingBack(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRollingBack", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRollingBack indicates an expected call of IsRollingBack.
func (mr *MockManagerMockRecorder) IsRollingBack(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRollingBack", reflect.TypeOf((*MockManager)(nil).IsRollingBack), ctx)
}

// IsUpgrading mocks base method.
func (m *MockManager) IsUpgrading() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpgrading")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUpgrading indicates an expected call of IsUpgrading.
func (mr *MockManagerMockRecorder) IsUpgrading() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpgrading", reflect.TypeOf((*MockManager)(nil).IsUpgrading))
}

// OSVersion mocks base method.
func (m *MockManager) OSVersion(specType Type) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OSVersion", specType)
	ret0, _ := ret[0].(string)
	return ret0
}

// OSVersion indicates an expected call of OSVersion.
func (mr *MockManagerMockRecorder) OSVersion(specType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OSVersion", reflect.TypeOf((*MockManager)(nil).OSVersion), specType)
}

// Read mocks base method.
func (m *MockManager) Read(specType Type) (*v1alpha1.RenderedDeviceSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", specType)
	ret0, _ := ret[0].(*v1alpha1.RenderedDeviceSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockManagerMockRecorder) Read(specType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockManager)(nil).Read), specType)
}

// RenderedVersion mocks base method.
func (m *MockManager) RenderedVersion(specType Type) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderedVersion", specType)
	ret0, _ := ret[0].(string)
	return ret0
}

// RenderedVersion indicates an expected call of RenderedVersion.
func (mr *MockManagerMockRecorder) RenderedVersion(specType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderedVersion", reflect.TypeOf((*MockManager)(nil).RenderedVersion), specType)
}

// Rollback mocks base method.
func (m *MockManager) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockManagerMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockManager)(nil).Rollback))
}

// SetClient mocks base method.
func (m *MockManager) SetClient(arg0 client.Management) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", arg0)
}

// SetClient indicates an expected call of SetClient.
func (mr *MockManagerMockRecorder) SetClient(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockManager)(nil).SetClient), arg0)
}

// SetUpgradeFailed mocks base method.
func (m *MockManager) SetUpgradeFailed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUpgradeFailed")
}

// SetUpgradeFailed indicates an expected call of SetUpgradeFailed.
func (mr *MockManagerMockRecorder) SetUpgradeFailed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpgradeFailed", reflect.TypeOf((*MockManager)(nil).SetUpgradeFailed))
}

// Status mocks base method.
func (m *MockManager) Status(arg0 context.Context, arg1 *v1alpha1.DeviceStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockManagerMockRecorder) Status(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockManager)(nil).Status), arg0, arg1)
}

// Upgrade mocks base method.
func (m *MockManager) Upgrade(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockManagerMockRecorder) Upgrade(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockManager)(nil).Upgrade), ctx)
}

// MockPriorityQueue is a mock of PriorityQueue interface.
type MockPriorityQueue struct {
	ctrl     *gomock.Controller
	recorder *MockPriorityQueueMockRecorder
}

// MockPriorityQueueMockRecorder is the mock recorder for MockPriorityQueue.
type MockPriorityQueueMockRecorder struct {
	mock *MockPriorityQueue
}

// NewMockPriorityQueue creates a new mock instance.
func NewMockPriorityQueue(ctrl *gomock.Controller) *MockPriorityQueue {
	mock := &MockPriorityQueue{ctrl: ctrl}
	mock.recorder = &MockPriorityQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriorityQueue) EXPECT() *MockPriorityQueueMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPriorityQueue) Add(item *Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPriorityQueueMockRecorder) Add(item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPriorityQueue)(nil).Add), item)
}

// Clear mocks base method.
func (m *MockPriorityQueue) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockPriorityQueueMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockPriorityQueue)(nil).Clear))
}

// IsEmpty mocks base method.
func (m *MockPriorityQueue) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty.
func (mr *MockPriorityQueueMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockPriorityQueue)(nil).IsEmpty))
}

// IsVersionFailed mocks base method.
func (m *MockPriorityQueue) IsVersionFailed(version string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVersionFailed", version)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVersionFailed indicates an expected call of IsVersionFailed.
func (mr *MockPriorityQueueMockRecorder) IsVersionFailed(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVersionFailed", reflect.TypeOf((*MockPriorityQueue)(nil).IsVersionFailed), version)
}

// Next mocks base method.
func (m *MockPriorityQueue) Next() (*Item, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*Item)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockPriorityQueueMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockPriorityQueue)(nil).Next))
}

// Remove mocks base method.
func (m *MockPriorityQueue) Remove(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", version)
}

// Remove indicates an expected call of Remove.
func (mr *MockPriorityQueueMockRecorder) Remove(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPriorityQueue)(nil).Remove), version)
}

// SetVersionFailed mocks base method.
func (m *MockPriorityQueue) SetVersionFailed(version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVersionFailed", version)
}

// SetVersionFailed indicates an expected call of SetVersionFailed.
func (mr *MockPriorityQueueMockRecorder) SetVersionFailed(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersionFailed", reflect.TypeOf((*MockPriorityQueue)(nil).SetVersionFailed), version)
}

// Size mocks base method.
func (m *MockPriorityQueue) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockPriorityQueueMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockPriorityQueue)(nil).Size))
}
