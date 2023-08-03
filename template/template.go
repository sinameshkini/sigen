package template

import (
	"errors"
	"fmt"
	"github.com/sinameshkini/sigen/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type FileType string

const (
	FSNone      FileType = ""
	FTFile      FileType = "file"
	FTDirectory FileType = "directory"
	FTLink      FileType = "link"
)

type File struct {
	Parent  string
	Name    string
	Type    FileType
	Content string
	Sub     []*File
}

func Make(template, out string, variables map[string]string) (err error) {
	var (
	//dirPath  string
	//filePath string
	)

	if template == "" {
		return errors.New("please select template from your config, example: sigen -t/--template <template key>")
	}

	// get template from config
	tmp, err := getTemplate(template)
	if err != nil {
		return err
	}

	// check output path is exist
	exist, err := utils.Exists(out)
	if err != nil {
		return err
	}

	// if not exist, try to create directory
	if !exist {
		if _, err = utils.Mkdir(out, ""); err != nil {
			return err
		}
	}

	tmp.Parent = out

	// make awesome  code from template :)
	if err = makeTemplate(tmp, variables); err != nil {
		return err
	}

	return nil
}

func makeTemplate(template *File, vars map[string]string) (err error) {
	if template == nil {
		return errors.New("template is null or empty")
	}

	switch template.Type {
	case FSNone:
		return fmt.Errorf("select file type for %s, file, directory or link ", template.Name)
	case FTFile:
		_, err = newFile(template, vars)
		if err != nil {
			return err
		}

	case FTDirectory:
		dirPath, err := newDirectory(template, vars)
		if err != nil {
			return err
		}

		for _, sub := range template.Sub {
			sub.Parent = dirPath
			if err = makeTemplate(sub, vars); err != nil {
				return err
			}
		}
	case FTLink:
		// 	TODO
		logrus.Warnln("link not implement, ", template.Name, "skipped!")
	}

	return nil
}

func newFile(template *File, vars map[string]string) (filePath string, err error) {
	template.Name = replaceVars(template.Name, vars)

	if filePath, err = utils.Touch(template.Parent, template.Name); err != nil {
		return "", err
	}

	template.Content = replaceVars(template.Content, vars)

	if err = utils.WriteToFile(filePath, template.Content); err != nil {
		return "", err
	}

	return filePath, nil
}

func newDirectory(template *File, vars map[string]string) (dirPath string, err error) {
	template.Name = replaceVars(template.Name, vars)

	if dirPath, err = utils.Mkdir(template.Parent, template.Name); err != nil {
		return "", err
	}

	return dirPath, nil
}

func replaceVars(str string, vars map[string]string) string {
	for key, value := range vars {
		str = strings.ReplaceAll(str, fmt.Sprintf("${%s}", key[1:]), value)
	}

	return str
}

func getTemplate(name string) (template *File, err error) {
	if err = viper.UnmarshalKey("templates."+name, &template); err != nil {
		return nil, err
	}

	return
}
