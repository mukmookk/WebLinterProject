package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
)

func get(ctx *gin.Context) {
	data, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(data)
	var coffees []Coffee
	json.Unmarshal(byteValue, &coffees)
	ctx.JSON(http.StatusOK, coffees)
}

func post(ctx *gin.Context) {
	coffee := Coffee{}
	if err := ctx.BindJSON(&coffee); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	coffees := []Coffee{}
	data, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(data)
	json.Unmarshal(byteValue, &coffees)
	coffees = append(coffees, coffee)
	coffeesByte, _ := json.MarshalIndent(coffees, "", "    ")
	os.WriteFile("sample.json", coffeesByte, 0644)
	ctx.JSON(http.StatusOK, coffee)
}

func put(ctx *gin.Context) {
	index, _ := strconv.ParseInt(ctx.Query("index"), 10, 32)
	coffee := Coffee{}
	if err := ctx.BindJSON(&coffee); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	data, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(data)
	var coffees []Coffee
	json.Unmarshal(byteValue, &coffees)
	coffees[index] = coffee
	coffeesByte, _ := json.MarshalIndent(coffees, "", "    ")
	os.WriteFile("sample.json", coffeesByte, 0644)
	ctx.JSON(http.StatusOK, coffee)
}

func delete(ctx *gin.Context) {
	index, _ := strconv.ParseInt(ctx.Query("index"), 10, 32)
	data, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(data)
	var coffees []Coffee
	json.Unmarshal(byteValue, &coffees)
	coffees = append(coffees[:index], coffees[index+1:]...)
	coffeesByte, _ := json.MarshalIndent(coffees, "", "    ")
	os.WriteFile("sample.json", coffeesByte, 0644)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
