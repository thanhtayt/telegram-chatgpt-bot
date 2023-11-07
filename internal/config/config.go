package config

import (
	"context"
	"sync"

	"github.com/iamwavecut/tool"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	TelegramAPIToken string `env:"BOT_TOKEN,required"`
	OpenAIToken      string `env:"OPENAI_TOKEN,required"`
	DefaultLanguage  string `env:"LANG,default=en"`
	ChatModel        string `env:"CHAT_MODEL,default=gpt-3.5-turbo-1106"`
}

var (
	once         sync.Once   //nolint:gochecknoglobals // desired behavior
	globalConfig = &Config{} //nolint:gochecknoglobals // desired behavior
)

func Get() Config {
	once.Do(func() {
		cfg := &Config{}
		tool.Must(envconfig.ProcessWith(context.Background(), cfg, envconfig.OsLookuper()))
		globalConfig = cfg
	})
	return *globalConfig
}
