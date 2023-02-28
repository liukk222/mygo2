package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

/* <!-- 将图片与音频文件放于assets文件夹下,，如果没有则新建一个  -->
   <!-- 将html文件放于templates下，如果没有则新建一个 --> */

func game() string {
	rand.Seed(time.Now().UnixNano())
	q := rand.Int63n(10) + 1
	a := "bug"
	switch q {
	case 1:
		a = "哦，你连女神家的一条狗都不如"
	case 2:
		a = "哦，原来你是备胎"
	case 3:
		a = "哦，很抱歉女神想不起来你是谁了"
	case 4:
		a = "哦，原来你是女神的提款机"
	case 5:
		a = "哦，在女神心中你只是一条单身狗"
	case 6:
		a = "哦，在女神心中你是一个好人"
	case 7:
		a = "哦，对不起女神有点讨厌你每天来烦她"
	case 8:
		a = "哦，对不起女神的心中还有另一个他"
	case 9:
		a = "哦，女神从来没有爱过你"
	case 10:
		a = "哦, 其实不用再想了，你自己心里也知道你就是一个小丑"
	default:
		fmt.Println("输入错误，请重新输入")
		a := "bug2"
		return a
	}

	return a
}

func Login(c *gin.Context) {
	c.HTML(200, "test.html", nil)
}
func Regsiter1(c *gin.Context) {
	nr := game()
	fmt.Println(nr)
	c.HTML(200, "test1.html", gin.H{
		"Nra": nr,
	})

}
func Regsiter2(c *gin.Context) {
	nr := game()
	c.HTML(200, "test2.html", gin.H{
		"Nre": nr,
	})
}
func Regsiter3(c *gin.Context) {
	c.HTML(200, "test3.html", nil)
}
func Regsiter4(c *gin.Context) {
	c.HTML(200, "test4.html", nil)
}
func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")
	e.GET("/hello", Login)
	e.POST("/hello1", Regsiter1)
	e.POST("/hello2", Regsiter2)
	e.POST("/hello3", Regsiter3)
	e.POST("/hello4", Regsiter4)
	e.Run(":8888")

}
