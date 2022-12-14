// Code generated by injection-gen. DO NOT EDIT.

package kfsystem

import (
	context "context"
	apiskfsystemv1alpha1 "kf-operator/pkg/apis/kfsystem/v1alpha1"
	versioned "kf-operator/pkg/client/clientset/versioned"
	v1alpha1 "kf-operator/pkg/client/informers/externalversions/kfsystem/v1alpha1"
	client "kf-operator/pkg/client/injection/client"
	factory "kf-operator/pkg/client/injection/informers/factory"
	kfsystemv1alpha1 "kf-operator/pkg/client/listers/kfsystem/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
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
	inf := f.Kf().V1alpha1().KfSystems()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

func withDynamicInformer(ctx context.Context) context.Context {
	inf := &wrapper{client: client.Get(ctx), resourceVersion: injection.GetResourceVersion(ctx)}
	return context.WithValue(ctx, Key{}, inf)
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.KfSystemInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch kf-operator/pkg/client/informers/externalversions/kfsystem/v1alpha1.KfSystemInformer from context.")
	}
	return untyped.(v1alpha1.KfSystemInformer)
}

type wrapper struct {
	client versioned.Interface

	resourceVersion string
}

var _ v1alpha1.KfSystemInformer = (*wrapper)(nil)
var _ kfsystemv1alpha1.KfSystemLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apiskfsystemv1alpha1.KfSystem{}, 0, nil)
}

func (w *wrapper) Lister() kfsystemv1alpha1.KfSystemLister {
	return w
}

// SetResourceVersion allows consumers to adjust the minimum resourceVersion
// used by the underlying client.  It is not accessible via the standard
// lister interface, but can be accessed through a user-defined interface and
// an implementation check e.g. rvs, ok := foo.(ResourceVersionSetter)
func (w *wrapper) SetResourceVersion(resourceVersion string) {
	w.resourceVersion = resourceVersion
}

func (w *wrapper) List(selector labels.Selector) (ret []*apiskfsystemv1alpha1.KfSystem, err error) {
	lo, err := w.client.KfV1alpha1().KfSystems().List(context.TODO(), v1.ListOptions{
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

func (w *wrapper) Get(name string) (*apiskfsystemv1alpha1.KfSystem, error) {
	return w.client.KfV1alpha1().KfSystems().Get(context.TODO(), name, v1.GetOptions{
		ResourceVersion: w.resourceVersion,
	})
}
