package config

import "context"

// Config is a config :)
type Treads struct {
	List map[string]context.CancelFunc
}

var (
	treads Treads
)

// Get reads config from environment
func GetTreads() *Treads {
	once.Do(func() {
		treads.List = make(map[string]context.CancelFunc)
	})
	return &treads
}
