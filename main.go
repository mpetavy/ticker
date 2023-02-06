package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"time"
)

func init() {
	common.Init("1.0.24", "", "", "2018", "service demo", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, start, stop, run, time.Second)
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
	defer common.Done()

	common.Run(nil)
}
