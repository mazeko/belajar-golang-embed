package belajar_golang_embed

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

//go:embed signature.png
var logo []byte

func TestByteEmbed(t *testing.T) {
	err := ioutil.WriteFile("signature_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
