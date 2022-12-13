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
// Source: k8s.io/client-go/listers/core/v1 (interfaces: NamespaceLister)

// Package route is a generated GoMock package.
package route

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	labels "k8s.io/apimachinery/pkg/labels"
)

// FakeNamespaceLister is a mock of NamespaceLister interface.
type FakeNamespaceLister struct {
	ctrl     *gomock.Controller
	recorder *FakeNamespaceListerMockRecorder
}

// FakeNamespaceListerMockRecorder is the mock recorder for FakeNamespaceLister.
type FakeNamespaceListerMockRecorder struct {
	mock *FakeNamespaceLister
}

// NewFakeNamespaceLister creates a new mock instance.
func NewFakeNamespaceLister(ctrl *gomock.Controller) *FakeNamespaceLister {
	mock := &FakeNamespaceLister{ctrl: ctrl}
	mock.recorder = &FakeNamespaceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeNamespaceLister) EXPECT() *FakeNamespaceListerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *FakeNamespaceLister) Get(arg0 string) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *FakeNamespaceListerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*FakeNamespaceLister)(nil).Get), arg0)
}

// List mocks base method.
func (m *FakeNamespaceLister) List(arg0 labels.Selector) ([]*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *FakeNamespaceListerMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeNamespaceLister)(nil).List), arg0)
}
