package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/Sifchain/sifnode/app"
	"github.com/Sifchain/sifnode/cmd/sifnoded/cmd"

	"log"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("sifnode"),
		tracer.WithEnv("test"),
	)
	defer tracer.Stop()

	if err := profiler.Start(
		profiler.WithService("sifnode"),
		profiler.WithEnv("test"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,

			// The profiles below are disabled by
			// default to keep overhead low, but
			// can be enabled as needed.
			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	); err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()
	// Start a root span.
	span := tracer.StartSpan("get.data")
	defer span.Finish()

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
