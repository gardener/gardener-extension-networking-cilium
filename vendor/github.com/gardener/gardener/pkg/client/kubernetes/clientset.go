// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package kubernetes

import (
	gardencoreclientset "github.com/gardener/gardener/pkg/client/core/clientset/versioned"

	apiextensionclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	apiregistrationclientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Applier returns the applier of this Clientset.
func (c *Clientset) Applier() ApplierInterface {
	return c.applier
}

// RESTConfig will return the config attribute of the Client object.
func (c *Clientset) RESTConfig() *rest.Config {
	return c.config
}

// Client returns the client of this Clientset.
func (c *Clientset) Client() client.Client {
	return c.client
}

// RESTMapper returns the restMapper of this Clientset.
func (c *Clientset) RESTMapper() meta.RESTMapper {
	return c.restMapper
}

// Kubernetes will return the kubernetes attribute of the Client object.
func (c *Clientset) Kubernetes() kubernetes.Interface {
	return c.kubernetes
}

// GardenCore will return the gardenCore attribute of the Client object.
func (c *Clientset) GardenCore() gardencoreclientset.Interface {
	return c.gardenCore
}

// APIExtension will return the apiextensionsClientset attribute of the Client object.
func (c *Clientset) APIExtension() apiextensionclientset.Interface {
	return c.apiextension
}

// APIRegistration will return the apiregistration attribute of the Client object.
func (c *Clientset) APIRegistration() apiregistrationclientset.Interface {
	return c.apiregistration
}

// RESTClient will return the restClient attribute of the Client object.
func (c *Clientset) RESTClient() rest.Interface {
	return c.restClient
}

// Version returns the GitVersion of the Kubernetes client stored on the object.
func (c *Clientset) Version() string {
	return c.version
}
