package paths

import (
	"os"
)

var PATHNAME = "./.config"

func CreateConfigDirectory() bool {
	if _, err := os.Stat(PATHNAME); os.IsNotExist(err) {
		err = os.Mkdir(PATHNAME, 0777)

		if err != nil {
			return false
		}
	}

	return true
}
