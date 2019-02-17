package file_detector_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/jtarchie/magic/file_detector"
)

//go:generate ragel -e -G2 -Z magic.rl -o magic.go
//go:generate stringer -type=Type
var (
	tarBuffer, zipBuffer []byte
)

func init() {
	var err error
	tarBuffer, err = ioutil.ReadFile("./filetype/fixtures/sample.tar")
	if err != nil {
		log.Fatal("could not load tar")
	}
	zipBuffer, err = ioutil.ReadFile("./filetype/fixtures/sample.zip")
	if err != nil {
		log.Fatal("could not load zip")
	}
}

var result file_detector.Type

func BenchmarkMatchTar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = file_detector.Detect(tarBuffer)
	}
}

func BenchmarkMatchZip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = file_detector.Detect(zipBuffer)
	}
}
