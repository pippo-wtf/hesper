package agents

import (
	"context"
	"encoding/json"
)

// Devices (internal/host removal.go) serves devices.list and
// devices.remove: every device of the owner, and removing one from the
// relay and this Mac. Only a person calls them (treePerson). Nil: the
// methods are unavailable.
type Devices interface {
	Call(ctx context.Context, method string, params json.RawMessage) (any, error)
}
