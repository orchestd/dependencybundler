package bundler

import (
	session2 "bitbucket.org/HeilaSystems/dependencybundler/constructors/session"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/session/repos"
	"bitbucket.org/HeilaSystems/session"
	"bitbucket.org/HeilaSystems/session/sessionresolver"
	"go.uber.org/fx"
)

func SessionFxOption()fx.Option {
	return fx.Options(
		fx.Provide(repos.DefaultCacheSessionRepo),
		fx.Provide(func()session.SessionResolverBuilder {
			return sessionresolver.Builder()
		}),
		fx.Provide(session2.DefaultSession),
		)
}

var DataVersionsKey = sessionresolver.DataVersionsKey
var DataNowKey = sessionresolver.DataNowKey
