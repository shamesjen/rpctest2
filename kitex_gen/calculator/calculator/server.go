// Code generated by Kitex v0.5.2. DO NOT EDIT.
package calculator

import (
	server "github.com/cloudwego/kitex/server"
	calculator "rpctest2/kitex_gen/calculator"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler calculator.Calculator, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
