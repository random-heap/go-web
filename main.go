package main

import (
	handler "go-web/handler"
	"go-web/logger"
	"net/http"
)

func main() {

	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//去首页
	http.HandleFunc("/main", handler.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", handler.Login)
	//去注销
	http.HandleFunc("/logout", handler.Logout)
	//去注册
	http.HandleFunc("/regist", handler.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", handler.CheckUserName)
	//获取所有图书
	// http.HandleFunc("/getBooks", controller.GetBooks)
	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", handler.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", handler.GetPageBooksByPrice)
	//添加图书
	// http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", handler.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", handler.ToUpdateBookPage)
	//更新或添加图书
	http.HandleFunc("/updateOraddBook", handler.UpdateOrAddBook)
	//添加图书到购物车中
	http.HandleFunc("/addBook2Cart", handler.AddBook2Cart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", handler.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", handler.DeleteCart)
	//删除购物项
	http.HandleFunc("/deleteCartItem", handler.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem", handler.UpdateCartItem)
	//去结账
	http.HandleFunc("/checkout", handler.Checkout)
	//获取所有订单
	http.HandleFunc("/getOrders", handler.GetOrders)
	//获取订单详情，即订单所对应的所有的订单项
	http.HandleFunc("/getOrderInfo", handler.GetOrderInfo)
	//获取我的订单
	http.HandleFunc("/getMyOrder", handler.GetMyOrders)
	//发货
	http.HandleFunc("/sendOrder", handler.SendOrder)
	//确认收货
	http.HandleFunc("/takeOrder", handler.TakeOrder)

	logger.Info.Println("server starting at http://localhost:8080")

	logger.Error.Fatal(http.ListenAndServe(":8080", nil))

}
