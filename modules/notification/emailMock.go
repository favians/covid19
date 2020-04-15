package notification

import (
	"net/http"
)

var HedwiginternalServerError = `{
	"errors": {
		"messages": [
			"This is error",
			"500"
		],
		"traces": [
			"  File \"/your/path/app.py\", line 1612, in full_dispatch_request\n    rv = self.dispatch_request()\n",
			"  File \"/your/path/app.py\", line 1598, in dispatch_request\n    return self.view_functions[rule.endpoint](**req.view_args)\n",
			"  File \"/your/path/__init__.py\", line 477, in wrapper\n    resp = resource(*args, **kwargs)\n",
			"  File \"/your/path/views.py\", line 84, in view\n    return self.dispatch_request(*args, **kwargs)\n",
			"  File \"/your/path/__init__.py\", line 587, in dispatch_request\n    resp = meth(*args, **kwargs)\n",
			"  File \"/your/path/views.py\", line 28, in post\n    raise Exception('This is error', 500)\n"
		],
		"type": "Exception"
	}
}`

func GetNotificationMock(service string, identifier interface{}) (string, int) {
	switch service {
	default:
		return "", 0
	case "hedwig_login":
		if identifier.(string) == "cMock3" {
			identifier = "cMock1"
		}
		return hedwigLogin(identifier.(string))
	case "hedwig_job_inquiry":
		return hedwigJobInquiry(identifier.(string))
	case "send_email":
		return hedwigSendEmail(identifier.(string))
	}
}

func hedwigLogin(identifier string) (string, int) {
	switch identifier {
	default:
		return "", 0
	case "cMock1":
		return `{
			"meta": {
				"messages": [
				  "valid"
				],
				"status": true
			  },
			  "result": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1NjUyNDQ0MjMsIm5iZiI6MTU2NTE1ODAyMywiZnJlc2giOmZhbHNlLCJqdGkiOiJhMDdjOTJiYy1kZmQ3LTRhMWMtYjIxNy0yOWM2MGYzZjQzY2YiLCJpYXQiOjE1NjUxNTgwMjMsImlkZW50aXR5IjoiTUFMTEFDQSIsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJjbGllbnRfaWQiOiJNQUxMQUNBIiwiZmNtX3NlcnZlcl9rZXkiOiJub25lIn19.sbddeK3YerF0hQqltZRb0D0IUryPstxFgnU53OW7jM4"
		}`, http.StatusOK
	case "cMock2":
		return HedwiginternalServerError, http.StatusInternalServerError
	}
}

func hedwigJobInquiry(identifier string) (string, int) {
	switch identifier {
	default:
		return "", 0
	case "cMock1":
		return `{
			"meta": {
			  "messages": [
				"valid"
			  ],
			  "status": true
			},
			"result": null
		  }`, http.StatusOK
	case "cMock2":
		return HedwiginternalServerError, http.StatusInternalServerError
	}
}

func hedwigSendEmail(identifier string) (string, int) {
	switch identifier {
	default:
		return "", 0
	case "cMock1":
		return `{
			"meta": {
				"messages": [
					"valid"
				],
				"status": true
			},
			"result": "c378bc3f-e682-4638-9a0e-1eb25ae1c563"
		}`, http.StatusOK
	case "cMock2":
		return HedwiginternalServerError, http.StatusInternalServerError
	case "cMock3":
		return `{
			"meta": {
				"messages": [
				  "Signature has expired"
				],
				"status": false
			  },
			  "result": null
		}`, http.StatusForbidden
	}
}
