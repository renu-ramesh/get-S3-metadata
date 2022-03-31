//go:generate go run main.go --generate-schema

package main

import (
	"flag"

	"aws-lambda-in-go-lang/cmd"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	runLambda      = flag.Bool("run-lambda", true, "Run service as lambda")
	runApi         = flag.Bool("run-api", false, "Run service as api")
	generateSchema = flag.Bool("generate-schema", false, "Run service as api")
)

func main() {

	// parse the input flags
	flag.Parse()

	// initiate and run as api server
	if *runApi {
		cmd.RunAPIServer()
		return
	}

	// generate schema
	if *generateSchema {
		cmd.GenerateSchema()
		return
	}

	// initiate for lambda
	if *runLambda {
		lambda.Start(cmd.Lambdahandler)
	}

}
