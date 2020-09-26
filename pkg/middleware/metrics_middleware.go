package middleware

import (
	"net/http"
	"strings"

	"github.com/hashicorp/go-hclog"
        //"github.com/sagargarg1/messageservice/pkg/data"
        "github.com/sagargarg1/messageservice/pkg/utils"
        "github.com/sagargarg1/messageservice/pkg/model"
)

func MetricsMiddleware(Logging hclog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			message := &model.Message{}

                	err := utils.FromJSON(message, r.Body)
                	if err != nil {
				Logging.Error("Failed to decode", "error", err)
				rw.WriteHeader(http.StatusBadRequest)
                		utils.ToJSON(&model.GenericError{Message: "Failed to deserialize"}, rw)
                		return
                	}

			//Only have handling for post but could be extended for delete and update
			if r.Method == http.MethodPost {
				Logging.Debug("Adding metrics \n")
				text := message.Text
				AddMetrics(text, Logging)
			}

			// Call the next handler
			next.ServeHTTP(rw, r)
		})
	}
}

//Function to add metrics
//TODO add functionality of delete and update to refresh the metrics
func AddMetrics(text string, Logging hclog.Logger) {

	str := strings.ToLower(text)
	
	if v, ok := model.Metrics["Number"]; ok {
		model.Metrics["Number"] = v + 1
	}

	if v, ok := model.Metrics["BirthdayMessages"]; ok {
		if strings.Contains(str, "birthday") {
			model.Metrics["BirthdayMessages"] = v + 1
		}
	}

	if v, ok := model.Metrics["GoodMorningMessages"]; ok {
		if strings.Contains(str, "morning") {
			model.Metrics["GoodMorningMessages"] = v + 1
		}
	}

	if v, ok := model.Metrics["SorryMessages"]; ok {
		if strings.Contains(str, "sorry") {
			model.Metrics["SorryMessages"] = v + 1
		}
	}

	if v, ok := model.Metrics["PalindromeMessages"]; ok {
		if utils.IsPalindrome(str) == true {
			Logging.Debug("Message is palindrome: %#s\n", str)
                	model.Metrics["PalindromeMessages"] = v + 1
		}
        }
}

