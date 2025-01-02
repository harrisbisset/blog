package models

type (
	PublicUserEvent struct {
		Id        string `json:"id"`
		EventType string `json:"type"`
		Actor     struct {
			Id           int    `json:"id"`
			Login        string `json:"login"`
			DisplayLogin string `json:"display_login"`
			GravaterId   string `json:"gravater_id"`
			Url          string `json:"url"`
			AvatarUrl    string `json:"avatar_url"`
		} `json:"actor"`
		Repo struct {
			Id int `json:"id"`
		} `json:"repo"`
		Payload struct {
			RepositoryId int    `json:"repository_id"`
			PushId       int    `json:"push_id"`
			Size         int    `json:"size"`
			DistinctSize int    `json:"distinct_size"`
			Ref          string `json:"ref"`
			Head         string `json:"head"`
			Before       string `json:"before"`
			Commits      []struct {
				Sha    string
				Author struct {
					Email string `json:"email"`
					Name  string `json:"name"`
				} `json:"author"`
				Message  string `json:"message"`
				Distinct bool   `json:"distinct"`
				Url      string `json:"url"`
			} `json:"commits"`
		} `json:"payload"`
		Public    bool   `json:"public"`
		CreatedAt string `json:"created_at"`
	}
)
