package pprof

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof" //init
	"time"

	"family-joint-school/common/app"

	"github.com/chanxuehong/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Start 方法
func Start() {
	http.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf("%s:%d", "", 9888)
	server := &http.Server{
		Addr:              addr,
		ReadTimeout:       10 * time.Minute,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      30 * time.Minute,
		IdleTimeout:       30 * time.Second,
	}

	// 程序退出的时候平滑关闭 http 服务
	app.Go(context.Background(), func(server *http.Server, addr string) func(ctx context.Context, closing <-chan struct{}) {
		return func(ctx context.Context, closing <-chan struct{}) {
			<-closing

			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Error("failed-to-shutdown-pprof-server", "addr", addr, "error", err.Error())
				return
			}
			log.Info("success-to-shutdown-pprof-server")
		}
	}(server, addr))
	// 启动 http 服务
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("failed-to-start-pprof-server", "addr", addr, "error", err.Error())
		app.Close()
		return
	}
	log.Info("pprof-server-closed")
}
