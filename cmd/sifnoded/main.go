package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/Sifchain/sifnode/app"
	"github.com/Sifchain/sifnode/cmd/sifnoded/cmd"
	// "io/ioutil"
	// "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	// "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	// rules := []tracer.SamplingRule{tracer.RateRule(1)}
	// tracer.Start(
	// 	tracer.WithSamplingRules(rules),
	// 	tracer.WithService("sifnode"),
	// 	tracer.WithEnv("test"),
	// )
	// defer tracer.Stop()

	// // Start a root span.
	// span := tracer.StartSpan("get.data")
	// defer span.Finish()

	// // Create a child of it, computing the time needed to read a file.
	// child := tracer.StartSpan("read.file", tracer.ChildOf(span.Context()))
	// child.SetTag(ext.ResourceName, "test.json")

	// // Perform an operation.
	// _, err := ioutil.ReadFile("~/test.json")

	// // We may finish the child span using the returned error. If it's
	// // nil, it will be disregarded.
	// child.Finish(tracer.WithError(err))

	rootCmd, _ := cmd.NewRootCmd()

	app.SetConfig(true)

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
