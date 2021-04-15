package rtc

import (
	grpc "google.golang.org/grpc"
	"stock.tao/module/core"
)

var (
	grpcServiceDesc = grpc.ServiceDesc{
		ServiceName: "RTC",
		HandlerType: (*RTCServer)(nil),
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "exchange",
				Handler:       exchange,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "grpc.proto",
	}
	ExchangeServerMap = make(map[string]*ExchangeServerImpl)
)

func init() {
	core.GrpcServer.RegisterService(&grpcServiceDesc, &ExchangeServerImpl{})
}

type RTCServer interface {
	Exchange(ExchangeServer) error
}

type RTCServerImpl struct {
}

func (RTCServerImpl) Exchange(s ExchangeServer) error {

	wd, err := s.Recv()
	if err != nil {
		return err
	}
	rs := ExchangeServerMap[wd.RemoteID]
	if rs == nil {
		ExchangeServerMap[wd.LocalID] = s.(*ExchangeServerImpl)
		return s.Send(&WebRTCDescription{LocalID: wd.LocalID, RemoteID: "", Sd: nil, Icd: nil})
	}
	return rs.Send(wd)

}

func exchange(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RTCServer).Exchange(&ExchangeServerImpl{stream})
}

// Server ==> server inteface
type ExchangeServer interface {
	Send(*WebRTCDescription) error
	Recv() (*WebRTCDescription, error)
	grpc.ServerStream
}

// ServerImpl ==> server implement
type ExchangeServerImpl struct {
	grpc.ServerStream
}

func (x *ExchangeServerImpl) Send(m *WebRTCDescription) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ExchangeServerImpl) Recv() (*WebRTCDescription, error) {
	m := new(WebRTCDescription)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
