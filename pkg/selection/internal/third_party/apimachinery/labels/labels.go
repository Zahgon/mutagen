// Subset of k8s.io/apimachinery vendored to avoid pulling in the full
// Kubernetes dependency tree (including klog and its init-time
// goroutines). Originally extracted from the following revision:
// https://github.com/kubernetes/apimachinery/tree/f916759cb6b8547418dc7708876ecab5c1961448
//
// The original code license:
//
// Copyright 2014 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// The original license header inside the code itself:
//

/*
Copyright 2014 The Kubernetes Authors.

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

package labels

import (
	"github.com/mutagen-io/mutagen/pkg/selection/internal/third_party/apimachinery/util/validation/field"
)

// Labels allows you to present labels independently from their storage.
type Labels interface {
	// Has returns whether the provided label exists.
	Has(label string) (exists bool)

	// Get returns the value for the provided label.
	Get(label string) (value string)
}

// Set is a map of label:value. It implements Labels.
type Set map[string]string

// String returns all labels listed as a human readable string.
// Conveniently, exactly the format that ParseSelector takes.
func (ls Set) String() string { _ = "STUB: not implemented"; return "" }

// Sort for determinism.

// Has returns whether the provided label exists in the map.
func (ls Set) Has(label string) bool { _ = "STUB: not implemented"; return false }

// Get returns the value in the map for the provided label.
func (ls Set) Get(label string) string {
	_ = "STUB: not implemented"

	// AsSelector converts labels into a selectors. It does not
	// perform any validation, which means the server will reject
	// the request if the Set contains invalid values.
	return ""
}

func (ls Set) AsSelector() Selector { _ = "STUB: not implemented"; return *new(Selector) }

// AsValidatedSelector converts labels into a selectors.
// The Set is validated client-side, which allows to catch errors early.
func (ls Set) AsValidatedSelector() (Selector, error) {
	_ = "STUB: not implemented"
	return *new(Selector), nil
}

// AsSelectorPreValidated converts labels into a selector, but
// assumes that labels are already validated and thus doesn't
// perform any validation.
// According to our measurements this is significantly faster
// in codepaths that matter at high scale.
func (ls Set) AsSelectorPreValidated() Selector { _ = "STUB: not implemented"; return *new(Selector) }

// FormatLabels converts label map into plain string
func FormatLabels(labelMap map[string]string) string { _ = "STUB: not implemented"; return "" }

// Conflicts takes 2 maps and returns true if there a key match between
// the maps but the value doesn't match, and returns false in other cases
func Conflicts(labels1, labels2 Set) bool { _ = "STUB: not implemented"; return false }

// Merge combines given maps, and does not check for any conflicts
// between the maps. In case of conflicts, second map (labels2) wins
func Merge(labels1, labels2 Set) Set { _ = "STUB: not implemented"; return *new(Set) }

// Equals returns true if the given maps are equal
func Equals(labels1, labels2 Set) bool { _ = "STUB: not implemented"; return false }

// ConvertSelectorToLabelsMap converts selector string to labels map
// and validates keys and values
func ConvertSelectorToLabelsMap(selector string, opts ...field.PathOption) (Set, error) {
	_ = "STUB: not implemented"
	return *new(Set), nil
}
