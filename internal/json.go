package internal

import (
	"github.com/bitly/go-simplejson"
	"os"
)

var (
	TypeInt    = "int"
	TypeString = "string"
	TypeBool   = "bool"
)

var JsonOptions struct {
	Paths []string
	V     string
	T     string
}

func Set(filePath string, paths []string, value interface{}) error {
	f, err := os.Open(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		f, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	j, err := simplejson.NewFromReader(f)
	if err != nil {
		return err
	}
	j.SetPath(paths, value)
	j.Encode()
	return nil
}

func Get(filePath string, paths []string) (*simplejson.Json, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	j, err := simplejson.NewFromReader(f)
	if err != nil {
		return nil, err
	}
	return j.GetPath(paths...), nil
}
