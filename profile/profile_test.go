package profile

import (
	"github.com/DaraDadachanji/switch-context/fixtures"
	"testing"
)

func TestProfileEnvExports(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfileEnvExports("usprod")
	if err != nil {
		t.Errorf("Error getting exports: %s", err)
	}
	expected := []string{"export AWS_PROFILE=\"default\"\n",
		"export AWS_REGION=\"us-east-1\"\n"}

	for _, v := range e {
		found := false
		for _, vv := range expected {
			if v == vv {
				found = true
			}
		}
		if !found {
			t.Errorf("Unexpected export: %s", v)
		}
	}

	for _, v := range e {
		t.Logf("%s", v)
	}
}

func TestProfilePS1Exports(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfilePS1Exports("usprod")
	if err != nil {
		t.Errorf("Error getting exports: %s", err)
	}
	expected := []string{"export PS1='\\h:\\W (usp) \\u\\$'\n"}

	if e[0] != expected[0] {
		t.Errorf("Unexpected export: %s", e)
	}

	for _, v := range e {
		t.Logf("%s", v)
	}
}

func TestPR(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	PR("usprod")
}
