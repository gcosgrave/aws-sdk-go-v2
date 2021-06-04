// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the permissions mode of a ledger.
func (c *Client) UpdateLedgerPermissionsMode(ctx context.Context, params *UpdateLedgerPermissionsModeInput, optFns ...func(*Options)) (*UpdateLedgerPermissionsModeOutput, error) {
	if params == nil {
		params = &UpdateLedgerPermissionsModeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLedgerPermissionsMode", params, optFns, addOperationUpdateLedgerPermissionsModeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLedgerPermissionsModeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLedgerPermissionsModeInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The permissions mode to assign to the ledger. This parameter can have one of the
	// following values:
	//
	// * ALLOW_ALL: A legacy permissions mode that enables access
	// control with API-level granularity for ledgers. This mode allows users who have
	// SendCommand permissions for this ledger to run all PartiQL commands (hence,
	// ALLOW_ALL) on any tables in the specified ledger. This mode disregards any
	// table-level or command-level IAM permissions policies that you create for the
	// ledger.
	//
	// * STANDARD: (Recommended) A permissions mode that enables access
	// control with finer granularity for ledgers, tables, and PartiQL commands. By
	// default, this mode denies all user requests to run any PartiQL commands on any
	// tables in this ledger. To allow PartiQL commands to run, you must create IAM
	// permissions policies for specific table resources and PartiQL actions, in
	// addition to SendCommand API permissions for the ledger.
	//
	// We strongly recommend
	// using the STANDARD permissions mode to maximize the security of your ledger
	// data.
	//
	// This member is required.
	PermissionsMode types.PermissionsMode
}

type UpdateLedgerPermissionsModeOutput struct {

	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string

	// The name of the ledger.
	Name *string

	// The current permissions mode of the ledger.
	PermissionsMode types.PermissionsMode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateLedgerPermissionsModeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLedgerPermissionsMode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLedgerPermissionsMode{}, middleware.After)
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
	if err = addOpUpdateLedgerPermissionsModeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLedgerPermissionsMode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLedgerPermissionsMode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "UpdateLedgerPermissionsMode",
	}
}
