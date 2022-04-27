package main

import "fmt"

type NotificationBuilder struct {
	Title     string
	Subtitle  string
	Message   string
	Image     string
	Icon      string
	Priority  int
	NotifType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subTitle string) {
	nb.Subtitle = subTitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotifType(notifType string) {
	nb.NotifType = notifType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	// Error checking
	if nb.Title == "" {
		return nil, fmt.Errorf("You need to provide a title")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("Priority must be 0 to 5")
	}

	// Return a newly created Notification
	return &Notification{
		title:     nb.Title,
		subTitle:  nb.Subtitle,
		message:   nb.Message,
		image:     nb.Image,
		icon:      nb.Icon,
		priority:  nb.Priority,
		notifType: nb.NotifType,
	}, nil
}
