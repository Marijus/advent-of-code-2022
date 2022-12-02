package common

import (
	"os"
)

func GetInput(path string) (result string) {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result = string(b)
	return
}
