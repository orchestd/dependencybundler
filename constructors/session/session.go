package session

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/session"
	session2 "bitbucket.org/HeilaSystems/session"
)

func DefaultSession(repo session.SessionRepo,builder session2.SessionResolverBuilder)(session.SessionResolver,error) {
	return builder.SetRepo(repo).Build()
}
