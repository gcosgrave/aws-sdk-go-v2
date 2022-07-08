// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the transition lifecycle of a recovery point. The lifecycle defines when a
// protected resource is transitioned to cold storage and when it expires. Backup
// transitions and expires backups automatically according to the lifecycle that
// you define. Backups transitioned to cold storage must be stored in cold storage
// for a minimum of 90 days. Therefore, the “retention” setting must be 90 days
// greater than the “transition to cold after days” setting. The “transition to
// cold after days” setting cannot be changed after a backup has been transitioned
// to cold. Resource types that are able to be transitioned to cold storage are
// listed in the "Lifecycle to cold storage" section of the  Feature availability
// by resource
// (https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html#features-by-resource)
// table. Backup ignores this expression for other resource types. This operation
// does not support continuous backups.
func (c *Client) UpdateRecoveryPointLifecycle(ctx context.Context, params *UpdateRecoveryPointLifecycleInput, optFns ...func(*Options)) (*UpdateRecoveryPointLifecycleOutput, error) {
	if params == nil {
		params = &UpdateRecoveryPointLifecycleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecoveryPointLifecycle", params, optFns, c.addOperationUpdateRecoveryPointLifecycleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecoveryPointLifecycleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRecoveryPointLifecycleInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	//
	// This member is required.
	RecoveryPointArn *string

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. Backup transitions and expires backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the
	// “retention” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold.
	Lifecycle *types.Lifecycle

	noSmithyDocumentSerde
}

type UpdateRecoveryPointLifecycleOutput struct {

	// An ARN that uniquely identifies a backup vault; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// A CalculatedLifecycle object containing DeleteAt and MoveToColdStorageAt
	// timestamps.
	CalculatedLifecycle *types.CalculatedLifecycle

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. Backup transitions and expires backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the
	// “retention” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold. Resource types that are able to be
	// transitioned to cold storage are listed in the "Lifecycle to cold storage"
	// section of the  Feature availability by resource
	// (https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html#features-by-resource)
	// table. Backup ignores this expression for other resource types.
	Lifecycle *types.Lifecycle

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRecoveryPointLifecycleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecoveryPointLifecycle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecoveryPointLifecycle{}, middleware.After)
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
	if err = addOpUpdateRecoveryPointLifecycleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecoveryPointLifecycle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecoveryPointLifecycle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "UpdateRecoveryPointLifecycle",
	}
}
