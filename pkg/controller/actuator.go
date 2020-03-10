// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package controller

import (
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	"github.com/gardener/gardener-extensions/pkg/controller/network"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	// StatusTypeMeta is the TypeMeta of Cilium Status
	StatusTypeMeta = metav1.TypeMeta{
		APIVersion: ciliumv1alpha1.SchemeGroupVersion.String(),
		Kind:       "NetworkStatus",
	}
)

type actuator struct {
	logger logr.Logger

	chartRendererFactory extensionscontroller.ChartRendererFactory
	restConfig           *rest.Config

	client  client.Client
	scheme  *runtime.Scheme
	decoder runtime.Decoder
}

const LogID = "network-cilium-actuator"

// NewActuator creates a new Actuator that updates the status of the handled Network resources.
func NewActuator(chartRendererFactory extensionscontroller.ChartRendererFactory) network.Actuator {
	return &actuator{
		logger:               log.Log.WithName(LogID),
		chartRendererFactory: chartRendererFactory,
	}
}

func (a *actuator) InjectScheme(scheme *runtime.Scheme) error {
	a.scheme = scheme
	a.decoder = serializer.NewCodecFactory(a.scheme).UniversalDecoder()
	return nil
}

func (a *actuator) InjectClient(client client.Client) error {
	a.client = client
	return nil
}

func (a *actuator) InjectConfig(config *rest.Config) error {
	a.restConfig = config
	return nil
}
