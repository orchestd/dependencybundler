package credentials

type Builder interface {
	DevMode() Builder
	Build() (CredentialsGetter, error)
}

type CredentialsGetter interface {
	GetCredentials() Credentials
	Implementation() interface{}
}
type Credentials struct {
	DbUsername string `envconfig:"DB_USERNAME"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbHost     string `envconfig:"DB_HOST"`
	DbName     string `envconfig:"DB_NAME"`
}
