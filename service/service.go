/**
所有服务开启或关闭再此调用
*/
package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"laji/v1/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//TODO
var service *Service

//TODO
type Service struct {
}

func init() {
	service = &Service{}
}

/**
 * @author lidong
 * @description 开启http服务
 * @date 10:20 2021/9/9
 * @param
 * @return
 **/
func Start() {
	httpTransport := api.NewHttpTransport()
	httpServer := httpTransport.HttpServer()
	handleSignals(httpServer)
}

/**
 * @author lidong
 * @description 信号捕捉，优雅退出
 * @date 10:19 2021/9/9
 * @param
 * @return
 **/
func handleSignals(httpServer *http.Server) {
	signCh := make(chan os.Signal, 2)
	signal.Notify(signCh, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
	exitCh := make(chan int)

	go func() {

		select {
		case s := <-signCh:
			stopHttpServer(httpServer)
			fmt.Println("捕捉到信号:", s)
			goto ExitProcess
		}

	ExitProcess:
		log.Println("Exit Service")
		exitCh <- 0
	}()

	code := <-exitCh
	os.Exit(code)
}

/**
 * @author lidong
 * @description 停止http服务
 * @date 10:20 2021/9/9
 * @param
 * @return
 **/
func stopHttpServer(httpServer *http.Server) {
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error("HttpServer Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Info("timeout of 2 seconds.")
	}
	log.Info("Server exiting")
}
