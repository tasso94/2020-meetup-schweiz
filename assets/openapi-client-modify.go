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
			Instructions: []openapi.ProcessInstanceModificationInstructionDto{
				{
					ActivityId: "zahl-raten-schwer",
					Type:       "cancel",
				},
				{
					ActivityId: "tipp-geben",
					Type:       "startBeforeActivity",
				},
			},
		}),
	}

	var ctx = context.Background()
	var processInstanceId = os.Args[1]

	var _, err = client.ProcessInstanceApi.ModifyProcessInstance(ctx, processInstanceId, &payload)

	if err != nil {
		panic(err)

	} else {
		println("Modifikation Erfolgreich!")

	}
}
