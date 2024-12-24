package repository

import (
	"kotakemail.id/pkg/container"
	appcontext "kotakemail.id/pkg/context"
)

const (
	MailboxRepoKey appcontext.RepoContextKey = "mailboxRepo"
)

func InitRepository(container *container.Container) {
	ctx := container.Context()
	db := container.GetDatabase("mongodb-main")

	// register repository
	ctx.SetRepo(MailboxRepoKey, NewMailboxRepo(db))
}
