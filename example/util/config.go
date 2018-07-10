package util

import (
	"os"

	sls "github.com/aliyun-fc/go-loghub"
)

// Project define Project for test
var Project = &sls.LogProject{
	Name:            os.Getenv("NAME"),
	Endpoint:        os.Getenv("ENDPOINT"),
	AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
	AccessKeySecret: os.Getenv("ACCESS_KEY_SECRET"),
}
