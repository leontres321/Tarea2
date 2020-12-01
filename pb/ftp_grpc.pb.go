// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FTPClient is the client API for FTP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FTPClient interface {
	Enviar(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Respuesta, error)
	Descargar(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Chunk, error)
}

type fTPClient struct {
	cc grpc.ClientConnInterface
}

func NewFTPClient(cc grpc.ClientConnInterface) FTPClient {
	return &fTPClient{cc}
}

func (c *fTPClient) Enviar(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/ftp.FTP/Enviar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fTPClient) Descargar(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/ftp.FTP/Descargar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FTPServer is the server API for FTP service.
// All implementations must embed UnimplementedFTPServer
// for forward compatibility
type FTPServer interface {
	Enviar(context.Context, *Chunk) (*Respuesta, error)
	Descargar(context.Context, *Nombre) (*Chunk, error)
	mustEmbedUnimplementedFTPServer()
}

// UnimplementedFTPServer must be embedded to have forward compatible implementations.
type UnimplementedFTPServer struct {
}

func (UnimplementedFTPServer) Enviar(context.Context, *Chunk) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enviar not implemented")
}
func (UnimplementedFTPServer) Descargar(context.Context, *Nombre) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Descargar not implemented")
}
func (UnimplementedFTPServer) mustEmbedUnimplementedFTPServer() {}

// UnsafeFTPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FTPServer will
// result in compilation errors.
type UnsafeFTPServer interface {
	mustEmbedUnimplementedFTPServer()
}

func RegisterFTPServer(s grpc.ServiceRegistrar, srv FTPServer) {
	s.RegisterService(&_FTP_serviceDesc, srv)
}

func _FTP_Enviar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTPServer).Enviar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.FTP/Enviar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTPServer).Enviar(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _FTP_Descargar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nombre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTPServer).Descargar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.FTP/Descargar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTPServer).Descargar(ctx, req.(*Nombre))
	}
	return interceptor(ctx, in, info, handler)
}

var _FTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ftp.FTP",
	HandlerType: (*FTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enviar",
			Handler:    _FTP_Enviar_Handler,
		},
		{
			MethodName: "Descargar",
			Handler:    _FTP_Descargar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ftp.proto",
}

// FTPDistribuidoClient is the client API for FTPDistribuido service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FTPDistribuidoClient interface {
	EnviarD(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Respuesta, error)
	DescargarD(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Chunk, error)
	AvisoEscrituraD(ctx context.Context, in *Aviso, opts ...grpc.CallOption) (*Respuesta, error)
}

type fTPDistribuidoClient struct {
	cc grpc.ClientConnInterface
}

func NewFTPDistribuidoClient(cc grpc.ClientConnInterface) FTPDistribuidoClient {
	return &fTPDistribuidoClient{cc}
}

func (c *fTPDistribuidoClient) EnviarD(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/ftp.FTPDistribuido/EnviarD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fTPDistribuidoClient) DescargarD(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/ftp.FTPDistribuido/DescargarD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fTPDistribuidoClient) AvisoEscrituraD(ctx context.Context, in *Aviso, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/ftp.FTPDistribuido/AvisoEscrituraD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FTPDistribuidoServer is the server API for FTPDistribuido service.
// All implementations must embed UnimplementedFTPDistribuidoServer
// for forward compatibility
type FTPDistribuidoServer interface {
	EnviarD(context.Context, *Chunk) (*Respuesta, error)
	DescargarD(context.Context, *Nombre) (*Chunk, error)
	AvisoEscrituraD(context.Context, *Aviso) (*Respuesta, error)
	mustEmbedUnimplementedFTPDistribuidoServer()
}

// UnimplementedFTPDistribuidoServer must be embedded to have forward compatible implementations.
type UnimplementedFTPDistribuidoServer struct {
}

func (UnimplementedFTPDistribuidoServer) EnviarD(context.Context, *Chunk) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarD not implemented")
}
func (UnimplementedFTPDistribuidoServer) DescargarD(context.Context, *Nombre) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescargarD not implemented")
}
func (UnimplementedFTPDistribuidoServer) AvisoEscrituraD(context.Context, *Aviso) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvisoEscrituraD not implemented")
}
func (UnimplementedFTPDistribuidoServer) mustEmbedUnimplementedFTPDistribuidoServer() {}

// UnsafeFTPDistribuidoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FTPDistribuidoServer will
// result in compilation errors.
type UnsafeFTPDistribuidoServer interface {
	mustEmbedUnimplementedFTPDistribuidoServer()
}

