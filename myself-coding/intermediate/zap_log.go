package intermediate

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Println("Error in initializing Zap logger")
	}
	defer logger.Sync()

	// Run some logging.
	logger.Info("This is an info message.")
	logger.Info(
		"User logged in",
		zap.String("username", "Tho Le"),
		zap.String("method", "GET"),
	)
}
