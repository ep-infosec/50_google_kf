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

package apps

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	v1alpha1 "github.com/google/kf/v2/pkg/apis/kf/v1alpha1"
	"github.com/google/kf/v2/pkg/internal/envutil"
	"github.com/google/kf/v2/pkg/kf/apps"
	"github.com/google/kf/v2/pkg/kf/apps/fake"
	"github.com/google/kf/v2/pkg/kf/commands/config"
	"github.com/google/kf/v2/pkg/kf/testutil"
	"github.com/google/kf/v2/pkg/reconciler/app/resources"
)

func TestEnvCommand(t *testing.T) {
	t.Parallel()

	for tn, tc := range map[string]struct {
		Space           string
		Args            []string
		ExpectedStrings []string
		ExpectedErr     error
		Setup           func(t *testing.T, fake *fake.FakeClient, p *config.KfParams)
	}{
		"wrong number of params": {
			Args:        []string{},
			ExpectedErr: errors.New("accepts 1 arg(s), received 0"),
		},
		"listing variables fails": {
			Space:       "default",
			Args:        []string{"app-name"},
			ExpectedErr: errors.New("some-error"),
			Setup: func(t *testing.T, fake *fake.FakeClient, _ *config.KfParams) {
				fake.EXPECT().Get(gomock.Any(), "default", "app-name").Return(nil, errors.New("some-error"))
			},
		},
		"custom namespace": {
			Args:  []string{"app-name"},
			Space: "some-namespace",
			Setup: func(t *testing.T, fake *fake.FakeClient, _ *config.KfParams) {
				fake.EXPECT().Get(gomock.Any(), "some-namespace", "app-name").Return(&v1alpha1.App{}, nil)
			},
		},
		"empty results": {
			Space: "default",
			Args:  []string{"app-name"},
			Setup: func(t *testing.T, fake *fake.FakeClient, _ *config.KfParams) {
				fake.EXPECT().Get(gomock.Any(), "default", "app-name").Return(&v1alpha1.App{}, nil)
			},
			ExpectedErr:     nil, // explicitly expecting no failure with zero length list
			ExpectedStrings: []string{"Space-Provided:", "User-Provided:", "System-Provided:", "[Resolved at runtime]"},
		},
		"with results": {
			Space: "default",
			Args:  []string{"app-name"},
			Setup: func(t *testing.T, fake *fake.FakeClient, p *config.KfParams) {
				out := apps.NewKfApp()
				out.SetEnvVars(envutil.MapToEnvVars(map[string]string{
					"name-1": "value-1",
					"name-2": "value-2",
				}))

				fake.EXPECT().Get(gomock.Any(), "default", "app-name").Return(out.ToApp(), nil)

				p.TargetSpace.Status.RuntimeConfig.Env = envutil.MapToEnvVars(map[string]string{
					"sn-1": "sv-1",
					"sn-2": "sv-2",
				})
			},
			ExpectedStrings: append(
				[]string{
					"name-1", "value-1", "name-2", "value-2", // app provided
					"sn-1", "sv-1", "sn-2", "sv-2", // space provided
				},
				resources.RuntimeEnvVarList(resources.CFRunning).List()..., // system provided
			),
		},
	} {
		t.Run(tn, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			fake := fake.NewFakeClient(ctrl)

			p := &config.KfParams{
				Space:       tc.Space,
				TargetSpace: &v1alpha1.Space{},
			}

			if tc.Setup != nil {
				tc.Setup(t, fake, p)
			}

			buf := new(bytes.Buffer)

			cmd := NewEnvCommand(p, fake)
			cmd.SetOutput(buf)
			cmd.SetArgs(tc.Args)
			_, actualErr := cmd.ExecuteC()
			if tc.ExpectedErr != nil || actualErr != nil {
				testutil.AssertErrorsEqual(t, tc.ExpectedErr, actualErr)
				return
			}

			testutil.AssertContainsAll(t, buf.String(), tc.ExpectedStrings)
			testutil.AssertEqual(t, "SilenceUsage", true, cmd.SilenceUsage)

		})
	}
}
