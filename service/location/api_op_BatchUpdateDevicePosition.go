// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads position update data for one or more devices to a tracker resource.
// Amazon Location uses the data when reporting the last known device position and
// position history. Only one position update is stored per sample time. Location
// data is sampled at a fixed rate of one position per 30-second interval and
// retained for 30 days before it's deleted.
func (c *Client) BatchUpdateDevicePosition(ctx context.Context, params *BatchUpdateDevicePositionInput, optFns ...func(*Options)) (*BatchUpdateDevicePositionOutput, error) {
	if params == nil {
		params = &BatchUpdateDevicePositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateDevicePosition", params, optFns, addOperationBatchUpdateDevicePositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateDevicePositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateDevicePositionInput struct {

	// The name of the tracker resource to update.
	//
	// This member is required.
	TrackerName *string

	// Contains the position update details for each device.
	//
	// This member is required.
	Updates []types.DevicePositionUpdate
}

type BatchUpdateDevicePositionOutput struct {

	// Contains error details for each device that failed to update its position.
	//
	// This member is required.
	Errors []types.BatchUpdateDevicePositionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchUpdateDevicePositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateDevicePosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateDevicePosition{}, middleware.After)
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
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addOpBatchUpdateDevicePositionValidationMiddleware(stack); err != nil {
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
