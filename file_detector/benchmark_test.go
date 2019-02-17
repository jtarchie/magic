package file_detector_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/jtarchie/magic/file_detector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate ragel -e -G2 -Z magic.rl -o magic.go
//go:generate stringer -type=Type
var (
	tarBuffer, zipBuffer, jpgBuffer, gifBuffer, pngBuffer []byte
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
	jpgBuffer, err = ioutil.ReadFile("./filetype/fixtures/sample.jpg")
	if err != nil {
		log.Fatal("could not load zip")
	}
	gifBuffer, err = ioutil.ReadFile("./filetype/fixtures/sample.gif")
	if err != nil {
		log.Fatal("could not load zip")
	}
	pngBuffer, err = ioutil.ReadFile("./filetype/fixtures/sample.png")
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

func BenchmarkMatchJpeg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = file_detector.Detect(jpgBuffer)
	}
}

func BenchmarkMatchGif(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = file_detector.Detect(gifBuffer)
	}
}
func BenchmarkMatchPng(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = file_detector.Detect(pngBuffer)
	}
}

var _ = Describe("Ensure results are correct", func() {
	It("for all files", func() {
		Expect(file_detector.Detect(pngBuffer).String()).To(Equal("Png"))
		Expect(file_detector.Detect(gifBuffer).String()).To(Equal("Gif"))
		Expect(file_detector.Detect(jpgBuffer).String()).To(Equal("Jpeg"))
		Expect(file_detector.Detect(zipBuffer).String()).To(Equal("Zip"))
		Expect(file_detector.Detect(tarBuffer).String()).To(Equal("Tar"))
	})
})

func TestFiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Results")
}
