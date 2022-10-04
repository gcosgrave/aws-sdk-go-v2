// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an impersonation role for the given WorkMail organization. Idempotency
// ensures that an API request completes no more than one time. With an idempotent
// request, if the original request completes successfully, any subsequent retries
// also complete successfully without performing any further actions.
func (c *Client) CreateImpersonationRole(ctx context.Context, params *CreateImpersonationRoleInput, optFns ...func(*Options)) (*CreateImpersonationRoleOutput, error) {
	if params == nil {
		params = &CreateImpersonationRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImpersonationRole", params, optFns, c.addOperationCreateImpersonationRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImpersonationRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImpersonationRoleInput struct {

	// The name of the new impersonation role.
	//
	// This member is required.
	Name *string

	// The WorkMail organization to create the new impersonation role within.
	//
	// This member is required.
	OrganizationId *string

	// The list of rules for the impersonation role.
	//
	// This member is required.
	Rules []types.ImpersonationRule

	// The impersonation role's type. The available impersonation role types are
	// READ_ONLY or FULL_ACCESS.
	//
	// This member is required.
	Type types.ImpersonationRoleType

	// The idempotency token for the client request.
	ClientToken *string

	// The description of the new impersonation role.
	Description *string

	noSmithyDocumentSerde
}

type CreateImpersonationRoleOutput struct {

	// The new impersonation role ID.
	ImpersonationRoleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateImpersonationRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateImpersonationRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateImpersonationRole{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateImpersonationRoleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateImpersonationRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImpersonationRole(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateImpersonationRole struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateImpersonationRole) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateImpersonationRole) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateImpersonationRoleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateImpersonationRoleInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateImpersonationRoleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateImpersonationRole{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateImpersonationRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "CreateImpersonationRole",
	}
}
