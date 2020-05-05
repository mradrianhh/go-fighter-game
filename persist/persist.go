package persist

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

var lock sync.Mutex

var marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Save saves an object v in json format to location given by path "path".
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create("../content/" + path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

var unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Load loads the file at path "path".
func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open("../content/" + path)
	if err != nil {
		return err
	}
	defer f.Close()
	return unmarshal(f, v)
}
