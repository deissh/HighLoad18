package main

import (
	"archive/zip"
	"encoding/json"
	"github.com/deissh/HighLoad18/core/account"
	"log"
	"strings"
)

func LoadFromZip(src string) {
	r, err := zip.OpenReader(src)
	if err != nil {
		log.Fatal(err)
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
				if err == nil {
					log.Println("converting to json ...")

					var acc account.Accounts
					if err := json.NewDecoder(rc).Decode(&acc); err != nil {
						log.Fatal(err)
					}

					accounts.Accounts = append(accounts.Accounts, acc.Accounts...)

					log.Println("file converted to json")
				}
			} else {
				log.Println("ignoring file ... " + f.Name)
			}
		}
	}

	for _, f := range r.File {
		extractAndWriteFile(f)
	}
}
