package bundler

import (
	"bitbucket.org/HeilaSystems/configurations/credentialsgetter"
	"go.uber.org/fx"
)

func CredentialsFxOption()fx.Option{
	return fx.Provide(func() {
		builder := credentialsgetter.Builder()
	})
}
