package main

import "fmt"

func main() {
	// Create a newNotificationBuilder
	var bldr = newNotificationBuilder()

	// Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.png")
	bldr.SetPriority(5)
	bldr.SetMessage("Basic Notification with some text in it")
	bldr.SetNotifType("alert")

	// Use the Build function
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification: ", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}
}
