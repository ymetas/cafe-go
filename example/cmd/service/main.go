package main

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/Louis0504/icafe/example/gen-go/thrift/content_thrift/content"
	"github.com/Louis0504/icafe/example/service"
	"github.com/Louis0504/icafe/server"
	"github.com/Louis0504/icafe/server/rpc"
)

func main() {

	srv := service.ContentService{}
	contentServiceProcessor := content.NewContentServiceProcessor(srv)

	servicesMap := map[string]thrift.TProcessor{
		"ContentService": contentServiceProcessor,
	}

	app := server.NewApplication()
	bundle := rpc.NewTRPCBundle("content-service", rpc.WithTRPCServiceMap(servicesMap), rpc.TRPCListen("0.0.0.0:9000"))
	app.AddBundle(bundle)
	app.Run()

	/*
		var transportFactory thrift.TTransportFactory
		var protocolFactory thrift.TProtocolFactory
		transport, _ := thrift.NewTServerSocket("0.0.0.0:8000")
		server := thrift.NewTSimpleServer4(contentServiceProcessor, transport, transportFactory, protocolFactory)
		server.Serve()
	*/
}
