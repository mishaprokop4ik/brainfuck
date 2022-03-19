package input

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type File struct {
	path string
	text string
}

func NewFile(param string, isFile bool) (*File, error) {
	var f = &File{}
	if isFile {
		f.path = param
		if err := f.read(); err != nil {
			return nil, err
		}
	} else {
		f.text = param
	}

	if err := f.removeSpecSymbols(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) read() error {
	if !strings.HasSuffix(f.path, ".b") {
		return fmt.Errorf("incorrect file type")
	}
	file, err := os.Open(f.path)
	if err != nil {
		return fmt.Errorf("can not open file %s", err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("can not read file %s", err)
	}

	f.text = string(b)
	return nil
}

func (f *File) removeSpecSymbols() error {
	r, err := regexp.Compile("[{} \n\t]")
	if err != nil {
		return fmt.Errorf("can complie regular expression %s", err)
	}
	f.text = r.ReplaceAllString(f.text, "")
	return nil
}

func (f *File) OutText() string {
	return f.text
}