// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified Spot Instance requests. You can use
// DescribeSpotInstanceRequests to find a running Spot Instance by examining the
// response. If the status of the Spot Instance is fulfilled, the instance ID
// appears in the response and contains the identifier of the instance.
// Alternatively, you can use DescribeInstances
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances)
// with a filter to look for instances where the instance lifecycle is spot. We
// recommend that you set MaxResults to a value between 5 and 1000 to limit the
// number of results returned. This paginates the output, which makes the list more
// manageable and returns the results faster. If the list of results exceeds your
// MaxResults value, then that number of results is returned along with a NextToken
// value that can be passed to a subsequent DescribeSpotInstanceRequests request to
// retrieve the remaining results. Spot Instance requests are deleted four hours
// after they are canceled and their instances are terminated.
func (c *Client) DescribeSpotInstanceRequests(ctx context.Context, params *DescribeSpotInstanceRequestsInput, optFns ...func(*Options)) (*DescribeSpotInstanceRequestsOutput, error) {
	if params == nil {
		params = &DescribeSpotInstanceRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotInstanceRequests", params, optFns, addOperationDescribeSpotInstanceRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotInstanceRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotInstanceRequests.
type DescribeSpotInstanceRequestsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * availability-zone-group - The Availability Zone
	// group.
	//
	// * create-time - The time stamp when the Spot Instance request was
	// created.
	//
	// * fault-code - The fault code related to the request.
	//
	// * fault-message
	// - The fault message related to the request.
	//
	// * instance-id - The ID of the
	// instance that fulfilled the request.
	//
	// * launch-group - The Spot Instance launch
	// group.
	//
	// * launch.block-device-mapping.delete-on-termination - Indicates whether
	// the EBS volume is deleted on instance termination.
	//
	// *
	// launch.block-device-mapping.device-name - The device name for the volume in the
	// block device mapping (for example, /dev/sdh or xvdh).
	//
	// *
	// launch.block-device-mapping.snapshot-id - The ID of the snapshot for the EBS
	// volume.
	//
	// * launch.block-device-mapping.volume-size - The size of the EBS volume,
	// in GiB.
	//
	// * launch.block-device-mapping.volume-type - The type of EBS volume: gp2
	// for General Purpose SSD, io1 or io2 for Provisioned IOPS SSD, st1 for Throughput
	// Optimized HDD, sc1for Cold HDD, or standard for Magnetic.
	//
	// * launch.group-id -
	// The ID of the security group for the instance.
	//
	// * launch.group-name - The name
	// of the security group for the instance.
	//
	// * launch.image-id - The ID of the
	// AMI.
	//
	// * launch.instance-type - The type of instance (for example, m3.medium).
	//
	// *
	// launch.kernel-id - The kernel ID.
	//
	// * launch.key-name - The name of the key pair
	// the instance launched with.
	//
	// * launch.monitoring-enabled - Whether detailed
	// monitoring is enabled for the Spot Instance.
	//
	// * launch.ramdisk-id - The RAM disk
	// ID.
	//
	// * launched-availability-zone - The Availability Zone in which the request
	// is launched.
	//
	// * network-interface.addresses.primary - Indicates whether the IP
	// address is the primary private IP address.
	//
	// *
	// network-interface.delete-on-termination - Indicates whether the network
	// interface is deleted when the instance is terminated.
	//
	// *
	// network-interface.description - A description of the network interface.
	//
	// *
	// network-interface.device-index - The index of the device for the network
	// interface attachment on the instance.
	//
	// * network-interface.group-id - The ID of
	// the security group associated with the network interface.
	//
	// *
	// network-interface.network-interface-id - The ID of the network interface.
	//
	// *
	// network-interface.private-ip-address - The primary private IP address of the
	// network interface.
	//
	// * network-interface.subnet-id - The ID of the subnet for the
	// instance.
	//
	// * product-description - The product description associated with the
	// instance (Linux/UNIX | Windows).
	//
	// * spot-instance-request-id - The Spot Instance
	// request ID.
	//
	// * spot-price - The maximum hourly price for any Spot Instance
	// launched to fulfill the request.
	//
	// * state - The state of the Spot Instance
	// request (open | active | closed | cancelled | failed). Spot request status
	// information can help you track your Amazon EC2 Spot Instance requests. For more
	// information, see Spot request status
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-bid-status.html) in
	// the Amazon EC2 User Guide for Linux Instances.
	//
	// * status-code - The short code
	// describing the most recent evaluation of your Spot Instance request.
	//
	// *
	// status-message - The message explaining the status of the Spot Instance
	// request.
	//
	// * tag: - The key/value combination of a tag assigned to the resource.
	// Use the tag key in the filter name and the tag value as the filter value. For
	// example, to find all resources that have a tag with the key Owner and the value
	// TeamA, specify tag:Owner for the filter name and TeamA for the filter value.
	//
	// *
	// tag-key - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	//
	// *
	// type - The type of Spot Instance request (one-time | persistent).
	//
	// * valid-from
	// - The start date of the request.
	//
	// * valid-until - The end date of the request.
	Filters []types.Filter

	// The maximum number of results to return in a single call. Specify a value
	// between 5 and 1000. To retrieve the remaining results, make another call with
	// the returned NextToken value.
	MaxResults int32

	// The token to request the next set of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// One or more Spot Instance request IDs.
	SpotInstanceRequestIds []string
}

// Contains the output of DescribeSpotInstanceRequests.
type DescribeSpotInstanceRequestsOutput struct {

	// The token to use to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// One or more Spot Instance requests.
	SpotInstanceRequests []types.SpotInstanceRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSpotInstanceRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotInstanceRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotInstanceRequests{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotInstanceRequests(options.Region), middleware.Before); err != nil {
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

// DescribeSpotInstanceRequestsAPIClient is a client that implements the
// DescribeSpotInstanceRequests operation.
type DescribeSpotInstanceRequestsAPIClient interface {
	DescribeSpotInstanceRequests(context.Context, *DescribeSpotInstanceRequestsInput, ...func(*Options)) (*DescribeSpotInstanceRequestsOutput, error)
}

var _ DescribeSpotInstanceRequestsAPIClient = (*Client)(nil)

// DescribeSpotInstanceRequestsPaginatorOptions is the paginator options for
// DescribeSpotInstanceRequests
type DescribeSpotInstanceRequestsPaginatorOptions struct {
	// The maximum number of results to return in a single call. Specify a value
	// between 5 and 1000. To retrieve the remaining results, make another call with
	// the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSpotInstanceRequestsPaginator is a paginator for
// DescribeSpotInstanceRequests
type DescribeSpotInstanceRequestsPaginator struct {
	options   DescribeSpotInstanceRequestsPaginatorOptions
	client    DescribeSpotInstanceRequestsAPIClient
	params    *DescribeSpotInstanceRequestsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSpotInstanceRequestsPaginator returns a new
// DescribeSpotInstanceRequestsPaginator
func NewDescribeSpotInstanceRequestsPaginator(client DescribeSpotInstanceRequestsAPIClient, params *DescribeSpotInstanceRequestsInput, optFns ...func(*DescribeSpotInstanceRequestsPaginatorOptions)) *DescribeSpotInstanceRequestsPaginator {
	options := DescribeSpotInstanceRequestsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeSpotInstanceRequestsInput{}
	}

	return &DescribeSpotInstanceRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSpotInstanceRequestsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeSpotInstanceRequests page.
func (p *DescribeSpotInstanceRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSpotInstanceRequestsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeSpotInstanceRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSpotInstanceRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotInstanceRequests",
	}
}
