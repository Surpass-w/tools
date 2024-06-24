package internal

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
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
	// 打开文件，如果不存在则创建
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// 如果文件是新创建的，写入空的 JSON 对象
	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		_, err = f.WriteString("{}")
		if err != nil {
			return err
		}
		// 需要重新打开文件以读取内容
		f.Close()
		f, err = os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	// 解析 JSON 数据
	j, err := simplejson.NewFromReader(f)
	if err != nil {
		return err
	}

	// 设置指定路径的值
	j.SetPath(paths, value)

	// 将修改后的 JSON 数据写回文件
	data, err := j.MarshalJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

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
