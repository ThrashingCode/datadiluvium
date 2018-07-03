package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Instructions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "For instructions, check out http://datadiluvium.com. Thanks!\n")
}

/*
Test instructions here:
 See README.md for test instructions.
*/

func DataGenerate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := ioutil.ReadAll(r.Body)
	ds := make([]DataStore, 0)
	json.Unmarshal(b, &ds)
	j, _ := json.Marshal(ds)
	w.Write(j)
}

func DataGenerateDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "For instructions on the schema generation, check out schema generation docs at http://datadiluvium.com. Thanks!\n")
}

