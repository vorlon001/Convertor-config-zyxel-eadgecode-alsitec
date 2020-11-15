package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func zipHandler(w http.ResponseWriter, r *http.Request) {
	filename := "toronja.jpeg"
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := writer.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", filename))
	w.Write(buf.Bytes())
	//	io.Copy(w, buf)
}

func main() {
	http.HandleFunc("/zip", zipHandler)
	http.ListenAndServe(":8080", nil)
}
