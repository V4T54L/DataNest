package main

import (
	"backend/internals/schemas"
	"backend/internals/utils"
	"log"
)

func Main() {
	user := schemas.UserDetails{
		ID:       45,
		Username: "asdadwa",
		Email:    "asadaeref@afkabskj.acsian",
		Name:     "asufiablfa dalnfa dsfwadaw",
	}
	tokenStr, err := utils.GenerateToken(user)
	if err != nil {
		log.Fatal("Error generating token : ", err)
	}

	decodedUser, err := utils.VerifyToken(tokenStr)
	if err != nil {
		log.Fatal("Error parsing token : ", err)
	}

	log.Println("Decoded User : \n", decodedUser)
}
