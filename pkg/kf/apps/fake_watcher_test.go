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
// Source: k8s.io/apimachinery/pkg/watch (interfaces: Interface)

// Package apps_test is a generated GoMock package.
package apps_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	watch "k8s.io/apimachinery/pkg/watch"
)

// FakeWatcher is a mock of Interface interface.
type FakeWatcher struct {
	ctrl     *gomock.Controller
	recorder *FakeWatcherMockRecorder
}

// FakeWatcherMockRecorder is the mock recorder for FakeWatcher.
type FakeWatcherMockRecorder struct {
	mock *FakeWatcher
}

// NewFakeWatcher creates a new mock instance.
func NewFakeWatcher(ctrl *gomock.Controller) *FakeWatcher {
	mock := &FakeWatcher{ctrl: ctrl}
	mock.recorder = &FakeWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeWatcher) EXPECT() *FakeWatcherMockRecorder {
	return m.recorder
}

// ResultChan mocks base method.
func (m *FakeWatcher) ResultChan() <-chan watch.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResultChan")
	ret0, _ := ret[0].(<-chan watch.Event)
	return ret0
}

// ResultChan indicates an expected call of ResultChan.
func (mr *FakeWatcherMockRecorder) ResultChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultChan", reflect.TypeOf((*FakeWatcher)(nil).ResultChan))
}

// Stop mocks base method.
func (m *FakeWatcher) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *FakeWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*FakeWatcher)(nil).Stop))
}
