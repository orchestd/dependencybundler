package credentials

import (
	"bitbucket.org/HeilaSystems/configurations/credentials"
	credentialsConstructor "bitbucket.org/HeilaSystems/dependencybundler/interfaces/credentials"
	"os"
)

func DefaultCredentials(builder credentials.Builder) (credentialsConstructor.CredentialsGetter,error) {
	env, isExist := os.LookupEnv("ENABLE_SECRET_MANAGER")
	if isExist && env == "true" {
		pID, _ := os.LookupEnv("PROJECT_ID")
		version , _ := os.LookupEnv("SECRET_MANAGER_VERSION")
		secretName , _ := os.LookupEnv("SECRET_NAME")
		builder = builder.UseGcpSecretManager(pID).SetSecretManagerVersion(version).SetSecretName(secretName)
	}
	return builder.Build()
}
