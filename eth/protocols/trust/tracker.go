package trust

import (
	"time"

	"github.com/Project-InitVerse/chain/p2p/tracker"
)

// requestTracker is a singleton tracker for request times.
var requestTracker = tracker.New(ProtocolName, time.Minute)
