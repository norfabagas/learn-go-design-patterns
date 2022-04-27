package main

// each property starts with lower case to make it immutable field
type Notification struct {
	title     string
	subTitle  string
	message   string
	image     string
	icon      string
	priority  int
	notifType string
}
