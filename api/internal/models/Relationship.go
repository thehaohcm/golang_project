package models

type Relationship struct {
	Requestor       string
	Target          string
	Is_friend       bool
	Friend_blocked  bool
	Subscribed      bool
	Subscribe_block bool
}
