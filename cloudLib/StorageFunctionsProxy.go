package cloudLib

import (
	"CloudStorage/shared"
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"os"
	"reflect"
)

type StorageFunctionsProxy struct {
	host      string
	port      int
	ObjectId  int
	requestor dist.Requestor
}

func NewStorageFunctionsProxy(host string, port int, objectId int) *StorageFunctionsProxy {
	return &StorageFunctionsProxy{host, port, objectId, dist.NewRequestorImpl(host, port)}
}

func (sfp StorageFunctionsProxy) Price(size float64) (price float64, err error) {
	inv := *dist.NewInvocation(sfp.ObjectId, sfp.host, sfp.port, lib.FunctionName(), []interface{}{size})
	termination, err := sfp.requestor.Invoke(inv)
	if err != nil {
		return price, err
	}

	price = termination.Result.([]interface{})[0].(float64)
	//err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return price, err
	}

	return price, nil
}

func (sfp StorageFunctionsProxy) Availability() (available bool, err error) {
	inv := *dist.NewInvocation(sfp.ObjectId, sfp.host, sfp.port, lib.FunctionName(), []interface{}{})
	termination, err := sfp.requestor.Invoke(inv)
	if err != nil {
		return available, err
	}

	available = termination.Result.([]interface{})[0].(bool)
	//err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return available, err
	}

	return available, nil
}

func (sfp StorageFunctionsProxy) SendFile(base64File string, fileName string, path string) (createdFile CloudFile, err error) {
	inv := *dist.NewInvocation(sfp.ObjectId, sfp.host, sfp.port, lib.FunctionName(), []interface{}{base64File, fileName, path})
	termination, err := sfp.requestor.Invoke(inv)
	if err != nil {
		return createdFile, err
	}

	createdFileValue := reflect.New(reflect.ValueOf(createdFile).Type())
	_, err = lib.Decode(termination.Result.([]interface{})[0].(map[string]interface{}), &createdFileValue)
	createdFile = createdFileValue.Elem().Interface().(CloudFile)

	var erro shared.RemoteError
	//err = mapstructure.Decode(termination.Result.([]interface{})[1], &erro)
	errValue := reflect.New(reflect.ValueOf(erro).Type())
	_, err = lib.Decode(termination.Result.([]interface{})[1].(map[string]interface{}), &errValue)
	err = errValue.Elem().Interface().(error)

	//err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return createdFile, err
	}

	return createdFile, erro
}

func (sfp StorageFunctionsProxy) GetFile(fileName string, path string) (file *os.File, err error) {
	inv := *dist.NewInvocation(sfp.ObjectId, sfp.host, sfp.port, lib.FunctionName(), []interface{}{fileName, path})
	termination, err := sfp.requestor.Invoke(inv)
	if err != nil {
		return file, err
	}

	file = termination.Result.([]interface{})[0].(*os.File)
	//err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return file, err
	}

	return file, nil
}

func (sfp StorageFunctionsProxy) List(path string) (files []CloudFile, err error) {
	inv := *dist.NewInvocation(sfp.ObjectId, sfp.host, sfp.port, lib.FunctionName(), []interface{}{path})
	termination, err := sfp.requestor.Invoke(inv)
	if err != nil {
		return files, err
	}

	files = termination.Result.([]interface{})[0].([]CloudFile)
	//err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return files, err
	}

	return files, nil
}

func (sfp StorageFunctionsProxy) Close() error {
	return sfp.requestor.Close()
}
