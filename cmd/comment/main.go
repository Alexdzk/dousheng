package main

import (
	"net"

	"github.com/Alexdzk/dousheng/dal"
	comment "github.com/Alexdzk/dousheng/kitex_gen/comment/commentservice"
	"github.com/Alexdzk/dousheng/pkg/bound"
	"github.com/Alexdzk/dousheng/pkg/constants"
	"github.com/Alexdzk/dousheng/pkg/middleware"
	tracer2 "github.com/Alexdzk/dousheng/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.CommentServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) //r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.CommentAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}), //server name
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       //address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), //limit
		server.WithMuxTransport(),                                          //Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    //tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                //BoundHandler
		server.WithRegistry(r),                                             //registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
