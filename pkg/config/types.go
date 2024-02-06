package config

const (
	Application = "multipliers-api"
	Version     = "v0.1.0"
)

type Config struct {
	Host                 *string `env:"HOST,default=localhost"`
	HttpPort             *string `env:"HTTP_PORT,default=8080"`
	TokenSignatureKey    *string `env:"TOKEN_SIGNATURE_KEY,default=SecretYouShouldHide"`
	TokenVerificationKey *string `env:"TOKEN_VERIFICATION_KEY,default=SecretYouShouldHide"`
	TokenTimeout         *string `env:"TOKEN_TIMEOUT,default=24h"`
}
