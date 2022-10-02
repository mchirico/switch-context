package logger

import (
	"github.com/mchirico/switch-context/config"
	"github.com/mchirico/switch-context/fixtures"
	"testing"
)

func TestLog(t *testing.T) {
	config.SetPath(fixtures.Path(".switchcontext"))
	OverrideFileLoc(fixtures.Path("../.switchcontext.log"))

	Log("Test more")
}
