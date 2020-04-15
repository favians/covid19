package notification

func Init(client_id string, client_secret string, from string, host string, version string, test_mode bool) (Hedwig, error) {
	hedwig := Hedwig{
		ClientID:     client_id,
		CLientSecret: client_secret,
		Host:         host,
		Version:      version,
		TestMode:     test_mode,
		From:         from,
	}

	_, err := hedwig.Login(0)
	if err != nil {
		return hedwig, err
	}

	return hedwig, nil
}
