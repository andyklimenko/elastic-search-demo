package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	_, err := Load()
	require.EqualError(t, err, ErrNoElasticPassword.Error())

	require.NoError(t, os.Setenv("ELASTIC_USERNAME", "Luke"))
	require.NoError(t, os.Setenv("ELASTIC_PASSWORD", "I'm your father"))
	_, err = Load()
	require.EqualError(t, err, ErrNoElasticAddress.Error())

	require.NoError(t, os.Setenv("ELASTIC_ADDRESS", "google.com"))
	cfg, err := Load()
	require.NoError(t, err)

	assert.Equal(t, "Luke", cfg.Elastic.UserName)
	assert.Equal(t, "I'm your father", cfg.Elastic.Password)
	assert.Equal(t, "google.com", cfg.Elastic.Address)
}
