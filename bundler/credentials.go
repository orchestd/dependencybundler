package bundler

import (
	"github.com/orchestd/configurations/credentials"
	"github.com/orchestd/configurations/credentials/credentialsgetter"

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
