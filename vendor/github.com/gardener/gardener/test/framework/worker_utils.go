// Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
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

package framework

import (
	"fmt"
	"strings"

	"k8s.io/utils/strings/slices"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/apis/core/v1beta1/helper"
	"github.com/gardener/gardener/pkg/utils"
)

// setShootWorkerSettings sets the Shoot's worker settings from the given config
func setShootWorkerSettings(shoot *gardencorev1beta1.Shoot, cfg *ShootCreationConfig, cloudProfile *gardencorev1beta1.CloudProfile) error {
	if StringSet(cfg.workersConfig) {
		workers, err := ParseFileAsWorkers(cfg.workersConfig)
		if err != nil {
			return err
		}
		shoot.Spec.Provider.Workers = workers
	} else {
		if err := SetupShootWorker(shoot, cloudProfile, cfg.workerZone); err != nil {
			return err
		}
	}

	if StringSet(cfg.shootMachineType) {
		for i := range shoot.Spec.Provider.Workers {
			shoot.Spec.Provider.Workers[i].Machine.Type = cfg.shootMachineType
		}
	}

	if StringSet(cfg.shootMachineImageName) {
		for i := range shoot.Spec.Provider.Workers {
			shoot.Spec.Provider.Workers[i].Machine.Image.Name = cfg.shootMachineImageName
		}
	}

	if StringSet(cfg.shootMachineImageVersion) {
		for i := range shoot.Spec.Provider.Workers {
			shoot.Spec.Provider.Workers[i].Machine.Image.Version = &cfg.shootMachineImageVersion
		}
	}

	return nil
}

// SetupShootWorker prepares the Shoot with one worker with provider specific volume. Clears the currently configured workers.
func SetupShootWorker(shoot *gardencorev1beta1.Shoot, cloudProfile *gardencorev1beta1.CloudProfile, workerZone string) error {
	if len(cloudProfile.Spec.MachineImages) < 1 {
		return fmt.Errorf("at least one different machine image has to be defined in the CloudProfile")
	}

	// clear current workers
	shoot.Spec.Provider.Workers = []gardencorev1beta1.Worker{}

	return AddWorker(shoot, cloudProfile, workerZone)
}

// AddWorker adds a valid default worker to the shoot for the given machineImage and CloudProfile.
func AddWorker(shoot *gardencorev1beta1.Shoot, cloudProfile *gardencorev1beta1.CloudProfile, workerZone string) error {
	if len(cloudProfile.Spec.MachineTypes) == 0 {
		return fmt.Errorf("no MachineTypes configured in the Cloudprofile '%s'", cloudProfile.Name)
	}

	machineType := cloudProfile.Spec.MachineTypes[0]

	//select first machine type of CPU architecture amd64
	for _, machine := range cloudProfile.Spec.MachineTypes {
		if *machine.Architecture == v1beta1constants.ArchitectureAMD64 {
			machineType = machine
			break
		}
	}

	if *machineType.Architecture != v1beta1constants.ArchitectureAMD64 {
		return fmt.Errorf("no MachineTypes of architecture amd64 configured in the Cloudprofile '%s'", cloudProfile.Name)
	}

	machineImage := firstMachineImageWithAMDSupport(cloudProfile.Spec.MachineImages)

	if machineImage == nil {
		return fmt.Errorf("no MachineImage that supports architecture amd64 configured in the Cloudprofile '%s'", cloudProfile.Name)
	}

	qualifyingVersionFound, shootMachineImage, err := helper.GetLatestQualifyingShootMachineImage(*machineImage)
	if err != nil {
		return err
	}

	if !qualifyingVersionFound {
		return fmt.Errorf("could not add worker. No latest qualifying Shoot machine image could be determined for machine image %q. Make sure the machine image in the CloudProfile has at least one version that is not expired and not in preview", machineImage.Name)
	}

	workerName, err := generateRandomWorkerName(fmt.Sprintf("%s-", shootMachineImage.Name))
	if err != nil {
		return err
	}

	shoot.Spec.Provider.Workers = append(shoot.Spec.Provider.Workers, gardencorev1beta1.Worker{
		Name:    workerName,
		Maximum: 2,
		Minimum: 2,
		Machine: gardencorev1beta1.Machine{
			Type:  machineType.Name,
			Image: shootMachineImage,
		},
	})

	if machineType.Storage == nil && len(cloudProfile.Spec.VolumeTypes) > 0 {
		shoot.Spec.Provider.Workers[0].Volume = &gardencorev1beta1.Volume{
			Type:       &cloudProfile.Spec.VolumeTypes[0].Name,
			VolumeSize: "35Gi",
		}
	}

	if StringSet(workerZone) {
		// using one zone as default
		shoot.Spec.Provider.Workers[0].Zones = []string{workerZone}
	}

	return nil
}

func generateRandomWorkerName(prefix string) (string, error) {
	var length int
	remainingCharacters := 15 - len(prefix)
	if remainingCharacters > 0 {
		length = remainingCharacters
	} else {
		prefix = WorkerNamePrefix
		length = 15 - len(WorkerNamePrefix)
	}

	randomString, err := utils.GenerateRandomString(length)
	if err != nil {
		return "", err
	}

	return prefix + strings.ToLower(randomString), nil
}

func firstMachineImageWithAMDSupport(machineImageFromCloudProfile []gardencorev1beta1.MachineImage) *gardencorev1beta1.MachineImage {
	for _, machineImage := range machineImageFromCloudProfile {
		for _, version := range machineImage.Versions {
			if slices.Contains(version.Architectures, v1beta1constants.ArchitectureAMD64) {
				return &machineImage
			}
		}
	}

	return nil
}
