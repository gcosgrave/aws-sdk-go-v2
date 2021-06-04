// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a block list used for query suggestions for an index. A deleted block
// list might not take effect right away. Amazon Kendra needs to refresh the entire
// suggestions list to add back the queries that were previously blocked.
func (c *Client) DeleteQuerySuggestionsBlockList(ctx context.Context, params *DeleteQuerySuggestionsBlockListInput, optFns ...func(*Options)) (*DeleteQuerySuggestionsBlockListOutput, error) {
	if params == nil {
		params = &DeleteQuerySuggestionsBlockListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteQuerySuggestionsBlockList", params, optFns, addOperationDeleteQuerySuggestionsBlockListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteQuerySuggestionsBlockListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQuerySuggestionsBlockListInput struct {

	// The unique identifier of the block list that needs to be deleted.
	//
	// This member is required.
	Id *string

	// The identifier of the you want to delete a block list from.
	//
	// This member is required.
	IndexId *string
}

type DeleteQuerySuggestionsBlockListOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteQuerySuggestionsBlockListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteQuerySuggestionsBlockList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteQuerySuggestionsBlockList{}, middleware.After)
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
	if err = addOpDeleteQuerySuggestionsBlockListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQuerySuggestionsBlockList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteQuerySuggestionsBlockList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DeleteQuerySuggestionsBlockList",
	}
}
