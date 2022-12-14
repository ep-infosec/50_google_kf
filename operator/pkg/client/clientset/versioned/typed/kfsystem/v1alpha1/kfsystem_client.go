// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kf-operator/pkg/apis/kfsystem/v1alpha1"
	"kf-operator/pkg/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type KfV1alpha1Interface interface {
	RESTClient() rest.Interface
	KfSystemsGetter
}

// KfV1alpha1Client is used to interact with features provided by the kf.dev group.
type KfV1alpha1Client struct {
	restClient rest.Interface
}

func (c *KfV1alpha1Client) KfSystems() KfSystemInterface {
	return newKfSystems(c)
}

// NewForConfig creates a new KfV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*KfV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KfV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new KfV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KfV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KfV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *KfV1alpha1Client {
	return &KfV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KfV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
