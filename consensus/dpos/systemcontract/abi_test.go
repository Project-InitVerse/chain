package systemcontract

import (
	"github.com/Project-InitVerse/chain/accounts/abi"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestJsonUnmarshalABI(t *testing.T) {
	for _, abiStr := range []string{DposFactoryInteractiveABI, PunishInteractiveABI, SysGovInteractiveABI, AddrListInteractiveABI} {
		_, err := abi.JSON(strings.NewReader(DposFactoryInteractiveABI))
		require.NoError(t, err, abiStr)
	}
}
