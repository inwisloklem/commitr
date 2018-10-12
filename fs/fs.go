package fs

import (
	"commitr/misc"
	"io/ioutil"
)

// ReadFromFile reads the file named by filename and returns the string of file content
func ReadFromFile(f string) string {
	b, err := ioutil.ReadFile(f)

	misc.HandleError(err, "Read file error")

	return string(b)
}
