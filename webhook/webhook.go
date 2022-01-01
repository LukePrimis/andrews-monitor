package webhook

import "fmt"

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Auth struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields"`
	Author      Auth    `json:"author"`
}

type Webhook struct {
	Content   string  `json:"content"`
	Embeds    []Embed `json:"embeds"`
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
}

func MakeWebhook(availableTimes []string) Webhook {
	fields := []Field{}

	for _, t := range availableTimes {
		newField := Field{
			Name:   t,
			Value:  "Available",
			Inline: true,
		}
		fields = append(fields, newField)
	}

	emb := Embed{
		Title:       "Andrews Time Slots Available!",
		Description: "Go quick! There are available time slots to order Andrews:",
		Color:       16750558,
		Fields:      fields,
		Author: Auth{
			Name:    "Luke Primis",
			URL:     "https://github.com/LukePrimis",
			IconURL: "https://pbs.twimg.com/profile_images/1291465390632112135/9Ewz8sWn_400x400.jpg",
		},
	}
	wh := Webhook{
		Content:   "@everyone",
		Embeds:    []Embed{emb},
		Username:  "Andrews Alerts",
		AvatarURL: "https://imgix.ranker.com/list_img_v2/1367/2681367/original/best-roy-mustang-quotes?w=817&h=427&fm=jpg&q=50&fit=crop",
	}

	return wh
}

func LoginWebhook() Webhook {
	emb := Embed{
		Title:       "Login Successful!",
		Description: "Andrews Alerts successfully logged in.",
		Color:       16750558,
		Author: Auth{
			Name:    "Luke Primis",
			URL:     "https://github.com/LukePrimis",
			IconURL: "https://pbs.twimg.com/profile_images/1291465390632112135/9Ewz8sWn_400x400.jpg",
		},
	}
	wh := Webhook{
		Embeds:    []Embed{emb},
		Username:  "Andrews Alerts",
		AvatarURL: "https://imgix.ranker.com/list_img_v2/1367/2681367/original/best-roy-mustang-quotes?w=817&h=427&fm=jpg&q=50&fit=crop",
	}
	return wh
}

func OOSWebhook(timeSlot string) Webhook {
	emb := Embed{
		Title:       "Time Slot Gone!",
		Description: fmt.Sprintf("Time slots for %s are sold out!", timeSlot),
		Color:       16750558,
		Author: Auth{
			Name:    "Luke Primis",
			URL:     "https://github.com/LukePrimis",
			IconURL: "https://pbs.twimg.com/profile_images/1291465390632112135/9Ewz8sWn_400x400.jpg",
		},
	}
	wh := Webhook{
		Embeds:    []Embed{emb},
		Username:  "Andrews Alerts",
		AvatarURL: "https://imgix.ranker.com/list_img_v2/1367/2681367/original/best-roy-mustang-quotes?w=817&h=427&fm=jpg&q=50&fit=crop",
	}
	return wh
}

func AllOOSWebhook() Webhook {
	emb := Embed{
		Title:       "All Time Slots Gone!",
		Description: "All time slots are sold out!",
		Color:       16750558,
		Author: Auth{
			Name:    "Luke Primis",
			URL:     "https://github.com/LukePrimis",
			IconURL: "https://pbs.twimg.com/profile_images/1291465390632112135/9Ewz8sWn_400x400.jpg",
		},
	}
	wh := Webhook{
		Embeds:    []Embed{emb},
		Username:  "Andrews Alerts",
		AvatarURL: "https://imgix.ranker.com/list_img_v2/1367/2681367/original/best-roy-mustang-quotes?w=817&h=427&fm=jpg&q=50&fit=crop",
	}
	return wh
}
