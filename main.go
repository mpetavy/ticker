package main

import (
	"fmt"
	"time"

	"github.com/mpetavy/common"
)

func init() {
	common.Init(true, "1.0.24", "2018", "service demo", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, start, stop, tick, time.Second)
}

func start() error {
	common.Info("ticker START!!")
	return nil
}

func stop() error {
	common.Info("ticker STOP!!")
	return nil
}

func tick() error {
	common.Info("ticker tick!!")
	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
