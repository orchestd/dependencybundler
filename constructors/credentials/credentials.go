package credentials

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/credentials"
	"os"
)

func DefaultCredentials(builder credentials.Builder) credentials.CredentialsGetter {
	env, isExist := os.LookupEnv("HEILA_ENV")
	if !isExist {
		panic("Cannot initialize new credentials, missing environment variable HEILA_ENV ")
	}
	if env == "LOCAL" || env == "DEV" {
		builder.DevMode()
	}
	if creds , err :=  builder.Build();err != nil {
		panic(err)
	}else{
		return creds
	}
}
