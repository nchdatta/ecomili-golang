package models

type Status string
type PlayPause string
type NewsSource string
type NewsStatus string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

const (
	Play  PlayPause = "play"
	Pause PlayPause = "pause"
)

const (
	Reuters NewsSource = "reuters"
	Local   NewsSource = "local"
)
const (
	Draft     NewsStatus = "draft"
	Published NewsStatus = "published"
	Archieved NewsStatus = "archieved"
)
