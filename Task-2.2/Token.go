package patterns

import "time"

type Token struct {
	date      time.Time
	sub       string
	tokenType string
}

type RequestData struct {
	cookie map[string]any
}

// json.Marshal(userData)
// if jsonErr := json.Unmarshal(body, registrationData); jsonErr != nil
func test() {
	cookie := map[string]any{
		"": nil,
	}
}
