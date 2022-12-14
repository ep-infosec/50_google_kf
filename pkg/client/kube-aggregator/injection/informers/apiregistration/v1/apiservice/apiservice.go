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

// Code generated by injection-gen. DO NOT EDIT.

package apiservice

import (
	context "context"

	versioned "github.com/google/kf/v2/pkg/client/kube-aggregator/clientset/versioned"
	v1 "github.com/google/kf/v2/pkg/client/kube-aggregator/informers/externalversions/apiregistration/v1"
	client "github.com/google/kf/v2/pkg/client/kube-aggregator/injection/client"
	factory "github.com/google/kf/v2/pkg/client/kube-aggregator/injection/informers/factory"
	apiregistrationv1 "github.com/google/kf/v2/pkg/client/kube-aggregator/listers/apiregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	apisapiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Apiregistration().V1().APIServices()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

func withDynamicInformer(ctx context.Context) context.Context {
	inf := &wrapper{client: client.Get(ctx), resourceVersion: injection.GetResourceVersion(ctx)}
	return context.WithValue(ctx, Key{}, inf)
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1.APIServiceInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/google/kf/v2/pkg/client/kube-aggregator/informers/externalversions/apiregistration/v1.APIServiceInformer from context.")
	}
	return untyped.(v1.APIServiceInformer)
}

type wrapper struct {
	client versioned.Interface

	resourceVersion string
}

var _ v1.APIServiceInformer = (*wrapper)(nil)
var _ apiregistrationv1.APIServiceLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apisapiregistrationv1.APIService{}, 0, nil)
}

func (w *wrapper) Lister() apiregistrationv1.APIServiceLister {
	return w
}

// SetResourceVersion allows consumers to adjust the minimum resourceVersion
// used by the underlying client.  It is not accessible via the standard
// lister interface, but can be accessed through a user-defined interface and
// an implementation check e.g. rvs, ok := foo.(ResourceVersionSetter)
func (w *wrapper) SetResourceVersion(resourceVersion string) {
	w.resourceVersion = resourceVersion
}

func (w *wrapper) List(selector labels.Selector) (ret []*apisapiregistrationv1.APIService, err error) {
	lo, err := w.client.ApiregistrationV1().APIServices().List(context.TODO(), metav1.ListOptions{
		LabelSelector:   selector.String(),
		ResourceVersion: w.resourceVersion,
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apisapiregistrationv1.APIService, error) {
	return w.client.ApiregistrationV1().APIServices().Get(context.TODO(), name, metav1.GetOptions{
		ResourceVersion: w.resourceVersion,
	})
}
