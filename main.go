package main

import (
	"embed"
	"github.com/mpetavy/common"
	"time"
)

//go:embed go.mod
var resources embed.FS

func init() {
	common.Init("", "", "", "", "service demo", "", "", "", &resources, start, stop, run, time.Second)
}

func start() error {
	common.Info("%s ticker START!!", time.Now().Format(common.DateTimeMilliMask))
	return nil
}

func stop() error {
	common.Info("%s ticker STOP!!", time.Now().Format(common.DateTimeMilliMask))
	return nil
}

func run() error {
	common.Info("%s ticker tick!!", time.Now().Format(common.DateTimeMilliMask))
	return nil
}

func main() {
	common.Run(nil)
}
