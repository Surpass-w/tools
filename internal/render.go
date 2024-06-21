package internal

import (
	"github.com/noirbizarre/gonja"
	"github.com/noirbizarre/gonja/ext/django"
	"github.com/pkg/errors"
	"os"
)

var RenderOptions struct {
	FromPath string
	ToPath   string
	MetaData string
}

func RenderFile(from string, to string, data map[string]interface{}) error {
	f, err := os.OpenFile(to, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0655)
	if err != nil {
		return errors.Wrap(err, "OpenFile")
	}
	defer f.Close()

	tpl := gonja.Must(gonja.FromFile(from))
	// 与 cluster 保持一致，添加扩展函数
	tpl.Env.Filters.Update(django.Filters)
	tpl.Env.Statements.Update(django.Statements)

	content, err := tpl.ExecuteBytes(data)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	return err
}

func Render(from []byte, data map[string]interface{}) ([]byte, error) {
	tpl := gonja.Must(gonja.FromBytes(from))
	return tpl.ExecuteBytes(data)
}
