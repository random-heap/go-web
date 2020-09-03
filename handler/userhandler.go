package handler

import (
	"go-web/dao"
	"go-web/logger"
	"html/template"
	"net/http"
)

//Register 处理用户的函注册数
func Register(w http.ResponseWriter, r *http.Request) {

	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	logger.Info.Println("enter Register method with username = {}", username)

	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		dao.SaveUser(username, password, email)
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}

}

func CheckUserName(w http.ResponseWriter, r *http.Request) {

	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登录
		//去首页
		//GetPageBooksByPrice(w, r)
	}
}