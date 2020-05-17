package imports

import (
	"os"
	"testing"

	"cuelang.org/go/tools/lsp/internal/testenv"
)

func TestMain(m *testing.M) {
	testenv.ExitIfSmallMachine()
	os.Exit(m.Run())
}
