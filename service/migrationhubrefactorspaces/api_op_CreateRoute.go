// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an Amazon Web Services Migration Hub Refactor Spaces route. The account
// owner of the service resource is always the environment owner, regardless of
// which account creates the route. Routes target a service in the application. If
// an application does not have any routes, then the first route must be created as
// a DEFAULTRouteType. When created, the default route defaults to an active state
// so state is not a required input. However, like all other state values the state
// of the default route can be updated after creation, but only when all other
// routes are also inactive. Conversely, no route can be active without the default
// route also being active. When you create a route, Refactor Spaces configures the
// Amazon API Gateway to send traffic to the target service as follows:
//
// * If the
// service has a URL endpoint, and the endpoint resolves to a private IP address,
// Refactor Spaces routes traffic using the API Gateway VPC link.
//
// * If the service
// has a URL endpoint, and the endpoint resolves to a public IP address, Refactor
// Spaces routes traffic over the public internet.
//
// * If the service has an Lambda
// function endpoint, then Refactor Spaces configures the Lambda function's
// resource policy to allow the application's API Gateway to invoke the
// function.
//
// A one-time health check is performed on the service when either the
// route is updated from inactive to active, or when it is created with an active
// state. If the health check fails, the route transitions the route state to
// FAILED, an error code of SERVICE_ENDPOINT_HEALTH_CHECK_FAILURE is provided, and
// no traffic is sent to the service. For Lambda functions, the Lambda function
// state is checked. If the function is not active, the function configuration is
// updated so that Lambda resources are provisioned. If the Lambda state is Failed,
// then the route creation fails. For more information, see the
// GetFunctionConfiguration's State response parameter
// (https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunctionConfiguration.html#SSS-GetFunctionConfiguration-response-State)
// in the Lambda Developer Guide. For Lambda endpoints, a check is performed to
// determine that a Lambda function with the specified ARN exists. If it does not
// exist, the health check fails. For public URLs, a connection is opened to the
// public endpoint. If the URL is not reachable, the health check fails. For
// private URLS, a target group is created on the Elastic Load Balancing and the
// target group health check is run. The HealthCheckProtocol, HealthCheckPort, and
// HealthCheckPath are the same protocol, port, and path specified in the URL or
// health URL, if used. All other settings use the default values, as described in
// Health checks for your target groups
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html).
// The health check is considered successful if at least one target within the
// target group transitions to a healthy state. Services can have HTTP or HTTPS URL
// endpoints. For HTTPS URLs, publicly-signed certificates are supported. Private
// Certificate Authorities (CAs) are permitted only if the CA's domain is also
// publicly resolvable.
func (c *Client) CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) {
	if params == nil {
		params = &CreateRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRoute", params, optFns, c.addOperationCreateRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRouteInput struct {

	// The ID of the application within which the route is being created.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The ID of the environment in which the route is created.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The route type of the route. DEFAULT indicates that all traffic that does not
	// match another route is forwarded to the default route. Applications must have a
	// default route before any other routes can be created. URI_PATH indicates a route
	// that is based on a URI path.
	//
	// This member is required.
	RouteType types.RouteType

	// The ID of the service in which the route is created. Traffic that matches this
	// route is forwarded to this service.
	//
	// This member is required.
	ServiceIdentifier *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Configuration for the default route type.
	DefaultRoute *types.DefaultRouteInput

	// The tags to assign to the route. A tag is a label that you assign to an Amazon
	// Web Services resource. Each tag consists of a key-value pair..
	Tags map[string]string

	// The configuration for the URI path route type.
	UriPathRoute *types.UriPathRouteInput

	noSmithyDocumentSerde
}

type CreateRouteOutput struct {

	// The ID of the application in which the route is created.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the route. The format for this ARN is
	// arn:aws:refactor-spaces:region:account-id:resource-type/resource-id . For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	Arn *string

	// The Amazon Web Services account ID of the route creator.
	CreatedByAccountId *string

	// A timestamp that indicates when the route is created.
	CreatedTime *time.Time

	// A timestamp that indicates when the route was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Web Services account ID of the route owner.
	OwnerAccountId *string

	// The unique identifier of the route.
	RouteId *string

	// The route type of the route.
	RouteType types.RouteType

	// The ID of service in which the route is created. Traffic that matches this route
	// is forwarded to this service.
	ServiceId *string

	// The current state of the route. Activation state only allows ACTIVE or INACTIVE
	// as user inputs. FAILED is a route state that is system generated.
	State types.RouteState

	// The tags assigned to the created route. A tag is a label that you assign to an
	// Amazon Web Services resource. Each tag consists of a key-value pair.
	Tags map[string]string

	// Configuration for the URI path route type.
	UriPathRoute *types.UriPathRouteInput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoute{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateRouteMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoute(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRoute struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRoute) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRoute) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRouteInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRouteInput ")
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
func addIdempotencyToken_opCreateRouteMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRoute{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "refactor-spaces",
		OperationName: "CreateRoute",
	}
}
