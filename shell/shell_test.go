package shell

import (
	"fmt"
	"testing"
)

func TestShellout(t *testing.T) {
	out, errout, err := Shellout("ls -l")
	fmt.Println(out, errout, err)
}
