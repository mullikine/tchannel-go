// Autogenerated by thrift-gen. Do not modify.
package keyvalue

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

type TChanAdmin interface {
	TChanBaseService

	ClearAll(ctx thrift.Context) error
}

type TChanKeyValue interface {
	TChanBaseService

	Get(ctx thrift.Context, key string) (string, error)
	Set(ctx thrift.Context, key string, value string) error
}

type TChanBaseService interface {
	HealthCheck(ctx thrift.Context) (string, error)
}

// Implementation of a client and service handler.

type tchanAdminClient struct {
	tchanBaseServiceClient

	client thrift.TChanClient
}

func newTChanAdminClient(client thrift.TChanClient) *tchanAdminClient {
	return &tchanAdminClient{
		*newTChanBaseServiceClient(client),
		client,
	}
}

func NewTChanAdminClient(client thrift.TChanClient) TChanAdmin {
	return newTChanAdminClient(client)
}

func (c *tchanAdminClient) ClearAll(ctx thrift.Context) error {
	var resp AdminClearAllResult
	args := AdminClearAllArgs{}
	success, err := c.client.Call(ctx, "Admin", "clearAll", &args, &resp)
	if err == nil && !success {
		if e := resp.NotAuthorized; e != nil {
			err = e
		}
	}

	return err
}

type tchanAdminServer struct {
	tchanBaseServiceServer

	handler TChanAdmin
}

func newTChanAdminServer(handler TChanAdmin) *tchanAdminServer {
	return &tchanAdminServer{
		*newTChanBaseServiceServer(handler),
		handler,
	}
}

func NewTChanAdminServer(handler TChanAdmin) thrift.TChanServer {
	return newTChanAdminServer(handler)
}

func (s *tchanAdminServer) Service() string {
	return "Admin"
}

func (s *tchanAdminServer) Methods() []string {
	return []string{
		"clearAll",

		"HealthCheck",
	}
}

func (s *tchanAdminServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "clearAll":
		return s.handleClearAll(ctx, protocol)
	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanAdminServer) handleClearAll(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req AdminClearAllArgs
	var res AdminClearAllResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ClearAll(ctx)

	if err != nil {
		switch v := err.(type) {
		case *NotAuthorized:
			res.NotAuthorized = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

type tchanKeyValueClient struct {
	tchanBaseServiceClient

	client thrift.TChanClient
}

func newTChanKeyValueClient(client thrift.TChanClient) *tchanKeyValueClient {
	return &tchanKeyValueClient{
		*newTChanBaseServiceClient(client),
		client,
	}
}

func NewTChanKeyValueClient(client thrift.TChanClient) TChanKeyValue {
	return newTChanKeyValueClient(client)
}

func (c *tchanKeyValueClient) Get(ctx thrift.Context, key string) (string, error) {
	var resp KeyValueGetResult
	args := KeyValueGetArgs{
		Key: key,
	}
	success, err := c.client.Call(ctx, "KeyValue", "Get", &args, &resp)
	if err == nil && !success {
		if e := resp.NotFound; e != nil {
			err = e
		}
		if e := resp.InvalidKey; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanKeyValueClient) Set(ctx thrift.Context, key string, value string) error {
	var resp KeyValueSetResult
	args := KeyValueSetArgs{
		Key:   key,
		Value: value,
	}
	success, err := c.client.Call(ctx, "KeyValue", "Set", &args, &resp)
	if err == nil && !success {
		if e := resp.InvalidKey; e != nil {
			err = e
		}
	}

	return err
}

type tchanKeyValueServer struct {
	tchanBaseServiceServer

	handler TChanKeyValue
}

func newTChanKeyValueServer(handler TChanKeyValue) *tchanKeyValueServer {
	return &tchanKeyValueServer{
		*newTChanBaseServiceServer(handler),
		handler,
	}
}

func NewTChanKeyValueServer(handler TChanKeyValue) thrift.TChanServer {
	return newTChanKeyValueServer(handler)
}

func (s *tchanKeyValueServer) Service() string {
	return "KeyValue"
}

func (s *tchanKeyValueServer) Methods() []string {
	return []string{
		"Get",
		"Set",

		"HealthCheck",
	}
}

func (s *tchanKeyValueServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "Get":
		return s.handleGet(ctx, protocol)
	case "Set":
		return s.handleSet(ctx, protocol)
	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanKeyValueServer) handleGet(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req KeyValueGetArgs
	var res KeyValueGetResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Get(ctx, req.Key)

	if err != nil {
		switch v := err.(type) {
		case *KeyNotFound:
			res.NotFound = v
		case *InvalidKey:
			res.InvalidKey = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = &r
	}

	return err == nil, &res, nil
}

func (s *tchanKeyValueServer) handleSet(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req KeyValueSetArgs
	var res KeyValueSetResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Set(ctx, req.Key, req.Value)

	if err != nil {
		switch v := err.(type) {
		case *InvalidKey:
			res.InvalidKey = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

type tchanBaseServiceClient struct {
	client thrift.TChanClient
}

func newTChanBaseServiceClient(client thrift.TChanClient) *tchanBaseServiceClient {
	return &tchanBaseServiceClient{
		client,
	}
}

func NewTChanBaseServiceClient(client thrift.TChanClient) TChanBaseService {
	return newTChanBaseServiceClient(client)
}

func (c *tchanBaseServiceClient) HealthCheck(ctx thrift.Context) (string, error) {
	var resp BaseServiceHealthCheckResult
	args := BaseServiceHealthCheckArgs{}
	success, err := c.client.Call(ctx, "baseService", "HealthCheck", &args, &resp)
	if err == nil && !success {
	}

	return resp.GetSuccess(), err
}

type tchanBaseServiceServer struct {
	handler TChanBaseService
}

func newTChanBaseServiceServer(handler TChanBaseService) *tchanBaseServiceServer {
	return &tchanBaseServiceServer{
		handler,
	}
}

func NewTChanBaseServiceServer(handler TChanBaseService) thrift.TChanServer {
	return newTChanBaseServiceServer(handler)
}

func (s *tchanBaseServiceServer) Service() string {
	return "baseService"
}

func (s *tchanBaseServiceServer) Methods() []string {
	return []string{
		"HealthCheck",
	}
}

func (s *tchanBaseServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "HealthCheck":
		return s.handleHealthCheck(ctx, protocol)
	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanBaseServiceServer) handleHealthCheck(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BaseServiceHealthCheckArgs
	var res BaseServiceHealthCheckResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.HealthCheck(ctx)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = &r
	}

	return err == nil, &res, nil
}
