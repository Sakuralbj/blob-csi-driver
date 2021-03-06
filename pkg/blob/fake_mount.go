/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package blob

import (
	"fmt"
	"strings"

	"k8s.io/utils/mount"
)

type fakeMounter struct {
	mount.FakeMounter
}

// MountSensitive overrides mount.FakeMounter.MountSensitive.
func (f *fakeMounter) MountSensitive(source string, target string, fstype string, options []string, sensitiveOptions []string) error {
	if strings.Contains(source, "ut-container") {
		return fmt.Errorf("fake MountSensitive: source error")
	} else if strings.Contains(target, "error_mount_sens") {
		return fmt.Errorf("fake MountSensitive: target error")
	}

	return nil
}