func RegisterFTPDistribuidoServer(s grpc.ServiceRegistrar, srv FTPDistribuidoServer) {
	s.RegisterService(&_FTPDistribuido_serviceDesc, srv)
}

func _FTPDistribuido_EnviarD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTPDistribuidoServer).EnviarD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.FTPDistribuido/EnviarD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTPDistribuidoServer).EnviarD(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _FTPDistribuido_DescargarD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nombre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTPDistribuidoServer).DescargarD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.FTPDistribuido/DescargarD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTPDistribuidoServer).DescargarD(ctx, req.(*Nombre))
	}
	return interceptor(ctx, in, info, handler)
}

func _FTPDistribuido_AvisoEscrituraD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Aviso)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTPDistribuidoServer).AvisoEscrituraD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.FTPDistribuido/AvisoEscrituraD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTPDistribuidoServer).AvisoEscrituraD(ctx, req.(*Aviso))
	}
	return interceptor(ctx, in, info, handler)
}

var _FTPDistribuido_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ftp.FTPDistribuido",
	HandlerType: (*FTPDistribuidoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarD",
			Handler:    _FTPDistribuido_EnviarD_Handler,
		},
		{
			MethodName: "DescargarD",
			Handler:    _FTPDistribuido_DescargarD_Handler,
		},
		{
			MethodName: "AvisoEscrituraD",
			Handler:    _FTPDistribuido_AvisoEscrituraD_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ftp.proto",
}

// LOGClient is the client API for LOG service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LOGClient interface {
	EnviarPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Propuesta, error)
	SolicitarUbicacion(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Propuesta, error)
	PedirLibros(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListaLibros, error)
}

type lOGClient struct {
	cc grpc.ClientConnInterface
}

func NewLOGClient(cc grpc.ClientConnInterface) LOGClient {
	return &lOGClient{cc}
}

