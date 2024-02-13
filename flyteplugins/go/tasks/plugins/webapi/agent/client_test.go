package agent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeClients(t *testing.T) {
	cfg := defaultConfig
	cfg.Agents = map[string]*Agent{
		"x": {
			Endpoint: "x",
		},
		"y": {
			Endpoint: "y",
			IsSync:   true,
		},
	}
	ctx := context.Background()
	err := SetConfig(&cfg)
	assert.NoError(t, err)
	cs, err := initializeClients(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, cs)
	_, ok := cs.syncAgentClients["y"]
	assert.True(t, ok)
	_, ok = cs.asyncAgentClients["x"]
	assert.True(t, ok)
}
