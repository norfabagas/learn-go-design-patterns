package main

type SammysangAdapter struct {
	sstv *SammysangTV
}

func (ss *SammysangAdapter) turnOn() {
	ss.sstv.setOnState(true)
}

func (ss *SammysangAdapter) turnOff() {
	ss.sstv.setOnState(false)
}

func (ss *SammysangAdapter) volumeUp() int {
	vol := ss.sstv.getVolume() + 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *SammysangAdapter) volumeDown() int {
	vol := ss.sstv.getVolume() - 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *SammysangAdapter) channelUp() int {
	ch := ss.sstv.getChannel() + 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *SammysangAdapter) channelDown() int {
	ch := ss.sstv.getChannel() - 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *SammysangAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
}
