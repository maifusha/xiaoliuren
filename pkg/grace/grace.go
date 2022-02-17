package grace

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type grace struct {
	srv  *http.Server
	quit chan os.Signal
}

func New(server *http.Server) *grace {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	return &grace{srv: server, quit: quit}
}

func (g *grace) Down() {
	<-g.quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := g.srv.Shutdown(ctx); err != nil {
		log.Println("Server shutdown：", err)
	}
}
