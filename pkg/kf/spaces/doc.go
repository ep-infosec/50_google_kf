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

// Package spaces provides a cf compatible way of managing spaces in the
// cluster. All spaces are denoted by the label `app.kubernetes.io/managed-by=kf`.
//
// Namespaces without that label are ignored.
package spaces

//go:generate go run ../internal/tools/clientgen/genclient.go client.yml
