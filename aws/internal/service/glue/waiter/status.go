package waiter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	MLTransformStatusUnknown = "Unknown"
	TriggerStatusUnknown     = "Unknown"
)

// MLTransformStatus fetches the MLTransform and its Status
func MLTransformStatus(conn *glue.Glue, transformId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		input := &glue.GetMLTransformInput{
			TransformId: aws.String(transformId),
		}

		output, err := conn.GetMLTransform(input)

		if err != nil {
			return nil, MLTransformStatusUnknown, err
		}

		if output == nil {
			return output, MLTransformStatusUnknown, nil
		}

		return output, aws.StringValue(output.Status), nil
	}
}

// TriggerStatus fetches the Trigger and its Status
func TriggerStatus(conn *glue.Glue, triggerName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		input := &glue.GetTriggerInput{
			Name: aws.String(triggerName),
		}

		output, err := conn.GetTrigger(input)

		if err != nil {
			return nil, TriggerStatusUnknown, err
		}

		if output == nil {
			return output, TriggerStatusUnknown, nil
		}

		return output, aws.StringValue(output.Trigger.State), nil
	}
}

func GlueDevEndpointStatus(conn *glue.Glue, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getDevEndpointInput := &glue.GetDevEndpointInput{
			EndpointName: aws.String(name),
		}
		endpoint, err := conn.GetDevEndpoint(getDevEndpointInput)
		if err != nil {
			if tfawserr.ErrCodeEquals(err, glue.ErrCodeEntityNotFoundException) {
				return nil, "", nil
			}

			return nil, "", err
		}

		if endpoint == nil {
			return nil, "", nil
		}

		return endpoint, aws.StringValue(endpoint.DevEndpoint.Status), nil
	}
}
