package grpc_handler

import (
	"context"
	"email-handler/service"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	appcontext "kotakemail.id/pkg/context"
	email_handler_pb "kotakemail.id/shared/grpc/email_handler"
	"kotakemail.id/shared/schema"
)

type MailboxGrpcHandlerStruct struct {
	email_handler_pb.UnimplementedMailboxServer
	mailboxService service.MailboxService
	appCtx         *appcontext.AppContext
}

func NewMailboxGrpcHandler(
	appCtx *appcontext.AppContext,
) email_handler_pb.MailboxServer {
	return &MailboxGrpcHandlerStruct{
		appCtx:         appCtx,
		mailboxService: appCtx.GetService(service.MailboxServiceKey).(service.MailboxService),
	}
}

func (h *MailboxGrpcHandlerStruct) prepareMailboxResponse(mailbox *schema.Mailbox, err error) *email_handler_pb.MailboxResponse {
	resp := &email_handler_pb.MailboxResponse{
		Success: true,
		Message: "success",
	}

	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp
	}

	data := &email_handler_pb.MailboxData{
		Id:          mailbox.ID.Hex(),
		Name:        mailbox.Name,
		UserId:      mailbox.UserID.Hex(),
		UnreadCount: mailbox.UnreadCount,
		CreatedAt:   timestamppb.New(time.Unix(mailbox.CreatedAt, 0)),
		UpdatedAt:   timestamppb.New(time.Unix(mailbox.UpdatedAt, 0)),
	}

	if mailbox.DeletedAt != 0 {
		data.DeletedAt = timestamppb.New(time.Unix(mailbox.DeletedAt, 0))
	}

	resp.Data = data
	return resp
}

func (h *MailboxGrpcHandlerStruct) CreateMailbox(ctx context.Context, req *email_handler_pb.CreateMailboxRequest) (*email_handler_pb.MailboxResponse, error) {
	mailbox, err := h.mailboxService.CreateMailbox(h.appCtx, req.Name, req.UserId)
	return h.prepareMailboxResponse(mailbox, err), err
}

func (h *MailboxGrpcHandlerStruct) GetMailbox(ctx context.Context, req *email_handler_pb.GetMailboxRequest) (*email_handler_pb.MailboxResponse, error) {
	mailbox, err := h.mailboxService.GetMailbox(h.appCtx, req.Id)
	return h.prepareMailboxResponse(mailbox, err), err
}

func (h *MailboxGrpcHandlerStruct) GetMailboxCredentials(ctx context.Context, req *email_handler_pb.GetMailboxCredentialsRequest) (*email_handler_pb.MailboxCredentialResponse, error) {
	resp := &email_handler_pb.MailboxCredentialResponse{
		Success: true,
		Message: "success",
	}
	mailbox, err := h.mailboxService.GetMailbox(h.appCtx, req.Id)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp, err
	}

	resp.Data = &email_handler_pb.MailboxCredentialData{
		Id:           mailbox.ID.Hex(),
		SMTPUserName: mailbox.SMTPUsername,
		SMTPPassword: mailbox.SMTPPassword,
		ApiKey:       mailbox.ApiKey,
	}
	return resp, nil
}

func (h *MailboxGrpcHandlerStruct) UpdateMailbox(ctx context.Context, req *email_handler_pb.UpdateMailboxRequest) (*email_handler_pb.MailboxResponse, error) {
	mailbox, err := h.mailboxService.UpdateMailbox(h.appCtx, req.Id, req.Name)
	return h.prepareMailboxResponse(mailbox, err), err
}

func (h *MailboxGrpcHandlerStruct) DeleteMailbox(ctx context.Context, req *email_handler_pb.DeleteMailboxRequest) (*email_handler_pb.MailboxResponse, error) {
	err := h.mailboxService.DeleteMailbox(h.appCtx, req.Id)
	return &email_handler_pb.MailboxResponse{
		Success: err == nil,
		Message: err.Error(),
	}, err
}
