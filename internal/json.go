package internal

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"os"
)

var JsonOptions struct {
	JsonPath string
	K        string
	V        string
	T        int64
}

func WriteFile(content, filepath string) error {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func Set(filepath, key string, value interface{}) error {
	var jsonstr string
	var err error
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		jsonstr = "{}"
	} else {
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("cannot read file %s, err: %s", filepath, err)
		}
		jsonstr = string(data)
	}
	res, err := sjson.Set(jsonstr, key, value)
	if err != nil {
		return err
	}
	err = WriteFile(res, filepath)
	if err != nil {
		return err
	}
	return nil
}

func Get(filepath, key string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("cannot read file %s, err: %s", filepath, err)
	}
	value := gjson.Get(string(data), key)
	return value.String(), nil
}
