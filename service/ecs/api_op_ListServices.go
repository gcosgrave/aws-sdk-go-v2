// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of services. You can filter the results by cluster, launch type,
// and scheduling strategy.
func (c *Client) ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) {
	if params == nil {
		params = &ListServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServices", params, optFns, addOperationListServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServicesInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster to use when
	// filtering the ListServices results. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	// The launch type to use when filtering the ListServices results.
	LaunchType types.LaunchType

	// The maximum number of service results returned by ListServices in paginated
	// output. When this parameter is used, ListServices only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListServices
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter is not used, then ListServices returns up to 10 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a ListServices request indicating that more
	// results are available to fulfill the request and further calls will be needed.
	// If maxResults was provided, it is possible the number of results to be fewer
	// than maxResults. This token should be treated as an opaque identifier that is
	// only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	// The scheduling strategy to use when filtering the ListServices results.
	SchedulingStrategy types.SchedulingStrategy
}

type ListServicesOutput struct {

	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// The list of full ARN entries for each service associated with the specified
	// cluster.
	ServiceArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServices(options.Region), middleware.Before); err != nil {
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

// ListServicesAPIClient is a client that implements the ListServices operation.
type ListServicesAPIClient interface {
	ListServices(context.Context, *ListServicesInput, ...func(*Options)) (*ListServicesOutput, error)
}

var _ ListServicesAPIClient = (*Client)(nil)

// ListServicesPaginatorOptions is the paginator options for ListServices
type ListServicesPaginatorOptions struct {
	// The maximum number of service results returned by ListServices in paginated
	// output. When this parameter is used, ListServices only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListServices
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter is not used, then ListServices returns up to 10 results and a
	// nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServicesPaginator is a paginator for ListServices
type ListServicesPaginator struct {
	options   ListServicesPaginatorOptions
	client    ListServicesAPIClient
	params    *ListServicesInput
	nextToken *string
	firstPage bool
}

// NewListServicesPaginator returns a new ListServicesPaginator
func NewListServicesPaginator(client ListServicesAPIClient, params *ListServicesInput, optFns ...func(*ListServicesPaginatorOptions)) *ListServicesPaginator {
	if params == nil {
		params = &ListServicesInput{}
	}

	options := ListServicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListServices page.
func (p *ListServicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListServices(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListServices",
	}
}
