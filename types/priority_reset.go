package types

var (
	// PriorityResetHeightInterval determines the interval at which the validator priority is reset.
	// If set to 100, the priority is reset every 100 blocks (e.g., at heights 100, 200, 300, ...).
	// When the chain reaches a height that is a multiple of this value, the validator priority is reset.
	PriorityResetHeightInterval int64 = 100
	// PriorityResetRoundInterval defines the interval at which the validator priority is reset.
	PriorityResetRoundInterval int32 = 20
)
