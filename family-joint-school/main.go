package main

import (
	"context"
	"family-joint-school/pprof"
	"fmt"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/chanxuehong/log"
	_ "github.com/go-sql-driver/mysql" // init

	"family-joint-school/common/app"
	httpx "family-joint-school/handler/http"
	"family-joint-school/handler/http/api"
	"family-joint-school/logic"
	"family-joint-school/mysql/ent"
)

func main() {
	dbClient, err := initEnt()
	if err != nil {
		panic(err)
	}
	logicInterface := logic.NewLogic(dbClient)
	handler := api.NewHandler(logicInterface)
	engine := httpx.Router(handler)

	srv := &http.Server{
		Addr:    ":8888",
		Handler: engine,
	}
	app.Go(context.Background(), func(ctx context.Context, closing <-chan struct{}) {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("run-server-failed")
			app.Close()
		}
	})
	app.Go(context.Background(), func(ctx context.Context, closing <-chan struct{}) {
		<-closing
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		log.Info("start-to-shutdown-engine")
		if err := srv.Shutdown(ctx); err != nil {
			log.Error("failed-to-shutdown-engine", "error", err.Error())
			return
		}
		log.Info("success-to-shutdown-engine")
	})
	go pprof.Start()

	log.Info("success-to-start-engine")
	app.Wait()
}

func initEnt() (*ent.Client, error) {
	driver, err := sql.Open(dialect.MySQL, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "root@123", "127.0.0.1", 3306, "familyJointSchool"))
	if err != nil {
		return nil, err
	}
	driver.DB().SetMaxOpenConns(20)
	driver.DB().SetMaxIdleConns(20)
	driver.DB().SetConnMaxLifetime(3600)
	client := ent.NewClient(ent.Driver(EntLogWithDriver(driver)))
	return client, nil
}

// EntLogWithDriver  add log
func EntLogWithDriver(d dialect.Driver) dialect.Driver {
	return dialect.DebugWithContext(d, func(ctx context.Context, i ...interface{}) {
		if len(i) == 0 {
			return
		}
		sqlQuery, ok := i[0].(string)
		if ok && strings.Contains(sqlQuery, "Exec") {
			log.InfoContext(ctx, "database-exec-log", "query", sqlQuery)
		}
	})
}
