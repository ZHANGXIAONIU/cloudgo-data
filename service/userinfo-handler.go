package service

import (
	"net/http"
	"strconv"

	"github.com/ajian/cloudgo-data/entities"
	"github.com/go-xorm/xorm"

	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		u.DepartName = req.Form["departname"][0]
		//entities.UserInfoService.Save(u)
		mySQLEngine, _ := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
		_, err := mySQLEngine.Table("userinfo").Insert(u)
		if err != nil {
			panic(err)
		}
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		mySQLEngine, _ := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
		if len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			//u := entities.UserInfoService.FindByID(int(i))

			var u entities.UserInfo
			mySQLEngine.Table("userinfo").Where("uid = ?", int(i)).Get(&u)
			formatter.JSON(w, http.StatusOK, u)
			return
		}
		//ulist := entities.UserInfoService.FindAll()
		ulist := make([]entities.UserInfo, 0)
		mySQLEngine.Table("userinfo").Find(&ulist)
		formatter.JSON(w, http.StatusOK, ulist)
	}
}
