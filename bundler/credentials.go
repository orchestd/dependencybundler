package bundler

import (
	"github.com/orchestd/dependencybundler/credentials"
	"github.com/orchestd/dependencybundler/credentials/credentialsgetter"

	credentialsConstructor "github.com/orchestd/dependencybundler/constructors/credentials"
	"go.uber.org/fx"
)

func CredentialsFxOption() fx.Option {
	return fx.Options(
		fx.Provide(func() credentials.Builder {
			builder := credentialsgetter.Builder()
			return builder
		}),
		fx.Provide(credentialsConstructor.DefaultCredentials),
	)
}
