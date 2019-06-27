package aws

import (
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
)

type AwsFunctionsProxy struct {
	host      string
	port      int
	ObjectId  int
	requestor dist.Requestor
}

func NewAwsFunctionsProxy(host string, port int, objectId int) *AwsFunctionsProxy {
	return &AwsFunctionsProxy{host, port, objectId, dist.NewRequestorImpl(host, port)}
}

func (afp AwsFunctionsProxy) Price(size float64) (price float64, err error) {
	inv := *dist.NewInvocation(afp.ObjectId, afp.host, afp.port, lib.FunctionName(), []interface{}{size})
	termination, err := afp.requestor.Invoke(inv)
	if err != nil {
		return price, err
	}

	price = termination.Result.([]interface{})[0].(float64)
	err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return price, err
	}

	return price, nil
}

func (afp AwsFunctionsProxy) Availability() (available bool, err error) {
	inv := *dist.NewInvocation(afp.ObjectId, afp.host, afp.port, lib.FunctionName(), []interface{}{})
	termination, err := afp.requestor.Invoke(inv)
	if err != nil {
		return available, err
	}

	available = termination.Result.([]interface{})[0].(bool)
	err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return available, err
	}

	return available, nil
}

func (afp AwsFunctionsProxy) Close() error {
	return afp.requestor.Close()
}
