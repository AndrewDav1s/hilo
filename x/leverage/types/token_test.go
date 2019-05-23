package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/cicizeo/hilo/x/leverage/types"
)

func TestUTokenFromTokenDenom(t *testing.T) {
	tokenDenom := "uhilo"
	uTokenDenom := types.UTokenFromTokenDenom(tokenDenom)
	require.Equal(t, "u/"+tokenDenom, uTokenDenom)
	require.NoError(t, sdk.ValidateDenom(uTokenDenom))
}
