package main

import (
	"time"

	"github.com/mpetavy/common"
)

func start() error {
	common.Info("ticker start!!")
	return nil
}

func stop() error {
	common.Info("ticker stop!!")
	return nil
}

func tick() error {
	common.Info("ticker tick!!")
	return nil
}

func main() {
	defer common.Cleanup()

	common.New(&common.App{"ticker", "1.0.24", "2018", "service demo", "mpetavy", common.APACHE, "https://github.com/mpetavy/worktime", true, start, stop, tick, time.Duration(5) * time.Second}, nil)
	common.Run()
}
