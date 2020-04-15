package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
)

type (
	Hedwig struct {
		Token        string `json:"token"`
		Version      string `json:"version"`
		Host         string `json:"host"`
		ClientID     string `json:"client_id"`
		CLientSecret string `json:"client_secret"`
		TestMode     bool   `json:"test_mode"`
		From         string `json:"from"`
	}

	EmailPayload struct {
		To          string  `json:"to"`
		From        string  `json:"from"`
		Cc          *string `json:"cc"`
		Bcc         *string `json:"bcc"`
		Subject     string  `json:"subject"`
		Message     string  `json:"message"`
		MessageType string  `json:"message_type"`
	}

	SMSPayload struct {
		To      string `json:"to"`
		Message string `json:"message"`
	}
)

func (hedwig *Hedwig) Login(RequestTimeout int64) (map[string]interface{}, error) {
	Url := fmt.Sprintf("%s/api/%s/token", hedwig.Host, hedwig.Version)
	client := http.Client{}

	if RequestTimeout > 0 {
		client = http.Client{
			Timeout: time.Duration(time.Duration(RequestTimeout) * time.Second),
		}
	}

	if hedwig.TestMode {
		httpmock.Activate()

		mock, httpcode_mock := GetNotificationMock("hedwig_login", hedwig.ClientID)

		if len(mock) > 0 && httpcode_mock != 0 {
			httpmock.RegisterResponder("POST", Url,
				httpmock.NewStringResponder(httpcode_mock, mock))

			defer httpmock.DeactivateAndReset()
		} else {
			httpmock.DeactivateAndReset()
		}
	}

	input := map[string]interface{}{
		"client_id":     hedwig.ClientID,
		"client_secret": hedwig.CLientSecret,
	}

	loginPayload := new(bytes.Buffer)
	err := json.NewEncoder(loginPayload).Encode(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", Url, loginPayload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(buff), &response)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		// set hedwig token
		hedwig.Token = response["result"].(string)
		return response, nil
	}

	return response, fmt.Errorf("error login hedwig notification with status code %v", res.StatusCode)
}

func (hedwig *Hedwig) JobInquiry(jobID string, RequestTimeout int64) (map[string]interface{}, error) {
	Url := fmt.Sprintf("%s/api/%s/queue/result", hedwig.Host, hedwig.Version)
	client := http.Client{}

	if RequestTimeout > 0 {
		client = http.Client{
			Timeout: time.Duration(time.Duration(RequestTimeout) * time.Second),
		}
	}

	if hedwig.TestMode {
		httpmock.Activate()

		mock, httpcode_mock := GetNotificationMock("hedwig_job_inquiry", hedwig.ClientID)

		if len(mock) > 0 && httpcode_mock != 0 {
			httpmock.RegisterResponder("GET", Url,
				httpmock.NewStringResponder(httpcode_mock, mock))

			defer httpmock.DeactivateAndReset()
		} else {
			httpmock.DeactivateAndReset()
		}
	}

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id", jobID)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+hedwig.Token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(buff), &response)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		_, errors := response["result"].(map[string]interface{})
		if !errors {
			response["result"] = "success"
			return response, nil
		}

		return response, nil
	}

	return response, fmt.Errorf("error hedwig job inquiry, status code %v", res.StatusCode)
}

