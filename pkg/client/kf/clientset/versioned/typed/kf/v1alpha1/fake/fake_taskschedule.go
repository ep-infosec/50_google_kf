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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/google/kf/v2/pkg/apis/kf/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTaskSchedules implements TaskScheduleInterface
type FakeTaskSchedules struct {
	Fake *FakeKfV1alpha1
	ns   string
}

var taskschedulesResource = schema.GroupVersionResource{Group: "kf.dev", Version: "v1alpha1", Resource: "taskschedules"}

var taskschedulesKind = schema.GroupVersionKind{Group: "kf.dev", Version: "v1alpha1", Kind: "TaskSchedule"}

// Get takes name of the taskSchedule, and returns the corresponding taskSchedule object, and an error if there is any.
func (c *FakeTaskSchedules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TaskSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(taskschedulesResource, c.ns, name), &v1alpha1.TaskSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TaskSchedule), err
}

// List takes label and field selectors, and returns the list of TaskSchedules that match those selectors.
func (c *FakeTaskSchedules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TaskScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(taskschedulesResource, taskschedulesKind, c.ns, opts), &v1alpha1.TaskScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TaskScheduleList{ListMeta: obj.(*v1alpha1.TaskScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.TaskScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested taskSchedules.
func (c *FakeTaskSchedules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(taskschedulesResource, c.ns, opts))

}

// Create takes the representation of a taskSchedule and creates it.  Returns the server's representation of the taskSchedule, and an error, if there is any.
func (c *FakeTaskSchedules) Create(ctx context.Context, taskSchedule *v1alpha1.TaskSchedule, opts v1.CreateOptions) (result *v1alpha1.TaskSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(taskschedulesResource, c.ns, taskSchedule), &v1alpha1.TaskSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TaskSchedule), err
}

// Update takes the representation of a taskSchedule and updates it. Returns the server's representation of the taskSchedule, and an error, if there is any.
func (c *FakeTaskSchedules) Update(ctx context.Context, taskSchedule *v1alpha1.TaskSchedule, opts v1.UpdateOptions) (result *v1alpha1.TaskSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(taskschedulesResource, c.ns, taskSchedule), &v1alpha1.TaskSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TaskSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTaskSchedules) UpdateStatus(ctx context.Context, taskSchedule *v1alpha1.TaskSchedule, opts v1.UpdateOptions) (*v1alpha1.TaskSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(taskschedulesResource, "status", c.ns, taskSchedule), &v1alpha1.TaskSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TaskSchedule), err
}

// Delete takes name of the taskSchedule and deletes it. Returns an error if one occurs.
func (c *FakeTaskSchedules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(taskschedulesResource, c.ns, name, opts), &v1alpha1.TaskSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTaskSchedules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(taskschedulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TaskScheduleList{})
	return err
}

// Patch applies the patch and returns the patched taskSchedule.
func (c *FakeTaskSchedules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TaskSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(taskschedulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.TaskSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TaskSchedule), err
}