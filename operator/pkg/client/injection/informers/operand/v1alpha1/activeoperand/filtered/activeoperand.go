// Code generated by injection-gen. DO NOT EDIT.

package filtered

import (
	context "context"
	apisoperandv1alpha1 "kf-operator/pkg/apis/operand/v1alpha1"
	versioned "kf-operator/pkg/client/clientset/versioned"
	v1alpha1 "kf-operator/pkg/client/informers/externalversions/operand/v1alpha1"
	client "kf-operator/pkg/client/injection/client"
	filtered "kf-operator/pkg/client/injection/informers/factory/filtered"
	operandv1alpha1 "kf-operator/pkg/client/listers/operand/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterFilteredInformers(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct {
	Selector string
}

func withInformer(ctx context.Context) (context.Context, []controller.Informer) {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	infs := []controller.Informer{}
	for _, selector := range labelSelectors {
		f := filtered.Get(ctx, selector)
		inf := f.Operand().V1alpha1().ActiveOperands()
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
		infs = append(infs, inf.Informer())
	}
	return ctx, infs
}

func withDynamicInformer(ctx context.Context) context.Context {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	for _, selector := range labelSelectors {
		inf := &wrapper{client: client.Get(ctx), selector: selector}
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
	}
	return ctx
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context, selector string) v1alpha1.ActiveOperandInformer {
	untyped := ctx.Value(Key{Selector: selector})
	if untyped == nil {
		logging.FromContext(ctx).Panicf(
			"Unable to fetch kf-operator/pkg/client/informers/externalversions/operand/v1alpha1.ActiveOperandInformer with selector %s from context.", selector)
	}
	return untyped.(v1alpha1.ActiveOperandInformer)
}

type wrapper struct {
	client versioned.Interface

	namespace string

	selector string
}

var _ v1alpha1.ActiveOperandInformer = (*wrapper)(nil)
var _ operandv1alpha1.ActiveOperandLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apisoperandv1alpha1.ActiveOperand{}, 0, nil)
}

func (w *wrapper) Lister() operandv1alpha1.ActiveOperandLister {
	return w
}

func (w *wrapper) ActiveOperands(namespace string) operandv1alpha1.ActiveOperandNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace, selector: w.selector}
}

func (w *wrapper) List(selector labels.Selector) (ret []*apisoperandv1alpha1.ActiveOperand, err error) {
	reqs, err := labels.ParseToRequirements(w.selector)
	if err != nil {
		return nil, err
	}
	selector = selector.Add(reqs...)
	lo, err := w.client.OperandV1alpha1().ActiveOperands(w.namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: selector.String(),
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apisoperandv1alpha1.ActiveOperand, error) {
	// TODO(mattmoor): Check that the fetched object matches the selector.
	return w.client.OperandV1alpha1().ActiveOperands(w.namespace).Get(context.TODO(), name, v1.GetOptions{
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
}
