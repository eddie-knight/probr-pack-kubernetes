package kubernetes

import (
	"testing"

	"github.com/citihub/probr-sdk/config"
	"github.com/citihub/probr-sdk/probeengine"
)

func TestGetProbes(t *testing.T) {
	pack := make([]probeengine.Probe, 0)
	pack = GetProbes()
	// No required vars set
	if len(pack) > 0 {
		t.Logf("Unexpected value returned from GetProbes")
		t.Fail()
	}
	// 1 of 2 required vars set
	config.Vars.ServicePacks.Kubernetes.AuthorisedContainerRegistry = "not-empty"
	pack = GetProbes()
	if len(pack) > 0 {
		t.Logf("Unexpected value returned from GetProbes")
		t.Fail()
	}
	// All required vars set
	config.Vars.ServicePacks.Kubernetes.UnauthorisedContainerRegistry = "not-empty"
	pack = GetProbes()
	if len(pack) == 0 {
		t.Logf("Expected value not returned from GetProbes")
		t.Fail()
	}
}
