package service

import (
	"email-handler/repository"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/shared/schema"
)

type MailboxService interface {
	CreateMailbox(ctx *appcontext.AppContext, name, userId string) (*schema.Mailbox, error)
	GetMailbox(ctx *appcontext.AppContext, id string) (*schema.Mailbox, error)
	UpdateMailbox(ctx *appcontext.AppContext, id, name string) (*schema.Mailbox, error)
	DeleteMailbox(ctx *appcontext.AppContext, id string) error
	ListMailbox(ctx *appcontext.AppContext) ([]*schema.Mailbox, error)
}

type mailboxService struct {
	mailboxRepo repository.MailboxRepo
}

func NewMailboxService(appCtx *appcontext.AppContext) MailboxService {
	return &mailboxService{
		mailboxRepo: appCtx.GetRepo(repository.MailboxRepoKey).(repository.MailboxRepo),
	}
}

func (s *mailboxService) CreateMailbox(ctx *appcontext.AppContext, name, userId string) (*schema.Mailbox, error) {
	mailboxObjID := bson.NewObjectID()
	userObjId, _ := bson.ObjectIDFromHex(userId)
	SMTPPassword, _ := bcrypt.GenerateFromPassword([]byte(mailboxObjID.Hex()), bcrypt.DefaultCost)
	mailbox := &schema.Mailbox{
		ID:           mailboxObjID,
		Name:         name,
		UserID:       userObjId,
		ApiKey:       uuid.Must(uuid.NewV7()).String(),
		SMTPUsername: userId,
		SMTPPassword: string(SMTPPassword),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		UnreadCount:  0,
	}
	if err := s.mailboxRepo.Create(ctx, mailbox); err != nil {
		return nil, err
	}
	return mailbox, nil
}

func (s *mailboxService) GetMailbox(ctx *appcontext.AppContext, id string) (*schema.Mailbox, error) {
	mailbox, err := s.mailboxRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mailbox, nil
}

func (s *mailboxService) UpdateMailbox(ctx *appcontext.AppContext, id, name string) (*schema.Mailbox, error) {
	mailbox, err := s.mailboxRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	mailbox.Name = name
	mailbox.UpdatedAt = time.Now().Unix()
	if err := s.mailboxRepo.Update(ctx, mailbox); err != nil {
		return nil, err
	}
	return mailbox, nil
}

func (s *mailboxService) DeleteMailbox(ctx *appcontext.AppContext, id string) error {
	return s.mailboxRepo.Delete(ctx, id)
}

func (s *mailboxService) ListMailbox(ctx *appcontext.AppContext) ([]*schema.Mailbox, error) {
	return s.mailboxRepo.List(ctx)
}
