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
// Source: knative.dev/pkg/controller (interfaces: Reconciler)

// Package route is a generated GoMock package.
package route

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// FakeReconciler is a mock of Reconciler interface.
type FakeReconciler struct {
	ctrl     *gomock.Controller
	recorder *FakeReconcilerMockRecorder
}

// FakeReconcilerMockRecorder is the mock recorder for FakeReconciler.
type FakeReconcilerMockRecorder struct {
	mock *FakeReconciler
}

// NewFakeReconciler creates a new mock instance.
func NewFakeReconciler(ctrl *gomock.Controller) *FakeReconciler {
	mock := &FakeReconciler{ctrl: ctrl}
	mock.recorder = &FakeReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeReconciler) EXPECT() *FakeReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method.
func (m *FakeReconciler) Reconcile(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile.
func (mr *FakeReconcilerMockRecorder) Reconcile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*FakeReconciler)(nil).Reconcile), arg0, arg1)
}
