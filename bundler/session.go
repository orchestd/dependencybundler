package bundler

import (
	session2 "github.com/orchestd/dependencybundler/constructors/session"
	"github.com/orchestd/dependencybundler/constructors/session/repos"
	"github.com/orchestd/session"
	"github.com/orchestd/session/sessionresolver"
	"go.uber.org/fx"
)

func SessionFxOption() fx.Option {
	return fx.Options(
		fx.Provide(repos.DefaultCacheSessionRepo),
		fx.Provide(func() session.SessionResolverBuilder {
			return sessionresolver.Builder()
		}),
		fx.Provide(session2.DefaultSession),
	)
}
