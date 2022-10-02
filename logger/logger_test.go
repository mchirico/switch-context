package logger

import (
	"github.com/DaraDadachanji/switch-context/config"
	"github.com/DaraDadachanji/switch-context/fixtures"
	"testing"
)

func TestLog(t *testing.T) {
	config.SetPath(fixtures.Path(".switchcontext"))
	OverrideFileLoc(fixtures.Path("../.switchcontext.log"))

	Log("Test more")
}
