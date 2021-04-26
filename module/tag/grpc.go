package tag

import (
	context "context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"stock.tao/module/core"
)

var (
	grpcServiceDesc = grpc.ServiceDesc{
		ServiceName: "Tag",
		HandlerType: (*Server)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "List",
				Handler:    grpcList,
			},
			{
				MethodName: "ListMember",
				Handler:    grpcListMember,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "tag.proto",
	}
)

func init() {
	core.GrpcServer.RegisterService(&grpcServiceDesc, &ServerImpl{})
}

// grpcList ==> grpc list handler
func grpcList(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PbTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tag/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Server).List(ctx, req.(*PbTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// grpcListMember ==> grpc list member handler
func grpcListMember(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PbTagMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Server).ListMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tag/ListMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Server).ListMember(ctx, req.(*PbTagMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Server ==> server inteface
type Server interface {
	List(context.Context, *PbTagListRequest) (*PbStockTao, error)
	ListMember(context.Context, *PbTagMemberRequest) (*PbStockTao, error)
}

// ServerImpl ==> server implement
type ServerImpl struct{}

// List ==> implement Server inteface
func (ServerImpl) List(ctx context.Context, request *PbTagListRequest) (*PbStockTao, error) {
	err := func(request *PbTagListRequest) error {
		// todo
		return nil
	}(request)
	if err != nil {
		return &PbStockTao{
			Code: uint32(codes.DataLoss),
			Msg:  codes.DataLoss.String(),
		}, nil
	}
	// todo
	return &PbStockTao{
		Code: uint32(core.Success),
		Msg:  codes.OK.String(),
		Data: nil,
	}, nil
}

// ListMember ==> implement Server inteface
func (ServerImpl) ListMember(ctx context.Context, request *PbTagMemberRequest) (*PbStockTao, error) {
	// todo
	return &PbStockTao{
		Code: uint32(core.Success),
		Msg:  codes.OK.String(),
		Data: nil,
	}, nil
}
