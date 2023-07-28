package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Application struct {
	logger    *zap.SugaredLogger
	ginEngine *gin.Engine
}

func newGinEngine() *gin.Engine {
	return gin.New()
}

func newApplication(logger *zap.SugaredLogger, ginEngine *gin.Engine) Application {
	return Application{logger, ginEngine}
}

func newSugaredLogger(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}

func newLogger() (*zap.Logger, error) {
	var zapCfg zap.Config

	zapCfg = zap.NewDevelopmentConfig()

	return zapCfg.Build()
}

func (app Application) Runserver() {

	// 確保Zap日誌訊息被寫入目的地(defer)
	defer app.logger.Sync()

	// 創建一個 Gin 的預設引擎
	router := app.ginEngine

	// 添加 Zap 中間件，將 Gin 請求日誌輸出到 Zap 日誌庫中
	zepLogger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("初始化 Zap 日誌庫失敗:", err)
		return
	}
	router.Use(ginzap.Ginzap(zepLogger, time.RFC3339, true))

	// 設置一個路由處理器，回傳一個簡單的訊息
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	// 添加 pprof 中間件，進行性能分析
	pprof.Register(router)

	// 創建一個 HTTP 伺服器
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 開始 HTTP 伺服器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("HTTP 伺服器啟動失敗:", err)
		}
	}()

	// 註冊 SIGINT 信號處理器
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	// 阻塞，等待接收 SIGINT 信號
	fmt.Println("程式已啟動，按下 Ctrl+C 可以終止程式...")
	<-sigChan

	// 收到 SIGINT 信號後先關閉 HTTP 伺服器
	fmt.Println("\n接收到 SIGINT 信號，正在關閉 HTTP 伺服器...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("關閉 HTTP 伺服器失敗:", err)
	}

	// 在這裡可以進行其他的清理工作，例如關閉資料庫連接等
	fmt.Println("正在執行其他清理工作...")
	// 進行一些模擬清理工作
	time.Sleep(2 * time.Second)

	// 正確地終止程式
	fmt.Println("程式已結束")
	os.Exit(0)
}
