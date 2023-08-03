package main

import (
	"github.com/sirupsen/logrus"
	"sigen/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
