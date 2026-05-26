package selection

import (
	k8slabels "github.com/mutagen-io/mutagen/pkg/selection/internal/third_party/apimachinery/labels"
)

// LabelSelector is a type that performs matching against a set of labels.
type LabelSelector interface {
	// Matches checks whether or not a set of labels is matched by the selector.
	Matches(labels map[string]string) bool
}

// labelSelector is the internal selector implementation. Internally it uses the
// Kubernetes label selection infrastructure.
type labelSelector struct {
	// k8sLabelSelector is the underlying Kubernetes label selector.
	k8sLabelSelector k8slabels.Selector
}

// Matches implements Selector.Matches.
func (s *labelSelector) Matches(labels map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

// ParseLabelSelector performs label selector parsing. The syntax is currently
// the same as that for Kubernetes.
func ParseLabelSelector(selector string) (LabelSelector, error) {
	_ = "STUB: not implemented"
	// Parse the selector using the Kubernetes label infrastructure.
	return *new(LabelSelector), nil
}

// Wrap up the Kubernetes selector.

// ExtractAndSortLabelKeys extracts a list of keys from the label set and sorts
// them.
func ExtractAndSortLabelKeys(labels map[string]string) []string {
	_ = "STUB: not implemented"
	// Avoid allocation in the event that there are no labels.
	return nil
}

// Create and populate the key slice.

// Sort keys.

// Done.

// EnsureLabelKeyValid verifies that a key conforms to label key requirements.
// These requirements are currently the same as those for Kubernetes label keys.
func EnsureLabelKeyValid(key string) error {
	_ = "STUB: not implemented"
	// Perform validation.
	return nil
}

// Success.

// EnsureLabelValueValid verifies that a value conforms to label value
// requirements. These requirements are currently the same as those for
// Kubernetes label values.
func EnsureLabelValueValid(value string) error {
	_ = "STUB: not implemented"
	// Perform validation.
	return nil
}

// Success.
