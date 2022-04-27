package main

import "fmt"

type SohneeTV struct {
	vol     int
	channel int
	isOn    bool
}

func (st *SohneeTV) turnOn() {
	fmt.Println("Sohnee TV is on")
	st.isOn = true
}

func (st *SohneeTV) turnOff() {
	fmt.Println("Sohnee TV is off")
	st.isOn = false
}

func (st *SohneeTV) volumeUp() int {
	st.vol++
	fmt.Println("Increasing Sohnee TV volume to", st.vol)
	return st.vol
}

func (st *SohneeTV) volumeDown() int {
	st.vol--
	fmt.Println("Decreasing Sohnee TV volume to", st.vol)
	return st.vol
}

func (st *SohneeTV) channelUp() int {
	st.channel++
	fmt.Println("Increasing Sohnee TV channel to", st.channel)
	return st.channel
}

func (st *SohneeTV) channelDown() int {
	st.channel--
	fmt.Println("Decreasing Sohnee TV channel to", st.channel)
	return st.channel
}

func (st *SohneeTV) goToChannel(ch int) {
	st.channel = ch
	fmt.Println("Setting Sohnee TV channel to", st.channel)
}
