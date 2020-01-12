package boot

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/config"
	file_manager "github.com/udayangaac/mobile-api/internal/lib/file-manager"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
	"os"
	"os/signal"
	"syscall"
)

func Init(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	config.Configurations{
		new(config.ServerConfig),
		new(config.DatabaseConfig),
	}.Init(file_manager.NewYamlManager())

	select {
	case <-sigs:
		log.Info(log_traceable.GetMessage(ctx, "Shutting down Application"))
		log.Info(log_traceable.GetMessage(ctx, "Application stopped"))
		os.Exit(0)
	}
}