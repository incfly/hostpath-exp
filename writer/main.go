// The program runs a simple webserver list the given path.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	fi, err := os.Stat(path)
	var buffer bytes.Buffer
	defer func() {
		w.Write(buffer.Bytes())
	}()

	if err != nil {
		buffer.WriteString(fmt.Sprintf("failed to stat the file %v", err))
		return
	}
	if fi.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("failed to list dir %v", err)))
			return
		}
		for _, f := range files {
			buffer.WriteString(fmt.Sprintf("%v%v<br/>", path, f.Name()))
		}
		return
	}
	// print out the content.
	c, err := ioutil.ReadFile(path)
	if err != nil {
		buffer.WriteString(fmt.Sprintf("failed to read file %v", err))
		return
	}
	buffer.WriteString(string(c))
}

// Populate the hostPath with per service account information.
func populateBySA(sa string) error {
	data := []byte(fmt.Sprintf("hello, %v!\n", sa))
	path := fmt.Sprintf("/host/driver/%v/content.txt", sa)
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		fmt.Printf("failed to populateBySA for %v, err %v\n", sa, err)
		return err
	}
	return nil
}

func main() {
	fmt.Println("hello world, i'm hostpath writer server")

	// Populate information for two sa.
	populateBySA("sa1")
	populateBySA("sa2")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
