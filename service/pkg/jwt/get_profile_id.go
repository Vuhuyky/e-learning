package jwtapp

import "github.com/gin-gonic/gin"

func GetProfileId(ctx *gin.Context) (uint, error) {
	token := GetToken(ctx)
	mapInfo, err := GetClaim(*token)

	if err != nil {
		return 0, err
	}

	profileId := uint(mapInfo["id"].(float64))

	return profileId, nil
}
