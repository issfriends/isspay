package config

import (
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/kelseyhightower/envconfig"
)

func initSecret() (*Secrets, error) {
	secrets := &Secrets{}

	if secretsFile := os.Getenv("SECRETS_FILE"); secretsFile != "" {
		viper.SetConfigFile(secretsFile)
		secretsType := os.Getenv("SECRETS_FILE_TYPE")
		if secretsType == "" {
			secretsType = "yaml"
		}
		viper.SetConfigType(secretsType)
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
		err = viper.Unmarshal(secrets, func(decode *mapstructure.DecoderConfig) {
			decode.TagName = "yaml"
		})
		if err != nil {
			return nil, err
		}
		return secrets, nil
	}

	secrets.Linebot = &chatbot.Config{
		AccessToken: os.Getenv("LINEBOT_ACCESS_TOKEN"),
		Secret:      os.Getenv("LINEBOT_ACCESS_Secret"),
	}

	err := envconfig.Process("ISSPAY", secrets)
	if err != nil {
		return nil, err
	}

	return secrets, nil
}

// Key base64 encoded secret key
type Key string

// Secrets secret config
type Secrets struct {
	Linebot     *chatbot.Config `yaml:"linebot"`
	TokenSecret Key             `yaml:"token_secret"`
}
