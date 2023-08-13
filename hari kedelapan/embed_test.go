package hari_kedelapan

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed 3.png
var logo []byte

func TestFile(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

	if err != nil {
		panic(err)
	}

}
