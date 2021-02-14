package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// PostKey defines stuff
	PostKey = "Post-value-"

	// PostCountKey defines stuff
	PostCountKey = "Post-count-"
)

// KeyPrefix does stuff
func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CommentKey= "Comment-value-"
	CommentCountKey= "Comment-count-"
)
