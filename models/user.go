package models

import "time"

type User struct {
	UserID    int       `json:"user_id"`
	GoogleID  string    `json:"google_id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUser struct {
	UserName string `json:"user_name"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type GoogleOAuthCode struct {
	Code string `json:"code"`
}

type GoogleOAuthToken struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type GoogleUserDataResponse struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Hd            string `json:"hd"`
	Id            string `json:"id"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
