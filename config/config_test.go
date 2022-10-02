package config

import (
	"github.com/DaraDadachanji/switch-context/fixtures"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Simple Test",
			want:    `\h:\W (usp) \u\$`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		SetPath(fixtures.Path(".switchcontext"))
		t.Run(tt.name, func(t *testing.T) {

			if Get("profiles.usprod.bash.PS1") != tt.want {
				t.Errorf("ReadConfigFile() = %v, want %v", Get("profiles.usprod.bash"), tt.want)
			}

		})
	}
}

func TestGetMap(t *testing.T) {
	SetPath(fixtures.Path(".switchcontext"))
	m := GetMap("profiles.usprod.env")
	for k, v := range m {
		t.Logf("%s: %s", k, v)
	}
}

func TestLogger(t *testing.T) {
	SetPath(fixtures.Path(".switchcontext"))
	m := Get("log.maxsize")
	t.Logf("%s", m)
}

func TestGetProfiles(t *testing.T) {
	SetPath(fixtures.Path(".switchcontext"))
	m := GetMap("profiles")
	for k, _ := range m {
		t.Logf("%s", k)
	}
}
