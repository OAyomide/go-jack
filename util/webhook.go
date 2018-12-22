package util

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		webhookVerify(w, r)
	}
}

//WebhookVerify allows us to return a function used to verify the webhook request
func WebhookVerify(token string) func(w http.ResponseWriter, r *http.Request) {
	hubChallenge := r.URL.Query().Get("hub.challenge")
	if token == r.URL.Query().Get("hub.verify_token") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hubChallenge))
		return
	} else {
		fmt.Fprint(w, "Nay! Tokens don't match")
	}
}
