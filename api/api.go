package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"laji/v1/config"
	"laji/v1/models"
	"net/http"
)

type HttpTransport struct {
	Engine   *gin.Engine
	StopHttp chan bool
}

/**
 * @author lidong
 * @description 结构体赋值，类似PHP中实例化
 * @date 10:24 2021/9/9
 * @param
 * @return
 **/
func NewHttpTransport() *HttpTransport {
	return &HttpTransport{}
}

/**
 * @author lidong
 * @description http服务主方法
 * @date 10:20 2021/9/9
 * @param
 * @return
 **/
func (h *HttpTransport) HttpServer() *http.Server {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	h.Engine = gin.Default()
	h.ApiRoutes()
	srv := &http.Server{
		Addr:    cfg.HttpAddr,
		Handler: h.Engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.WithFields(log.Fields{
		"address": cfg.HttpAddr,
	}).Info("api: Running HTTP server")
	return srv
}

/**
 * @author lidong
 * @description 设置路由
 * @date 10:21 2021/9/9
 * @param
 * @return
 **/
func (h *HttpTransport) ApiRoutes() {
	h.Engine.GET("/ping", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{
			"code": 0,
			"info": "pong",
		})
	})

	v1 := h.Engine.Group("/v1")
	v1.GET("/login", h.Login)
	v1.POST("/create", h.createUser)
}

/**
 * @author lidong
 * @description 登录接口
 * @date 10:21 2021/9/9
 * @param
 * @return
 **/
func (h *HttpTransport) Login(ctx *gin.Context) {
	fmt.Println("hello world!!")
}

/**
 * @author lidong
 * @description 创建用户
 * @date 17:10 2021/9/9
 * @param
 * @return
 **/
func (h *HttpTransport) createUser(ctx *gin.Context) {
	userModel := models.NewUser()
	userModel.Username = "leedong1"
	userModel.Age = 12

	data, err := userModel.Create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"info": "create user failed!!",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"info": "create user success!!",
	})
}
