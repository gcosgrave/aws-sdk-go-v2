// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the tracker resource details.
func (c *Client) DescribeTracker(ctx context.Context, params *DescribeTrackerInput, optFns ...func(*Options)) (*DescribeTrackerOutput, error) {
	if params == nil {
		params = &DescribeTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTracker", params, optFns, addOperationDescribeTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrackerInput struct {

	// The name of the tracker resource.
	//
	// This member is required.
	TrackerName *string
}

type DescribeTrackerOutput struct {

	// The timestamp for when the tracker resource was created in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// The optional description for the tracker resource.
	//
	// This member is required.
	Description *string

	// The pricing plan selected for the specified tracker resource. For additional
	// details and restrictions on each pricing plan option, see the Amazon Location
	// Service pricing page (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan types.PricingPlan

	// The Amazon Resource Name (ARN) for the tracker resource. Used when you need to
	// specify a resource across all AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:tracker/ExampleTracker
	//
	// This member is required.
	TrackerArn *string

	// The name of the tracker resource.
	//
	// This member is required.
	TrackerName *string

	// The timestamp for when the tracker resource was last updated in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// A key identifier for an AWS KMS customer managed key
	// (https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html)
	// assigned to the Amazon Location resource.
	KmsKeyId *string

	// The specified data provider for the tracker resource.
	PricingPlanDataSource *string

	// The tags associated with the tracker resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTracker{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeTrackerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}
