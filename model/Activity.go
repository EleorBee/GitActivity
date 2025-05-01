package model

import "time"

type Activity struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		Id           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarId   string `json:"gravatar_id"`
		Url          string `json:"url"`
		AvatarUrl    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
