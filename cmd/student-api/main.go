package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http" 
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/annuvrat/go-rest/internal/config"
	student "github.com/annuvrat/go-rest/internal/http/handlers"
)


func main(){
//load config
cfg:= config.MustLoad()


//router setup
router :=http.NewServeMux()

router.HandleFunc("POST /api/students",student.New())

server:=http.Server{
	Addr: cfg.Addr,
	Handler: router,
}

slog.Info("server started", slog.String("addr", cfg.Addr))
fmt.Printf("server started %s", cfg.Addr)


done:= make(chan os.Signal, 1)

signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
go func (){
	err:=server.ListenAndServe()
	if err!= nil{
		log.Fatal("failed to start server")
	}
  
}()

<-done

slog.Info("shutting down the server")

ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)

defer cancel()
err :=server.Shutdown(ctx)

if err!= nil{
	slog.Error("failed to shutdown server",slog.String("error",err.Error()))
}

slog.Info("server shutdown successfully")
}
