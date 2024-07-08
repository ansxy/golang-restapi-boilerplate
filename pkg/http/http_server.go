package custome_http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewHttpServer(server *http.Server) (err error) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown http server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("shutdown http server success")
	case <-time.After(time.Second * 5):
		log.Println("shutdown http server timeout")
	}

	log.Println("exit")

	return
}
