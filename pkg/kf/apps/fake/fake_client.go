// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kf/v2/pkg/kf/apps/fake (interfaces: Client)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/google/kf/v2/pkg/apis/kf/v1alpha1"
	apps "github.com/google/kf/v2/pkg/kf/apps"
)

// FakeClient is a mock of Client interface.
type FakeClient struct {
	ctrl     *gomock.Controller
	recorder *FakeClientMockRecorder
}

// FakeClientMockRecorder is the mock recorder for FakeClient.
type FakeClientMockRecorder struct {
	mock *FakeClient
}

// NewFakeClient creates a new mock instance.
func NewFakeClient(ctrl *gomock.Controller) *FakeClient {
	mock := &FakeClient{ctrl: ctrl}
	mock.recorder = &FakeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeClient) EXPECT() *FakeClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *FakeClient) Create(arg0 context.Context, arg1 string, arg2 *v1alpha1.App) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *FakeClientMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*FakeClient)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *FakeClient) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *FakeClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*FakeClient)(nil).Delete), arg0, arg1, arg2)
}

// DeployLogs mocks base method.
func (m *FakeClient) DeployLogs(arg0 context.Context, arg1 io.Writer, arg2, arg3, arg4 string, arg5 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployLogs", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployLogs indicates an expected call of DeployLogs.
func (mr *FakeClientMockRecorder) DeployLogs(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployLogs", reflect.TypeOf((*FakeClient)(nil).DeployLogs), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DeployLogsForApp mocks base method.
func (m *FakeClient) DeployLogsForApp(arg0 context.Context, arg1 io.Writer, arg2 *v1alpha1.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployLogsForApp", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployLogsForApp indicates an expected call of DeployLogsForApp.
func (mr *FakeClientMockRecorder) DeployLogsForApp(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployLogsForApp", reflect.TypeOf((*FakeClient)(nil).DeployLogsForApp), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *FakeClient) Get(arg0 context.Context, arg1, arg2 string) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *FakeClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*FakeClient)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *FakeClient) List(arg0 context.Context, arg1 string) ([]v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *FakeClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeClient)(nil).List), arg0, arg1)
}

// Restage mocks base method.
func (m *FakeClient) Restage(arg0 context.Context, arg1, arg2 string) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Restage indicates an expected call of Restage.
func (mr *FakeClientMockRecorder) Restage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restage", reflect.TypeOf((*FakeClient)(nil).Restage), arg0, arg1, arg2)
}

// Restart mocks base method.
func (m *FakeClient) Restart(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restart", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restart indicates an expected call of Restart.
func (mr *FakeClientMockRecorder) Restart(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*FakeClient)(nil).Restart), arg0, arg1, arg2)
}

// Transform mocks base method.
func (m *FakeClient) Transform(arg0 context.Context, arg1, arg2 string, arg3 apps.Mutator) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transform indicates an expected call of Transform.
func (mr *FakeClientMockRecorder) Transform(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*FakeClient)(nil).Transform), arg0, arg1, arg2, arg3)
}

// Upsert mocks base method.
func (m *FakeClient) Upsert(arg0 context.Context, arg1 string, arg2 *v1alpha1.App, arg3 apps.Merger) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *FakeClientMockRecorder) Upsert(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*FakeClient)(nil).Upsert), arg0, arg1, arg2, arg3)
}

// WaitFor mocks base method.
func (m *FakeClient) WaitFor(arg0 context.Context, arg1, arg2 string, arg3 time.Duration, arg4 apps.Predicate) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitFor", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitFor indicates an expected call of WaitFor.
func (mr *FakeClientMockRecorder) WaitFor(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitFor", reflect.TypeOf((*FakeClient)(nil).WaitFor), arg0, arg1, arg2, arg3, arg4)
}

// WaitForConditionKnativeServiceReadyTrue mocks base method.
func (m *FakeClient) WaitForConditionKnativeServiceReadyTrue(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForConditionKnativeServiceReadyTrue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForConditionKnativeServiceReadyTrue indicates an expected call of WaitForConditionKnativeServiceReadyTrue.
func (mr *FakeClientMockRecorder) WaitForConditionKnativeServiceReadyTrue(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConditionKnativeServiceReadyTrue", reflect.TypeOf((*FakeClient)(nil).WaitForConditionKnativeServiceReadyTrue), arg0, arg1, arg2, arg3)
}

// WaitForConditionReadyTrue mocks base method.
func (m *FakeClient) WaitForConditionReadyTrue(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForConditionReadyTrue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForConditionReadyTrue indicates an expected call of WaitForConditionReadyTrue.
func (mr *FakeClientMockRecorder) WaitForConditionReadyTrue(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConditionReadyTrue", reflect.TypeOf((*FakeClient)(nil).WaitForConditionReadyTrue), arg0, arg1, arg2, arg3)
}

// WaitForConditionRoutesReadyTrue mocks base method.
func (m *FakeClient) WaitForConditionRoutesReadyTrue(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForConditionRoutesReadyTrue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForConditionRoutesReadyTrue indicates an expected call of WaitForConditionRoutesReadyTrue.
func (mr *FakeClientMockRecorder) WaitForConditionRoutesReadyTrue(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConditionRoutesReadyTrue", reflect.TypeOf((*FakeClient)(nil).WaitForConditionRoutesReadyTrue), arg0, arg1, arg2, arg3)
}

// WaitForConditionServiceInstanceBindingsReadyTrue mocks base method.
func (m *FakeClient) WaitForConditionServiceInstanceBindingsReadyTrue(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForConditionServiceInstanceBindingsReadyTrue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForConditionServiceInstanceBindingsReadyTrue indicates an expected call of WaitForConditionServiceInstanceBindingsReadyTrue.
func (mr *FakeClientMockRecorder) WaitForConditionServiceInstanceBindingsReadyTrue(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConditionServiceInstanceBindingsReadyTrue", reflect.TypeOf((*FakeClient)(nil).WaitForConditionServiceInstanceBindingsReadyTrue), arg0, arg1, arg2, arg3)
}

// WaitForDeletion mocks base method.
func (m *FakeClient) WaitForDeletion(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) (*v1alpha1.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDeletion", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForDeletion indicates an expected call of WaitForDeletion.
func (mr *FakeClientMockRecorder) WaitForDeletion(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDeletion", reflect.TypeOf((*FakeClient)(nil).WaitForDeletion), arg0, arg1, arg2, arg3)
}