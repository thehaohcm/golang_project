package models

type Relationship struct {
	Requestor      string
	Target         string
	IsFriend       bool
	FriendBlocked  bool
	Subscribed     bool
	SubscribeBlock bool
}
