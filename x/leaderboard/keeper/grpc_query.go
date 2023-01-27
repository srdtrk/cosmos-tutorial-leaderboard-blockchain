package keeper

import (
	"github.com/srdtrk/leaderboard/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
