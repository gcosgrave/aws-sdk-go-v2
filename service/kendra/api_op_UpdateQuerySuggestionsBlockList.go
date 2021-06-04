// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a block list used for query suggestions for an index. Updates to a block
// list might not take effect right away. Amazon Kendra needs to refresh the entire
// suggestions list to apply any updates to the block list. Other changes not
// related to the block list apply immediately. If a block list is updating, then
// you need to wait for the first update to finish before submitting another
// update. Amazon Kendra supports partial updates, so you only need to provide the
// fields you want to update.
func (c *Client) UpdateQuerySuggestionsBlockList(ctx context.Context, params *UpdateQuerySuggestionsBlockListInput, optFns ...func(*Options)) (*UpdateQuerySuggestionsBlockListOutput, error) {
	if params == nil {
		params = &UpdateQuerySuggestionsBlockListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQuerySuggestionsBlockList", params, optFns, addOperationUpdateQuerySuggestionsBlockListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQuerySuggestionsBlockListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQuerySuggestionsBlockListInput struct {

	// The unique identifier of a block list.
	//
	// This member is required.
	Id *string

	// The identifier of the index for a block list.
	//
	// This member is required.
	IndexId *string

	// The description for a block list.
	Description *string

	// The name of a block list.
	Name *string

	// The IAM (Identity and Access Management) role used to access the block list text
	// file in S3.
	RoleArn *string

	// The S3 path where your block list text file sits in S3. If you update your block
	// list and provide the same path to the block list text file in S3, then Amazon
	// Kendra reloads the file to refresh the block list. Amazon Kendra does not
	// automatically refresh your block list. You need to call the
	// UpdateQuerySuggestionsBlockList API to refresh you block list. If you update
	// your block list, then Amazon Kendra asynchronously refreshes all query
	// suggestions with the latest content in the S3 file. This means changes might not
	// take effect immediately.
	SourceS3Path *types.S3Path
}

type UpdateQuerySuggestionsBlockListOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateQuerySuggestionsBlockListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateQuerySuggestionsBlockList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateQuerySuggestionsBlockList{}, middleware.After)
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
	if err = addOpUpdateQuerySuggestionsBlockListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQuerySuggestionsBlockList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQuerySuggestionsBlockList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "UpdateQuerySuggestionsBlockList",
	}
}
