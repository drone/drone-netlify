// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

// This program downloads an updated pipeline.go file, which
// provides common pipeline structures and environment variable
// mappings.

// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/drone/boilr-plugin/master/template/plugin/pipeline.go")
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("pipeline_gen.go", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
