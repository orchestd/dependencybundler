package session

import (
	"github.com/orchestd/dependencybundler/interfaces/session"
	session2 "github.com/orchestd/session"
)

func DefaultSession(repo session.SessionRepo, builder session2.SessionResolverBuilder) (session.SessionResolver, error) {
	return builder.SetRepo(repo).Build()
}
