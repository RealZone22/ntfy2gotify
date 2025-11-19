package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"ntfy2gotify/pkg/utils"
	"path"

	"github.com/labstack/echo/v4"
)

func HandleNtfyRequests(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(400, "Invalid request body")
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		data = make(map[string]interface{})
		data["message"] = string(body)
		utils.Logger.Debug().Msgf("Parsed as plain text, message: %s\n", data["message"])
	}

	topic := c.Param("topic")
	if topic == "" {
		if topicValue, ok := data["topic"]; ok {
			topic = topicValue.(string)
		} else {
			return c.String(400, "Topic is required")
		}
	}

	gotifyUrl := ""
	for entry, gotify := range utils.Config.Subscriptions {
		u, err := url.Parse(entry)
		if err != nil {
			continue
		}
		if path.Base(u.Path) == topic {
			gotifyUrl = gotify
			break
		}
	}
	if gotifyUrl != "" {
		req, _ := http.NewRequest("POST", gotifyUrl, nil)
		q := req.URL.Query()
		if title, ok := data["title"].(string); ok {
			q.Add("title", title)
		} else {
			q.Add("title", topic)
		}
		if message, ok := data["message"].(string); ok {
			q.Add("message", message)
		} else {
			return c.String(400, "Message is required")
		}
		req.URL.RawQuery = q.Encode()
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			utils.Logger.Error().Err(err).Msg("Error sending ntfy message to gotify")
			return c.String(500, "Error sending ntfy message to gotify")
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				utils.Logger.Error().Err(err).Msg("Error closing response body")
			}
		}(resp.Body)
		return c.String(resp.StatusCode, "Message forwarded to gotify")
	}
	return c.String(404, "No subscription found for this topic")
}
