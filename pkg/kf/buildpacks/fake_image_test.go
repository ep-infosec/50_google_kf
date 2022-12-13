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
// Source: github.com/google/go-containerregistry/pkg/v1 (interfaces: Image)

// Package buildpacks_test is a generated GoMock package.
package buildpacks_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	types "github.com/google/go-containerregistry/pkg/v1/types"
)

// FakeImage is a mock of Image interface.
type FakeImage struct {
	ctrl     *gomock.Controller
	recorder *FakeImageMockRecorder
}

// FakeImageMockRecorder is the mock recorder for FakeImage.
type FakeImageMockRecorder struct {
	mock *FakeImage
}

// NewFakeImage creates a new mock instance.
func NewFakeImage(ctrl *gomock.Controller) *FakeImage {
	mock := &FakeImage{ctrl: ctrl}
	mock.recorder = &FakeImageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FakeImage) EXPECT() *FakeImageMockRecorder {
	return m.recorder
}

// ConfigFile mocks base method.
func (m *FakeImage) ConfigFile() (*v1.ConfigFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFile")
	ret0, _ := ret[0].(*v1.ConfigFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigFile indicates an expected call of ConfigFile.
func (mr *FakeImageMockRecorder) ConfigFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFile", reflect.TypeOf((*FakeImage)(nil).ConfigFile))
}

// ConfigName mocks base method.
func (m *FakeImage) ConfigName() (v1.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigName")
	ret0, _ := ret[0].(v1.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigName indicates an expected call of ConfigName.
func (mr *FakeImageMockRecorder) ConfigName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigName", reflect.TypeOf((*FakeImage)(nil).ConfigName))
}

// Digest mocks base method.
func (m *FakeImage) Digest() (v1.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Digest")
	ret0, _ := ret[0].(v1.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Digest indicates an expected call of Digest.
func (mr *FakeImageMockRecorder) Digest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Digest", reflect.TypeOf((*FakeImage)(nil).Digest))
}

// LayerByDiffID mocks base method.
func (m *FakeImage) LayerByDiffID(arg0 v1.Hash) (v1.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerByDiffID", arg0)
	ret0, _ := ret[0].(v1.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerByDiffID indicates an expected call of LayerByDiffID.
func (mr *FakeImageMockRecorder) LayerByDiffID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerByDiffID", reflect.TypeOf((*FakeImage)(nil).LayerByDiffID), arg0)
}

// LayerByDigest mocks base method.
func (m *FakeImage) LayerByDigest(arg0 v1.Hash) (v1.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerByDigest", arg0)
	ret0, _ := ret[0].(v1.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerByDigest indicates an expected call of LayerByDigest.
func (mr *FakeImageMockRecorder) LayerByDigest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerByDigest", reflect.TypeOf((*FakeImage)(nil).LayerByDigest), arg0)
}

// Layers mocks base method.
func (m *FakeImage) Layers() ([]v1.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Layers")
	ret0, _ := ret[0].([]v1.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Layers indicates an expected call of Layers.
func (mr *FakeImageMockRecorder) Layers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Layers", reflect.TypeOf((*FakeImage)(nil).Layers))
}

// Manifest mocks base method.
func (m *FakeImage) Manifest() (*v1.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*v1.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Manifest indicates an expected call of Manifest.
func (mr *FakeImageMockRecorder) Manifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*FakeImage)(nil).Manifest))
}

// MediaType mocks base method.
func (m *FakeImage) MediaType() (types.MediaType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MediaType")
	ret0, _ := ret[0].(types.MediaType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MediaType indicates an expected call of MediaType.
func (mr *FakeImageMockRecorder) MediaType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MediaType", reflect.TypeOf((*FakeImage)(nil).MediaType))
}

// RawConfigFile mocks base method.
func (m *FakeImage) RawConfigFile() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawConfigFile")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawConfigFile indicates an expected call of RawConfigFile.
func (mr *FakeImageMockRecorder) RawConfigFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawConfigFile", reflect.TypeOf((*FakeImage)(nil).RawConfigFile))
}

// RawManifest mocks base method.
func (m *FakeImage) RawManifest() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawManifest")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawManifest indicates an expected call of RawManifest.
func (mr *FakeImageMockRecorder) RawManifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawManifest", reflect.TypeOf((*FakeImage)(nil).RawManifest))
}

// Size mocks base method.
func (m *FakeImage) Size() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Size indicates an expected call of Size.
func (mr *FakeImageMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*FakeImage)(nil).Size))
}
