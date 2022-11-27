package test

import (
	"embed"
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

//go:embed files/file1.txt
//go:embed files/file2.txt
//go:embed files/file3.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	file1, _ := files.ReadFile("files/file1.txt")
	fmt.Println(string(file1))

	file2, _ := files.ReadFile("files/file2.txt")
	fmt.Println(string(file2))

	file3, _ := files.ReadFile("files/file3.txt")
	fmt.Println(string(file3))
}

//go:embed files/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
