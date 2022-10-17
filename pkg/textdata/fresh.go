//for importing a completely new text file

package textdata

import (
	"fmt"
	"os"
	"path/filepath"
)

func NewFile(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	rawPath, err := filepath.Abs("data/raw/" + filepath.Base(path))
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile(rawPath, b, 0777)

	procPath, err := filepath.Abs("data/processed/" + filepath.Base(path))
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile(procPath, b, 0777)
}
