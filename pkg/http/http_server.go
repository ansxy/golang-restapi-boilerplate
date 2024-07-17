package custom_http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ansxy/golang-boilerplate-gin/config"
)

func NewHttpServer(cnf *config.Config, handler http.Handler) (err error) {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cnf.App.Port),
		Handler: handler,
	}

	serverCtx, ServerStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		log.Println("Gracefull Shutdownm")

		shutdowCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdowCtx.Done()

			if shutdowCtx.Err() == context.DeadlineExceeded {
				log.Printf("Graceful-shutdown-timeout \n%v\n ", err.Error())
			}
		}()

		defer cancel()

		err = srv.Shutdown(shutdowCtx)

		if err != nil {
			log.Printf("Gracefull-Shutdown-Error \n%v\n", err.Error())
		}

		ServerStopCtx()
	}()

	log.Printf("Http-server-online: http://%v:%v", cnf.App.Host, cnf.App.Port)

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("Failed to start Http Server \n%v\n", err.Error())
		return err
	}

	<-serverCtx.Done()

	return
}
