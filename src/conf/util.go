package conf

import (
	"os"
	"path/filepath"

	"github.com/triole/logseal"
)

func absPathFatal(path string, lg logseal.Logseal) string {
	fullPath, err := filepath.Abs(path)
	lg.IfErrFatal(
		"unable to construct full path",
		logseal.F{"path": path, "error": err},
	)
	return fullPath
}

func existsFatal(path string, lg logseal.Logseal) {
	_, err := os.Stat(path)
	lg.IfErrFatal(
		"file does not exist",
		logseal.F{"path": path, "error": err},
	)
}
