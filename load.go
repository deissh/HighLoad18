package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"strings"
)

func LoadFromZip(src string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()
		if f.FileInfo().IsDir() {
			//os.MkdirAll(path, f.Mode())
			log.Println("loading folder ... " + f.Name)
		} else {
			log.Println("loading file ... " + f.Name)

			if strings.Contains(f.Name, ".json") {
				byteValue, err := ioutil.ReadAll(rc)
				if err == nil {
					go parseFile(byteValue)
				} else {
					log.Println("ignoring file ... " + f.Name)
					log.Fatal(err)
				}
			}
		}
	}

	for _, f := range r.File {
		extractAndWriteFile(f)
	}

	return nil
}