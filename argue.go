package argue

import (
	"os"
	"path/filepath"
)

func binaryName() string {
	return filepath.Base(os.Args[0])
}
