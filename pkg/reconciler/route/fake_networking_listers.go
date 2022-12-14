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
// Source: github.com/google/kf/v2/pkg/client/networking/listers/networking/v1alpha3 (interfaces: VirtualServiceLister,VirtualServiceNamespaceLister)

// Package route is a generated GoMock package.
package route

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/google/kf/v2/pkg/apis/networking/v1alpha3"
	v1alpha30 "github.com/google/kf/v2/pkg/client/networking/listers/networking/v1alpha3"
	labels "k8s.io/apimachinery/pkg/labels"
)

// FakeVirtualServiceLister is a mock of VirtualServiceLister interface.
type FakeVirtualServiceLister struct {
	ctrl     *gomock.Controller
	recorder *FakeVirtualServiceListerMockRecorder
}

// FakeVirtualServiceListerMockRecorder is the mock recorder for FakeVirtualServiceLister.
type FakeVirtualServiceListerMockRecorder struct {
	mock *FakeVirtualServiceLister
}

// NewFakeVirtualServiceLister creates a new mock instance.
func NewFakeVirtualServiceLister(ctrl *gomock.Controller) *FakeVirtualServiceLister {
	mock := &FakeVirtualServiceLister{ctrl: ctrl}
	mock.recorder = &FakeVirtualServiceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeVirtualServiceLister) EXPECT() *FakeVirtualServiceListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *FakeVirtualServiceLister) List(arg0 labels.Selector) ([]*v1alpha3.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v1alpha3.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *FakeVirtualServiceListerMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeVirtualServiceLister)(nil).List), arg0)
}

// VirtualServices mocks base method.
func (m *FakeVirtualServiceLister) VirtualServices(arg0 string) v1alpha30.VirtualServiceNamespaceLister {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualServices", arg0)
	ret0, _ := ret[0].(v1alpha30.VirtualServiceNamespaceLister)
	return ret0
}

// VirtualServices indicates an expected call of VirtualServices.
func (mr *FakeVirtualServiceListerMockRecorder) VirtualServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualServices", reflect.TypeOf((*FakeVirtualServiceLister)(nil).VirtualServices), arg0)
}

// FakeVirtualServiceNamespaceLister is a mock of VirtualServiceNamespaceLister interface.
type FakeVirtualServiceNamespaceLister struct {
	ctrl     *gomock.Controller
	recorder *FakeVirtualServiceNamespaceListerMockRecorder
}

// FakeVirtualServiceNamespaceListerMockRecorder is the mock recorder for FakeVirtualServiceNamespaceLister.
type FakeVirtualServiceNamespaceListerMockRecorder struct {
	mock *FakeVirtualServiceNamespaceLister
}

// NewFakeVirtualServiceNamespaceLister creates a new mock instance.
func NewFakeVirtualServiceNamespaceLister(ctrl *gomock.Controller) *FakeVirtualServiceNamespaceLister {
	mock := &FakeVirtualServiceNamespaceLister{ctrl: ctrl}
	mock.recorder = &FakeVirtualServiceNamespaceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeVirtualServiceNamespaceLister) EXPECT() *FakeVirtualServiceNamespaceListerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *FakeVirtualServiceNamespaceLister) Get(arg0 string) (*v1alpha3.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v1alpha3.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *FakeVirtualServiceNamespaceListerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*FakeVirtualServiceNamespaceLister)(nil).Get), arg0)
}

// List mocks base method.
func (m *FakeVirtualServiceNamespaceLister) List(arg0 labels.Selector) ([]*v1alpha3.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v1alpha3.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *FakeVirtualServiceNamespaceListerMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeVirtualServiceNamespaceLister)(nil).List), arg0)
}
