package profile

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/mchirico/switch-context/fixtures"
	"os"
	"testing"
)

func TestProfileAliasExportsNoFound(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfileAliasExports("ukprod")

	if err != nil {
		t.Errorf("Expected error, got nil")
	}
	fmt.Println(e)
}

func TestProfileAliasExports(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfileAliasExports("usprod")
	if err != nil {
		t.Errorf("Error getting exports: %s", err)
	}

	// Get Expected values
	data, _ := os.ReadFile(fixtures.Path("expected_output/alias_exports.data"))
	var buf bytes.Buffer
	buf.Write(data)

	dec := gob.NewDecoder(&buf)
	expected := []string{}
	_ = dec.Decode(&expected)

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

	fmt.Println(e)

}

func TestProfileArgoExports(t *testing.T) {
	err := SetPath(fixtures.Path(".switchcontext"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfileArgoExports("usprod")
	if err != nil {
		t.Errorf("Error getting exports: %s", err)
	}

	// Get Expected values
	data, _ := os.ReadFile(fixtures.Path("expected_output/argo_exports.data"))
	var buf bytes.Buffer
	buf.Write(data)

	dec := gob.NewDecoder(&buf)
	expected := []string{}
	_ = dec.Decode(&expected)

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
