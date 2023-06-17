package configmanager

import (
	"fmt"
	"github.com/kardolus/chatgpt-cli/config"
	"github.com/kardolus/chatgpt-cli/types"
)

type ConfigManager struct {
	configStore config.ConfigStore
	Config      types.Config
}

func New(cs config.ConfigStore) (*ConfigManager, error) {
	c, err := cs.ReadDefaults()
	if err != nil {
		return nil, fmt.Errorf("error reading default values: %w", err)
	}

	configured, err := cs.Read()
	if err == nil {
		if configured.Model != "" {
			c.Model = configured.Model
		}
		if configured.MaxTokens != 0 {
			c.MaxTokens = configured.MaxTokens
		}
		if configured.URL != "" {
			c.URL = configured.URL
		}
		if configured.CompletionsPath != "" {
			c.CompletionsPath = configured.CompletionsPath
		}
		if configured.ModelsPath != "" {
			c.ModelsPath = configured.ModelsPath
		}
	}

	return &ConfigManager{configStore: cs, Config: c}, nil
}

func (c *ConfigManager) WriteModel(model string) error {
	c.Config.Model = model

	return c.configStore.Write(c.Config)
}
