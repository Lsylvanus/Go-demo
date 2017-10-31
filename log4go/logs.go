package main

import (
	"fmt"
	"os"
	logs "github.com/YoungPioneers/blog4go"
)

// optionally set user defined hook for logging
type MyHook struct {
	something string
}

// when log-level exceed level, call the hook
// level is the level associate with that logging action.
// message is the formatted string already written.
func (self *MyHook) Fire(level logs.LevelType, args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	// init a file write using xml config file
	// log.SetBufferSize(0) // close buffer for in time logging when debugging
	err := logs.NewWriterFromConfigAsFile("./etc/config.xml")
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer logs.Close()

	// initialize your hook instance
	hook := new(MyHook)
	logs.SetHook(hook) // writersFromConfig can be replaced with writers
	logs.SetHookLevel(logs.INFO)
	logs.SetHookAsync(true) // hook will be called in async mode

	// optionally set output colored
	logs.SetColored(true)

	logs.Debugf("Good morning, %s", "eddie")
	logs.Warn("It's time to have breakfast")

	logs.Info("hallo go.")
	logs.Error("there is a error.")
}