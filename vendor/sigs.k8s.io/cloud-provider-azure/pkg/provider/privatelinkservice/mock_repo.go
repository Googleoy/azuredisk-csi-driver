// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_repo.go -package=privatelinkservice -copyright_file ../../../hack/boilerplate/boilerplate.generatego.txt -source=repo.go Repository
//
// Package privatelinkservice is a generated GoMock package.
package privatelinkservice

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	gomock "go.uber.org/mock/gomock"
	cache "sigs.k8s.io/cloud-provider-azure/pkg/cache"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockRepository) CreateOrUpdate(ctx context.Context, resourceGroup string, pls armnetwork.PrivateLinkService) (*armnetwork.PrivateLinkService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroup, pls)
	ret0, _ := ret[0].(*armnetwork.PrivateLinkService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockRepositoryMockRecorder) CreateOrUpdate(ctx, resourceGroup, pls any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockRepository)(nil).CreateOrUpdate), ctx, resourceGroup, pls)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, resourceGroup, plsName, lbFrontendID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroup, plsName, lbFrontendID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, resourceGroup, plsName, lbFrontendID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, resourceGroup, plsName, lbFrontendID)
}

// DeletePEConnection mocks base method.
func (m *MockRepository) DeletePEConnection(ctx context.Context, resourceGroup, plsName, peConnName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePEConnection", ctx, resourceGroup, plsName, peConnName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePEConnection indicates an expected call of DeletePEConnection.
func (mr *MockRepositoryMockRecorder) DeletePEConnection(ctx, resourceGroup, plsName, peConnName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePEConnection", reflect.TypeOf((*MockRepository)(nil).DeletePEConnection), ctx, resourceGroup, plsName, peConnName)
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, resourceGroup, frontendIPConfigID string, crt cache.AzureCacheReadType) (*armnetwork.PrivateLinkService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroup, frontendIPConfigID, crt)
	ret0, _ := ret[0].(*armnetwork.PrivateLinkService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, resourceGroup, frontendIPConfigID, crt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, resourceGroup, frontendIPConfigID, crt)
}

// List mocks base method.
func (m *MockRepository) List(ctx context.Context, resourceGroup string) ([]*armnetwork.PrivateLinkService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroup)
	ret0, _ := ret[0].([]*armnetwork.PrivateLinkService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(ctx, resourceGroup any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), ctx, resourceGroup)
}