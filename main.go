package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/jtarchie/magic/file_detector"
)

func main() {
	filename := flag.String("filename", "", "name to detect")
	flag.Parse()

	if *filename == "" {
		log.Fatalf("please specify a filename")
	}

	log.Printf("filename: %s", *filename)
	contents, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("detected: %s", file_detector.Detect(contents))
}
