package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllFootballers(ctx *gin.Context) {
	fmt.Println(" Reached GetAllFootballers")
	result := DBGetAllPlayers()
	fmt.Println(" result", result)

	ctx.JSON(http.StatusOK, result)
}

func PostNewFootballer(ctx *gin.Context) {
	fmt.Println(" Reached DBAddPlayers")
	var newPlayer Footballer

	if err := ctx.BindJSON(&newPlayer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := DBAddPlayer(newPlayer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusCreated, "")
}

func DeleteFootballer(dfb *gin.Context) {
	fmt.Println("reached DBDeleteByUUID")
	UUID := dfb.Param("UUID")
	fmt.Println("UUID :", UUID)
	ppp, _ := strconv.Atoi(UUID)
	DBDeleteByUUID(ppp)
	dfb.Status(http.StatusNoContent)
}

func UpdateFootballerByUuid(ctx *gin.Context) {
	fmt.Println("reached DBUpdateByUUID")
	var update Footballer

	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	UUID := ctx.Param("UUID")
	fmt.Println("UUID:", UUID)
	trd, _ := strconv.Atoi(UUID)
	err := DBUpdateByUuid(trd, update.Name)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusNotAcceptable, gin.H{
		"error": "ID NOT FOUND",
	})

}

func SearchFootballerAny(ctx *gin.Context) {
	key := ctx.Param("key")
	fmt.Println(key)
	var result []Footballer
	for _, player := range footballers {
		if strings.Contains(strings.ToUpper(player.Name), strings.ToUpper(key)) || strings.Contains(strings.ToUpper(player.Club), strings.ToUpper(key)) || strings.Contains(strings.ToUpper(player.Country), strings.ToUpper(key)) {
			result = append(result, player)
		}
	}

	ctx.JSON(http.StatusOK, result)
}
