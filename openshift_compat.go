package ginkgo

// specPriorityShim is a no-op type for SpecPriority compatibility
type specPriorityShim struct{}

// SpecPriority is a deprecated decorator from Ginkgo v1 that set test priority.
// In Ginkgo v2, test ordering is controlled differently and priorities are not supported.
// This is provided as a no-op compatibility shim for legacy code that uses it.
//
// This is an OpenShift-specific patch to maintain compatibility with vendored
// k8s.io/kubernetes/test/e2e/framework code that still uses SpecPriority.
func SpecPriority(priority int) specPriorityShim {
	// No-op: Ginkgo v2 does not support test priorities
	return specPriorityShim{}
}
