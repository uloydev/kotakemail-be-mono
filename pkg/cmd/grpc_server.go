package cmd

import (
	"net"

	"google.golang.org/grpc"
	"kotakemail.id/config"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/logger"
)

type GrpcServer struct {
	app     *grpc.Server
	conn    net.Listener
	logger  *logger.Logger
	cfg     *config.Config
	options GrpcServerOptions
	name    string
}

type RegisterServiceFunc func(s grpc.ServiceRegistrar, ctx *appcontext.AppContext)
type GrpcServerOptions struct {
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	RegisterService    RegisterServiceFunc
}

func (o *GrpcServerOptions) AddUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	o.UnaryInterceptors = append(o.UnaryInterceptors, interceptors...)
}

func (o *GrpcServerOptions) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	o.StreamInterceptors = append(o.StreamInterceptors, interceptors...)
}

func NewGrpcServer(
	name string,
	cfg *config.Config,
	appLogger *logger.Logger,
	options GrpcServerOptions,
) Command {
	return &GrpcServer{
		name:    name,
		cfg:     cfg,
		logger:  appLogger,
		options: options,
	}
}

func (g *GrpcServer) Execute() error {
	g.logger.Info().Msgf("starting grpc server %s on %s:%s", g.name, g.cfg.Grpc.Server.Host, g.cfg.Grpc.Server.Port)
	conn, err := net.Listen("tcp", g.cfg.Grpc.Server.Host+":"+g.cfg.Grpc.Server.Port)
	if err != nil {
		g.logger.Error().Err(err).Msgf("failed to listen on %s:%s", g.cfg.Grpc.Server.Host, g.cfg.Grpc.Server.Port)
	}
	g.conn = conn

	g.app = grpc.NewServer(
		grpc.ChainUnaryInterceptor(g.options.UnaryInterceptors...),
		grpc.ChainStreamInterceptor(g.options.StreamInterceptors...),
	)

	return g.app.Serve(g.conn)
}

func (g *GrpcServer) Shutdown() error {
	g.logger.Info().Msgf("shutting down grpc server %s", g.name)
	g.app.GracefulStop()
	return g.conn.Close()
}

func (g *GrpcServer) App() interface{} {
	return g.app
}

func (g *GrpcServer) Name() string {
	return g.name
}
