// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/query"
)

// A load balancer distributes incoming traffic across your EC2 instances. This
// enables you to increase the availability of your application. The load balancer
// also monitors the health of its registered instances and ensures that it
// routes traffic only to healthy instances. You configure your load balancer
// to accept incoming traffic by specifying one or more listeners, which are
// configured with a protocol and port number for connections from clients to
// the load balancer and a protocol and port number for connections from the
// load balancer to the instances.
//
// Elastic Load Balancing supports two types of load balancers: Classic load
// balancers and Application load balancers (new). A Classic load balancer makes
// routing and load balancing decisions either at the transport layer (TCP/SSL)
// or the application layer (HTTP/HTTPS), and supports either EC2-Classic or
// a VPC. An Application load balancer makes routing and load balancing decisions
// at the application layer (HTTP/HTTPS), supports path-based routing, and can
// route requests to one or more ports on each EC2 instance or container instance
// in your virtual private cloud (VPC). For more information, see the .
//
// This reference covers the 2012-06-01 API, which supports Classic load balancers.
// The 2015-12-01 API supports Application load balancers.
//
// To get started, create a load balancer with one or more listeners using CreateLoadBalancer.
// Register your instances with the load balancer using RegisterInstancesWithLoadBalancer.
//
// All Elastic Load Balancing operations are idempotent, which means that they
// complete at most one time. If you repeat an operation, it succeeds with a
// 200 OK response code.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01
type ELB struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "elasticloadbalancing" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName            // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ELB client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ELB client from just a session.
//     svc := elb.New(mySession)
//
//     // Create a ELB client with additional configuration
//     svc := elb.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ELB {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ELB {
	svc := &ELB{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-06-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ELB operation and runs any
// custom request initialization.
func (c *ELB) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
