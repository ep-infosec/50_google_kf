// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration_test

import (
	"context"
	"fmt"
	"testing"

	logging "cloud.google.com/go/logging"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/kf/v2/pkg/kf/testutil"
	"github.com/google/kf/v2/pkg/kf/testutil/integration"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
)

func TestWriteMetric_Int64(t *testing.T) {
	testutil.WithGomockController(t, func(ctrl *gomock.Controller) {
		fakeLogger := integration.NewFakeMetricLogger(ctrl)
		fakeLogger.EXPECT().
			LogSync(gomock.Any(), logging.Entry{
				Payload: map[string]interface{}{
					"name":  "some_name",
					"value": int64(99),
				},
				Labels: map[string]string{"a": "1", "b": "2"},
				Resource: &mrpb.MonitoredResource{
					Type: "global",
				},
			})
		ctx := integration.ContextWithMetricLogger(context.Background(), fakeLogger)
		integration.WriteMetric(ctx, integration.Int64Gauge{
			Name:   "some_name",
			Value:  99,
			Labels: integration.MetricLabels{"a": "1", "b": "2"},
		})
	})
}

func TestWriteMetric_Float64(t *testing.T) {
	testutil.WithGomockController(t, func(ctrl *gomock.Controller) {
		fakeLogger := integration.NewFakeMetricLogger(ctrl)
		fakeLogger.EXPECT().
			LogSync(gomock.Any(), logging.Entry{
				Payload: map[string]interface{}{
					"name":  "some_name",
					"value": float64(99.9),
				},
				Labels: map[string]string{"a": "1", "b": "2"},
				Resource: &mrpb.MonitoredResource{
					Type: "global",
				},
			})

		ctx := integration.ContextWithMetricLogger(context.Background(), fakeLogger)
		integration.WriteMetric(ctx, integration.Float64Gauge{
			Name:   "some_name",
			Value:  99.9,
			Labels: integration.MetricLabels{"a": "1", "b": "2"},
		})
	})
}

func TestWriteMetric_CombineLabels(t *testing.T) {
	testutil.WithGomockController(t, func(ctrl *gomock.Controller) {
		fakeLogger := integration.NewFakeMetricLogger(ctrl)
		fakeLogger.EXPECT().
			LogSync(gomock.Any(), logging.Entry{
				Payload: map[string]interface{}{
					"name":  "some_name",
					"value": int64(99),
				},
				Labels: map[string]string{"a": "1", "b": "2", "c": "3"},
				Resource: &mrpb.MonitoredResource{
					Type: "global",
				},
			})

		ctx := integration.ContextWithMetricLogger(context.Background(), fakeLogger)
		ctx = integration.ContextWithMetricLabels(ctx, integration.MetricLabels{
			"a": "wrong",
			"c": "3",
		})
		integration.WriteMetric(ctx, integration.Int64Gauge{
			Name:   "some_name",
			Value:  99,
			Labels: integration.MetricLabels{"a": "1", "b": "2"},
		})
	})
}

func TestWriteMetric_NoWriter(t *testing.T) {
	ctx := context.Background()

	integration.WriteMetric(ctx, integration.Int64Gauge{
		Name:   "some_name",
		Value:  99,
		Labels: integration.MetricLabels{"a": "1", "b": "2"},
	})

	// Getting this far indicates success
}

func ExampleContextWithMetricLabels() {
	ctx := integration.ContextWithMetricLabels(context.Background(), integration.MetricLabels{
		"a": "wrong",
		"c": "3",
	})
	ctx = integration.ContextWithMetricLabels(ctx, integration.MetricLabels{
		"a": "1",
		"b": "2",
	})
	m := integration.MetricLabelsFromContext(ctx)
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["c"])

	// Output: 1
	// 2
	// 3

}
