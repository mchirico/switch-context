package profile

import (
	"fmt"
	"github.com/mchirico/switch-context/fixtures"
	"testing"
)

func TestProfileArgoExports(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfileArgoExports("usprod")
	if err != nil {
		t.Errorf("Error getting exports: %s", err)
	}
	found := false
	for _, v := range e {
		if v == "unset ARGO_TOKEN\n" {
			found = true
		}
	}
	if !found {
		t.Errorf("Expected unset ARGO_TOKEN")
	}
	fmt.Println(e)

}

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

func Test_ListAllProfiles(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	profiles := ProfilesAvailable()
	count := 0
	for _, v := range profiles {
		if v == "usprod" || v == "ukprod" {
			count++
		}
		t.Logf("%s", v)
	}
	if count != 2 {
		t.Errorf("Expected 2 profiles, got %d", count)
	}
}
