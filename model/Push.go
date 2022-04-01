package model

import (
	"encoding/json"
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
)

var (
	VapidPrivateKey = "Ix-taUAI1eSi8S2pPk-zvLrmYS-JLNpF4fh4zLsGAqY"
	VapidPublicKey  = "BIZ20hSa0uDmIRkuw1A5nUl92alEUy8-hvpg7TCMNJ7L5CnMEuPkkd2MPwCSifjSivEjY_aMw3eDSx8MWnDh2hY"
	TTL             = 60
)

type Push struct {
	VapidPublicKey  string
	VapidPrivateKey string
	TTL             int
}

func (p *Push) InitFromVar() *Push {
	p.VapidPrivateKey = VapidPrivateKey
	p.VapidPublicKey = VapidPublicKey
	p.TTL = TTL
	return p
}
func (p *Push) SendToallsub(sub []webpush.Subscription, mail string, notif Notification) {
	if p.VapidPrivateKey == "" {
		p.InitFromVar()
	}
	for _, subscription := range sub {
		p.SendNotif(subscription, mail, notif)
	}

}
func (p *Push) SendNotif(sub webpush.Subscription, mail string, notif Notification) {
	parsed, err := json.Marshal(notif)
	if err != nil {
		fmt.Print("Erreur d'envoie de notification")
	}
	resp, err := webpush.SendNotification(parsed, &sub, &webpush.Options{
		Subscriber:      mail,
		VAPIDPublicKey:  p.VapidPublicKey,
		VAPIDPrivateKey: p.VapidPrivateKey,
		TTL:             p.TTL,
	})
	if err != nil {
		fmt.Print("Erreur d'envoie de notification")
	}
	defer resp.Body.Close()
}
