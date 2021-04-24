package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"time"
)

func init() {
	common.Init(true, "1.0.24", "", "", "2018", "service demo", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, start, stop, run, 0)

	common.Events.NewFuncReceiver(common.EventFlagsParsed{}, func(ev common.Event) {
		if common.IsRunningAsService() {
			common.App().RunTime = time.Second
		}
	})
}

func start() error {
	common.Info("ticker START!!")
	return nil
}

func stop() error {
	common.Info("ticker STOP!!")
	return nil
}

func run() error {
	common.Info("ticker TICK!!")
	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
