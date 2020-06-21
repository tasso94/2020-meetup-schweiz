package main

import (
	"context"
	"github.com/antihax/optional"
	"openapi"
	"os"
)

func main() {
	var config = openapi.NewConfiguration()
	var client = openapi.NewAPIClient(config)

	var payload = openapi.ModifyProcessInstanceOpts{
		ProcessInstanceModificationDto: optional.NewInterface(openapi.ProcessInstanceModificationDto{
			Instructions: []openapi.ProcessInstanceModificationInstructionDto{},
		}),
	}

	var ctx = context.Background()
	var processInstanceId = os.Args[1]

}
