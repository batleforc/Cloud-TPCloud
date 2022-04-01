package model

type Notification struct {
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	Icon    string   `json:"icon,omitempty"`
	Image   string   `json:"image,omitempty"`
	Vibrate []int    `json:"vibrate,omitempty"`
	Tag     string   `json:"tag,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Action struct {
	Action string `json:"action"`
	Title  string `json:"title"`
	Icon   string `json:"icon,omitempty"`
}

func DefaultNotification() Notification {
	return Notification{
		Title:   "",
		Body:    "",
		Icon:    "/favicon.ico",
		Vibrate: []int{200, 100, 200},
		Tag:     "pos",
	}
}