func (c *lOGClient) EnviarPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Propuesta, error) {
	out := new(Propuesta)
	err := c.cc.Invoke(ctx, "/ftp.LOG/EnviarPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOGClient) SolicitarUbicacion(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Propuesta, error) {
	out := new(Propuesta)
	err := c.cc.Invoke(ctx, "/ftp.LOG/SolicitarUbicacion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOGClient) PedirLibros(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListaLibros, error) {
	out := new(ListaLibros)
	err := c.cc.Invoke(ctx, "/ftp.LOG/PedirLibros", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LOGServer is the server API for LOG service.
// All implementations must embed UnimplementedLOGServer
// for forward compatibility
type LOGServer interface {
	EnviarPropuesta(context.Context, *Propuesta) (*Propuesta, error)
	SolicitarUbicacion(context.Context, *Nombre) (*Propuesta, error)
	PedirLibros(context.Context, *Empty) (*ListaLibros, error)
	mustEmbedUnimplementedLOGServer()
}

// UnimplementedLOGServer must be embedded to have forward compatible implementations.
type UnimplementedLOGServer struct {
}

func (UnimplementedLOGServer) EnviarPropuesta(context.Context, *Propuesta) (*Propuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPropuesta not implemented")
}
func (UnimplementedLOGServer) SolicitarUbicacion(context.Context, *Nombre) (*Propuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarUbicacion not implemented")
}
func (UnimplementedLOGServer) PedirLibros(context.Context, *Empty) (*ListaLibros, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirLibros not implemented")
}
func (UnimplementedLOGServer) mustEmbedUnimplementedLOGServer() {}

// UnsafeLOGServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LOGServer will
// result in compilation errors.
type UnsafeLOGServer interface {
	mustEmbedUnimplementedLOGServer()
}

func RegisterLOGServer(s grpc.ServiceRegistrar, srv LOGServer) {
	s.RegisterService(&_LOG_serviceDesc, srv)
}

func _LOG_EnviarPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Propuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOGServer).EnviarPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.LOG/EnviarPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOGServer).EnviarPropuesta(ctx, req.(*Propuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOG_SolicitarUbicacion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nombre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOGServer).SolicitarUbicacion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.LOG/SolicitarUbicacion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOGServer).SolicitarUbicacion(ctx, req.(*Nombre))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOG_PedirLibros_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOGServer).PedirLibros(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.LOG/PedirLibros",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOGServer).PedirLibros(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LOG_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ftp.LOG",
	HandlerType: (*LOGServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarPropuesta",
			Handler:    _LOG_EnviarPropuesta_Handler,
		},
		{
			MethodName: "SolicitarUbicacion",
			Handler:    _LOG_SolicitarUbicacion_Handler,
		},
		{
			MethodName: "PedirLibros",
			Handler:    _LOG_PedirLibros_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ftp.proto",
}

// LOGDistribuidoClient is the client API for LOGDistribuido service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LOGDistribuidoClient interface {
	EnviarPropuestaD(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Respuesta, error)
	SolicitarUbicacionD(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Propuesta, error)
	PedirLibrosD(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListaLibros, error)
}

type lOGDistribuidoClient struct {
	cc grpc.ClientConnInterface
}

func NewLOGDistribuidoClient(cc grpc.ClientConnInterface) LOGDistribuidoClient {
	return &lOGDistribuidoClient{cc}
}

func (c *lOGDistribuidoClient) EnviarPropuestaD(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/ftp.LOGDistribuido/EnviarPropuestaD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOGDistribuidoClient) SolicitarUbicacionD(ctx context.Context, in *Nombre, opts ...grpc.CallOption) (*Propuesta, error) {
	out := new(Propuesta)
	err := c.cc.Invoke(ctx, "/ftp.LOGDistribuido/SolicitarUbicacionD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOGDistribuidoClient) PedirLibrosD(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListaLibros, error) {
	out := new(ListaLibros)
	err := c.cc.Invoke(ctx, "/ftp.LOGDistribuido/PedirLibrosD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LOGDistribuidoServer is the server API for LOGDistribuido service.
// All implementations must embed UnimplementedLOGDistribuidoServer
// for forward compatibility
type LOGDistribuidoServer interface {
	EnviarPropuestaD(context.Context, *Propuesta) (*Respuesta, error)
	SolicitarUbicacionD(context.Context, *Nombre) (*Propuesta, error)
	PedirLibrosD(context.Context, *Empty) (*ListaLibros, error)
	mustEmbedUnimplementedLOGDistribuidoServer()
}

// UnimplementedLOGDistribuidoServer must be embedded to have forward compatible implementations.
type UnimplementedLOGDistribuidoServer struct {
}

func (UnimplementedLOGDistribuidoServer) EnviarPropuestaD(context.Context, *Propuesta) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPropuestaD not implemented")
}
func (UnimplementedLOGDistribuidoServer) SolicitarUbicacionD(context.Context, *Nombre) (*Propuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarUbicacionD not implemented")
}
func (UnimplementedLOGDistribuidoServer) PedirLibrosD(context.Context, *Empty) (*ListaLibros, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirLibrosD not implemented")
}
func (UnimplementedLOGDistribuidoServer) mustEmbedUnimplementedLOGDistribuidoServer() {}

// UnsafeLOGDistribuidoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LOGDistribuidoServer will
// result in compilation errors.
type UnsafeLOGDistribuidoServer interface {
	mustEmbedUnimplementedLOGDistribuidoServer()
}

func RegisterLOGDistribuidoServer(s grpc.ServiceRegistrar, srv LOGDistribuidoServer) {
	s.RegisterService(&_LOGDistribuido_serviceDesc, srv)
}

func _LOGDistribuido_EnviarPropuestaD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Propuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOGDistribuidoServer).EnviarPropuestaD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.LOGDistribuido/EnviarPropuestaD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOGDistribuidoServer).EnviarPropuestaD(ctx, req.(*Propuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOGDistribuido_SolicitarUbicacionD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nombre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOGDistribuidoServer).SolicitarUbicacionD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.LOGDistribuido/SolicitarUbicacionD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOGDistribuidoServer).SolicitarUbicacionD(ctx, req.(*Nombre))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOGDistribuido_PedirLibrosD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOGDistribuidoServer).PedirLibrosD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ftp.LOGDistribuido/PedirLibrosD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOGDistribuidoServer).PedirLibrosD(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LOGDistribuido_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ftp.LOGDistribuido",
	HandlerType: (*LOGDistribuidoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarPropuestaD",
			Handler:    _LOGDistribuido_EnviarPropuestaD_Handler,
		},
		{
			MethodName: "SolicitarUbicacionD",
			Handler:    _LOGDistribuido_SolicitarUbicacionD_Handler,
		},
		{
			MethodName: "PedirLibrosD",
			Handler:    _LOGDistribuido_PedirLibrosD_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ftp.proto",
}
