package controllers

import (
	"encoding/json"
	"fmt"
	"testproject/models"

	"github.com/astaxie/beego"
	//"reflect"
	//"strings"
)

type HomeController struct {
	beego.Controller
}

type User struct {
	Id string
	//Pageno int
}

type FinalResponse struct {
	Code    int
	Message string
	Data    string
}

type ErrorResponse struct {
	Code    int
	Message string
}

func (h *HomeController) Post() {

	var ob User
	json.Unmarshal(h.Ctx.Input.RequestBody, &ob)
	key := ob.Id
	res, err := getRedisData(key)
	if err != nil {
		fmt.Println("Error Occured : ", err)
		r := ErrorResponse{404, "Error in getting data from redis"}
		h.Data["json"] = r
		h.ServeJSON()
	}

	r := FinalResponse{200, "success", res}
	h.Data["json"] = r
	h.ServeJSON()
}

func getRedisData(key string) (string, error) {
	res, err := models.RedisGet(key)
	if err != nil {
		return "", err
	}
	return res, nil
}
