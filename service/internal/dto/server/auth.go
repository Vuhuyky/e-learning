package responsedata

import (
	"app/internal/entity"
)

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	AccessToken  string         `json:"accessToken"`
	RefreshToken string         `json:"refreshToken"`
	Profile      entity.Profile `json:"profile"`
}

type RefreshTokenResponse struct {
	AccessToken  string         `json:"accessToken"`
	RefreshToken string         `json:"refreshToken"`
	Profile      entity.Profile `json:"profile"`
}
