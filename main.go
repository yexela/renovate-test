package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(ctx context.Context, event events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
		fmt.Println("event", event)

		metadata := make(map[string]interface{})

		metadata["httpMethod"] = event.HTTPMethod
		metadata["path"] = event.Path
		metadata["queryStringParameters"] = event.QueryStringParameters
		metadata["headers"] = event.Headers
		metadata["body"] = event.Body
		metadata["isBase64Encoded"] = event.IsBase64Encoded

		metadata["functionName"] = ctx.Value("functionName")
		metadata["functionVersion"] = ctx.Value("functionVersion")

		body, err := json.Marshal(metadata)
		if err != nil {
			return events.ALBTargetGroupResponse{
				StatusCode:        500,
				Body:              err.Error(),
				MultiValueHeaders: MultiValueHeaders,
			}, err
		}

		fmt.Println("metadata", string(body))

		return events.ALBTargetGroupResponse{
			StatusCode:        200,
			Body:              string(body),
			MultiValueHeaders: MultiValueHeaders,
		}, nil
	})
}

var MultiValueHeaders = map[string][]string{
	"Content-Type": {"application/json"},
}
