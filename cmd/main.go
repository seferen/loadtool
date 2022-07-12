package main

import (
	"github.com/seferen/loadtesttool/core"
	"github.com/seferen/loadtesttool/httpDsl"
	"github.com/seferen/logger"
)

var log = logger.GetLoger("main.go")

func main() {

	log.Info("Start Application")

	core.SetUp(
		*core.Scenario("test").ExecFn(
			func(session *core.Session) {
				session.Set("test", "test")
			}).ExecFn(func(session *core.Session) {
			session.Set("legacy", "fuck")
		}).Exec(httpDsl.Http("test")),
	)
}

func init() {
	logger.SetRootLevel(logger.DebugLogLevel)
}
