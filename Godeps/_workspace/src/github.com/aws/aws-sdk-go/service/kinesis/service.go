// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kinesis

import (
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/convox/agent/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Amazon Kinesis is a managed service that scales elastically for real time
// processing of streaming big data.
type Kinesis struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new Kinesis client.
func New(config *aws.Config) *Kinesis {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "kinesis",
			APIVersion:   "2013-12-02",
			JSONVersion:  "1.1",
			TargetPrefix: "Kinesis_20131202",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &Kinesis{service}
}

// newRequest creates a new request for a Kinesis operation and runs any
// custom request initialization.
func (c *Kinesis) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}