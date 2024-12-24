package grpc_handler

import (
	"google.golang.org/grpc"
	appcontext "kotakemail.id/pkg/context"
	email_handler_pb "kotakemail.id/shared/grpc/email_handler"
)

func RegisterGrpcServices(s grpc.ServiceRegistrar, ctx *appcontext.AppContext) {
	// mailbox Handler
	email_handler_pb.RegisterMailboxServer(s, NewMailboxGrpcHandler(ctx))
}
