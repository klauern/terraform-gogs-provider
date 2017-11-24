package gogs

import (
	gogs "github.com/gogits/go-gogs-client"
)

// Config represents the configuration necessary to connect to a Gogs API Server.
type Config struct {
	Token   string
	BaseURL string
}

// Client will create a new Gogs HTTP client based on the Config.
func (c *Config) Client() *gogs.Client {
	return gogs.NewClient(c.BaseURL, c.Token)
}
