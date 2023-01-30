package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"main/pkg/conf"
	"main/pkg/result"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("empty or invalid id\n"))
	}
}
func init() {

}
func Router() {
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(hanlder)) // Вывод городов в соотвествии с запросом

	srv := &http.Server{
		Addr:    ":" + conf.Con.Apport,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
func hanlder(w http.ResponseWriter, r *http.Request) error {
	Fr, _ := json.Marshal(result.MakeStruct())
	//fmt.Println("result: ", Fr)
	w.Write([]byte(Fr))
	render.Status(r, http.StatusOK)
	return nil
}
