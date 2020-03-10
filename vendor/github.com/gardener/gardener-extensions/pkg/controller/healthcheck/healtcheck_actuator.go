// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package healthcheck

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/gardener/gardener-extensions/pkg/util"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Actuator contains all the health checks and the means to execute them
type Actuator struct {
	logger logr.Logger

	restConfig *rest.Config

	seedClient          client.Client
	scheme              *runtime.Scheme
	decoder             runtime.Decoder
	provider            string
	extensionKind       string
	healthCheckMappings map[HealthCheck]string
}

// NewActuator creates a new Actuator.
func NewActuator(provider, extensionKind string, healthChecks map[HealthCheck]string) HealthCheckActuator {
	return &Actuator{
		healthCheckMappings: healthChecks,
		provider:            provider,
		extensionKind:       extensionKind,
		logger:              log.Log.WithName(fmt.Sprintf("%s-%s-healthcheck-actuator", provider, extensionKind)),
	}
}

func (a *Actuator) InjectScheme(scheme *runtime.Scheme) error {
	a.scheme = scheme
	a.decoder = serializer.NewCodecFactory(a.scheme).UniversalDecoder()
	return nil
}

func (a *Actuator) InjectClient(client client.Client) error {
	a.seedClient = client
	return nil
}

func (a *Actuator) InjectConfig(config *rest.Config) error {
	a.restConfig = config
	return nil
}

type healthCheckUnsuccessful struct {
	reason string
	detail string
}

type channelResult struct {
	healthConditionType string
	healthCheckResult   *SingleCheckResult
	error               error
}

type checkResultForConditionType struct {
	failedChecks       []error
	unsuccessfulChecks []healthCheckUnsuccessful
	successfulChecks   int
}

// ExecuteHealthCheckFunctions executes all the health check functions, injects clients and logger & aggregates the results.
// returns an Result for each HealthConditionTyp (e.g  ControlPlaneHealthy)
func (a *Actuator) ExecuteHealthCheckFunctions(ctx context.Context, request types.NamespacedName) (*[]Result, error) {
	_, shootClient, err := util.NewClientForShoot(ctx, a.seedClient, request.Namespace, client.Options{})
	if err != nil {
		msg := fmt.Sprintf("failed to create shoot client in namespace '%s'", request.Namespace)
		a.logger.Error(err, msg)
		return nil, fmt.Errorf(msg)
	}
	channel := make(chan channelResult)
	var wg sync.WaitGroup
	wg.Add(len(a.healthCheckMappings))
	for healthCheck, healthConditionType := range a.healthCheckMappings {
		// clone to avoid problems during parallel execution
		check := healthCheck.DeepCopy()
		check.InjectSeedClient(a.seedClient)
		check.InjectShootClient(shootClient)
		check.SetLoggerSuffix(a.provider, a.extensionKind)
		go func(ctx context.Context, request types.NamespacedName, check HealthCheck, healthConditionType string) {
			defer wg.Done()
			healthCheckResult, err := check.Check(ctx, request)
			channel <- channelResult{
				healthCheckResult:   healthCheckResult,
				error:               err,
				healthConditionType: healthConditionType,
			}
		}(ctx, request, check, healthConditionType)
	}

	// close channel when wait group has 0 counter
	go func() {
		wg.Wait()
		close(channel)
	}()

	groupedHealthCheckResults := make(map[string]*checkResultForConditionType)
	// loop runs until channel is closed
	for channelResult := range channel {
		if groupedHealthCheckResults[channelResult.healthConditionType] == nil {
			groupedHealthCheckResults[channelResult.healthConditionType] = &checkResultForConditionType{}
		}
		if channelResult.error != nil {
			groupedHealthCheckResults[channelResult.healthConditionType].failedChecks = append(groupedHealthCheckResults[channelResult.healthConditionType].failedChecks, channelResult.error)
			continue
		}
		if !channelResult.healthCheckResult.IsHealthy {
			groupedHealthCheckResults[channelResult.healthConditionType].unsuccessfulChecks = append(groupedHealthCheckResults[channelResult.healthConditionType].unsuccessfulChecks, healthCheckUnsuccessful{reason: channelResult.healthCheckResult.Reason, detail: channelResult.healthCheckResult.Detail})
			continue
		}
		groupedHealthCheckResults[channelResult.healthConditionType].successfulChecks++
	}

	var checkResults []Result
	for conditionType, result := range groupedHealthCheckResults {
		if len(result.unsuccessfulChecks) > 0 || len(result.failedChecks) > 0 {
			var details strings.Builder
			if len(result.unsuccessfulChecks) > 0 {
				details.WriteString("Unsuccessful checks: ")
			}
			for index, unsuccessfulCheck := range result.unsuccessfulChecks {
				details.WriteString(fmt.Sprintf("%d) %s: %s. ", index+1, unsuccessfulCheck.reason, unsuccessfulCheck.detail))
			}
			if len(result.failedChecks) > 0 {
				details.WriteString("Failed checks: ")
			}
			for index, err := range result.failedChecks {
				details.WriteString(fmt.Sprintf("%d) %s. ", index+1, err.Error()))
			}
			failureDetails := details.String()
			checkResults = append(checkResults, Result{
				HealthConditionType: conditionType,
				IsHealthy:           false,
				Detail:              &failureDetails,
				SuccessfulChecks:    result.successfulChecks,
				UnsuccessfulChecks:  len(result.unsuccessfulChecks),
				FailedChecks:        len(result.failedChecks),
			})
			continue
		}
		checkResults = append(checkResults, Result{
			HealthConditionType: conditionType,
			IsHealthy:           true,
			SuccessfulChecks:    result.successfulChecks,
		})
	}
	return &checkResults, nil
}
