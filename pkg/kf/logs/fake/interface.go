// Copyright 2019 Google LLC
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

// This file was generated with option-builder.go, DO NOT EDIT IT.

package fake

import (
	"github.com/google/kf/v2/pkg/kf/logs"
)

//go:generate mockgen --package=fake --copyright_file ../../internal/tools/option-builder/LICENSE_HEADER --destination=fake_logs.go --mock_names=Tailer=FakeTailer github.com/google/kf/v2/pkg/kf/logs/fake Tailer

// Tailer is implemented by logs.Tailer.
type Tailer interface {
	logs.Tailer
}
