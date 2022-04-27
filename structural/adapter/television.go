package main

type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
	turnOn() int
	turnOff() int
	goToChannel(ch int)
}
