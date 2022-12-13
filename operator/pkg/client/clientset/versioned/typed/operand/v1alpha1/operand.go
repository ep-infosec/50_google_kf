// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "kf-operator/pkg/apis/operand/v1alpha1"
	scheme "kf-operator/pkg/client/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperandsGetter has a method to return a OperandInterface.
// A group's client should implement this interface.
type OperandsGetter interface {
	Operands() OperandInterface
}

// OperandInterface has methods to work with Operand resources.
type OperandInterface interface {
	Create(ctx context.Context, operand *v1alpha1.Operand, opts v1.CreateOptions) (*v1alpha1.Operand, error)
	Update(ctx context.Context, operand *v1alpha1.Operand, opts v1.UpdateOptions) (*v1alpha1.Operand, error)
	UpdateStatus(ctx context.Context, operand *v1alpha1.Operand, opts v1.UpdateOptions) (*v1alpha1.Operand, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Operand, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OperandList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Operand, err error)
	OperandExpansion
}

// operands implements OperandInterface
type operands struct {
	client rest.Interface
}

// newOperands returns a Operands
func newOperands(c *OperandV1alpha1Client) *operands {
	return &operands{
		client: c.RESTClient(),
	}
}

// Get takes name of the operand, and returns the corresponding operand object, and an error if there is any.
func (c *operands) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Operand, err error) {
	result = &v1alpha1.Operand{}
	err = c.client.Get().
		Resource("operands").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Operands that match those selectors.
func (c *operands) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OperandList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OperandList{}
	err = c.client.Get().
		Resource("operands").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operands.
func (c *operands) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("operands").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a operand and creates it.  Returns the server's representation of the operand, and an error, if there is any.
func (c *operands) Create(ctx context.Context, operand *v1alpha1.Operand, opts v1.CreateOptions) (result *v1alpha1.Operand, err error) {
	result = &v1alpha1.Operand{}
	err = c.client.Post().
		Resource("operands").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operand).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a operand and updates it. Returns the server's representation of the operand, and an error, if there is any.
func (c *operands) Update(ctx context.Context, operand *v1alpha1.Operand, opts v1.UpdateOptions) (result *v1alpha1.Operand, err error) {
	result = &v1alpha1.Operand{}
	err = c.client.Put().
		Resource("operands").
		Name(operand.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operand).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *operands) UpdateStatus(ctx context.Context, operand *v1alpha1.Operand, opts v1.UpdateOptions) (result *v1alpha1.Operand, err error) {
	result = &v1alpha1.Operand{}
	err = c.client.Put().
		Resource("operands").
		Name(operand.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operand).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the operand and deletes it. Returns an error if one occurs.
func (c *operands) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("operands").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operands) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("operands").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched operand.
func (c *operands) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Operand, err error) {
	result = &v1alpha1.Operand{}
	err = c.client.Patch(pt).
		Resource("operands").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
