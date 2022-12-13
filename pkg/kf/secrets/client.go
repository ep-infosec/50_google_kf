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

package secrets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/kf/v2/pkg/apis/kf/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	cv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"knative.dev/pkg/kmeta"
)

// ClientExtension holds additional functions that should be exposed by client.
type ClientExtension interface {
	CreateParamsSecret(ctx context.Context, owner kmeta.OwnerRefable, name string, params json.RawMessage) (*v1.Secret, error)
	UpdateParamsSecret(ctx context.Context, namespace string, name string, params json.RawMessage) (*v1.Secret, error)
}

// NewClient creates a new service client.
func NewClient(kclient cv1.SecretsGetter) Client {
	return &coreClient{
		kclient: kclient,
	}
}

// CreateParamsSecret creates a secret with parameters for the given object.
// The created secret will have a name matching that generated by
// ParamsSecretName and value matching BuildParamsSecret.
func (core *coreClient) CreateParamsSecret(ctx context.Context, owner kmeta.OwnerRefable, name string, params json.RawMessage) (*v1.Secret, error) {
	return core.Create(ctx, owner.GetObjectMeta().GetNamespace(), BuildParamsSecret(owner, name, params))
}

// UpdateParamsSecret updates a secret with parameters for the given object.
// The new parameters replace the existing parameters in the secret.
func (core *coreClient) UpdateParamsSecret(ctx context.Context, namespace string, name string, params json.RawMessage) (*v1.Secret, error) {
	data, err := CreateJSONPatch(params)
	if err != nil {
		return nil, err
	}
	return core.kclient.Secrets(namespace).Patch(ctx, name, types.JSONPatchType, data, metav1.PatchOptions{})
}

// BuildParamsSecret builds a secret that holds parameters for a service instance
// or binding.
func BuildParamsSecret(owner kmeta.OwnerRefable, name string, params json.RawMessage) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: owner.GetObjectMeta().GetNamespace(),
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(owner),
			},
		},
		Data: map[string][]byte{
			v1alpha1.ServiceInstanceParamsSecretKey: params,
		},
	}
}

// CreateJSONPatch creates a json patch object with the given json raw message.
func CreateJSONPatch(params json.RawMessage) (data []byte, err error) {
	patch := []interface{}{
		map[string]interface{}{
			"op":    "replace",
			"path":  fmt.Sprintf("/data/%s", v1alpha1.ServiceInstanceParamsSecretKey),
			"value": []byte(params), // this conversion is necessary to ensure the value is encoded as base64 like secret wants
		},
	}
	return json.Marshal(patch)
}
