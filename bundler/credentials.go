package bundler

import (
	"bitbucket.org/HeilaSystems/configurations/credentials"
	"bitbucket.org/HeilaSystems/configurations/credentials/credentialsgetter"

	credentialsConstructor "bitbucket.org/HeilaSystems/dependencybundler/constructors/credentials"
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
