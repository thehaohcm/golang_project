package models

// Relationship struct used when mapping to get a Relationship model after querying data from Relationship table in database
type Relationship struct {
	Requestor      string
	Target         string
	IsFriend       bool
	FriendBlocked  bool
	Subscribed     bool
	SubscribeBlock bool
}
