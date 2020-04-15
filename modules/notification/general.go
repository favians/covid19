package notification

import (
	"fmt"
	"log"
)

func SendEmail(notif NotificationList, payload EmailPayload) (map[string]interface{}, error) {

	hedwigResponse, err := notif.SendEmail(payload, 0)
	if err != nil {
		return hedwigResponse, fmt.Errorf("error in sending notification")
	}

	log.Println(hedwigResponse)
	return hedwigResponse, nil
}

func SendSMS(notif NotificationList, payload SMSPayload) (map[string]interface{}, error) {

	hedwigResponse, err := notif.SendSMS(payload, 0)
	if err != nil {
		return hedwigResponse, fmt.Errorf("error in sending notification")
	}

	log.Println(hedwigResponse)
	return hedwigResponse, nil
}
