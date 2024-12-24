package service

import (
	"kotakemail.id/pkg/container"
	appcontext "kotakemail.id/pkg/context"
)

const (
	MailboxServiceKey appcontext.ServiceContextKey = "mailboxService"
)

func InitServicesitory(container *container.Container) {
	ctx := container.Context()
	// register repository
	ctx.SetService(MailboxServiceKey, NewMailboxService(ctx))
}
