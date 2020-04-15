package notification

type NotificationList interface {
	Login(int64) (map[string]interface{}, error)
	JobInquiry(string, int64) (map[string]interface{}, error)
	SendEmail(EmailPayload, int64) (map[string]interface{}, error)
	SendSMS(SMSPayload, int64) (map[string]interface{}, error)
}
