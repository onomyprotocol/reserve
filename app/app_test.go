package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/reserve/app"
)

func TestGaiaApp_Export(t *testing.T) {
	app := app.Setup(t, false)
	_, err := app.ExportAppStateAndValidators(true, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