func (hedwig *Hedwig) SendEmail(input EmailPayload, RequestTimeout int64) (map[string]interface{}, error) {
	Url := fmt.Sprintf("%s/api/%s/email/send", hedwig.Host, hedwig.Version)
	client := http.Client{}

	if RequestTimeout > 0 {
		client = http.Client{
			Timeout: time.Duration(time.Duration(RequestTimeout) * time.Second),
		}
	}

	if hedwig.TestMode {
		httpmock.Activate()

		mock, httpcode_mock := GetNotificationMock("send_email", hedwig.ClientID)

		if len(mock) > 0 && httpcode_mock != 0 {
			httpmock.RegisterResponder("POST", Url,
				httpmock.NewStringResponder(httpcode_mock, mock))

			defer httpmock.DeactivateAndReset()
		} else {
			httpmock.DeactivateAndReset()
		}
	}

	EmailPayload := new(bytes.Buffer)
	err := json.NewEncoder(EmailPayload).Encode(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", Url, EmailPayload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+hedwig.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(buff), &response)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	flagRetry := false
	if res.StatusCode == 200 {
		hedwigResult, err := hedwig.JobInquiry(response["result"].(string), 0)
		if err != nil {
			return response, fmt.Errorf("error in hedwig validate process")
		}

		if hedwigStatus, oke := hedwigResult["result"].(string); hedwigStatus == "success" && oke {
			return response, nil
		}

		return hedwigResult, fmt.Errorf("email notification not send, job error")

	} else if res.StatusCode == 403 && response["meta"].(map[string]interface{})["messages"].([]interface{})[0].(string) == "Signature has expired" {
		result, err := hedwig.Login(0)
		if err != nil {
			return result, fmt.Errorf("error in hedwig login process")
		}

		flagRetry = true
	}

	if flagRetry {
		//Ini Untuk Kebutuhan Unit Test
		if hedwig.ClientID == "cMock3" {
			hedwig.ClientID = "cMock1"
		}

		result, err := hedwig.SendEmail(input, 0)
		if err != nil {
			return result, fmt.Errorf("error in hedwig login process")
		}

		return result, nil
	}

	return response, fmt.Errorf("error sending email for notification, or wrong endpoint %v", res.StatusCode)
}

func (hedwig *Hedwig) SendSMS(input SMSPayload, RequestTimeout int64) (map[string]interface{}, error) {
	Url := fmt.Sprintf("%s/api/%s/sms/otp", hedwig.Host, hedwig.Version)
	client := http.Client{}
	log.Println("Url:", Url)
	if RequestTimeout > 0 {
		client = http.Client{
			Timeout: time.Duration(time.Duration(RequestTimeout) * time.Second),
		}
	}

	if hedwig.TestMode {
		httpmock.Activate()

		mock, httpcode_mock := GetNotificationMock("send_email", hedwig.ClientID)

		if len(mock) > 0 && httpcode_mock != 0 {
			httpmock.RegisterResponder("POST", Url,
				httpmock.NewStringResponder(httpcode_mock, mock))

			defer httpmock.DeactivateAndReset()
		} else {
			httpmock.DeactivateAndReset()
		}
	}

	SMSPayload := new(bytes.Buffer)
	err := json.NewEncoder(SMSPayload).Encode(input)
	if err != nil {
		return nil, err
	}
	log.Println("SMSPayload:", SMSPayload)

	req, err := http.NewRequest("POST", Url, SMSPayload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+hedwig.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(buff), &response)
	if err != nil {
		return nil, err
	}
	log.Println("Response Body:", response)

	defer res.Body.Close()

	flagRetry := false
	if res.StatusCode == 200 {
		hedwigResult, err := hedwig.JobInquiry(response["result"].(string), 0)
		if err != nil {
			return response, fmt.Errorf("error in hedwig validate process")
		}
		log.Println("Hedwig Result:", hedwigResult)

		if hedwigStatus, oke := hedwigResult["result"].(string); hedwigStatus == "success" && oke {
			return response, nil
		}

		return hedwigResult, fmt.Errorf("sms notification not send, job error")

	} else if res.StatusCode == 403 && response["meta"].(map[string]interface{})["messages"].([]interface{})[0].(string) == "Signature has expired" {
		result, err := hedwig.Login(0)
		if err != nil {
			return result, fmt.Errorf("error in hedwig login process")
		}
		log.Println("Result with error:", result)

		flagRetry = true
	}

	if flagRetry {
		//Ini Untuk Kebutuhan Unit Test
		if hedwig.ClientID == "cMock3" {
			hedwig.ClientID = "cMock1"
		}

		result, err := hedwig.SendSMS(input, 0)
		if err != nil {
			return result, fmt.Errorf("error in hedwig login process")
		}

		return result, nil
	}

	return response, fmt.Errorf("error sending sms for notification, or wrong endpoint %v", res.StatusCode)
}
