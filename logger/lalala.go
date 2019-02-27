package logger

import (
	"google.golang.org/grpc/grpclog"
	"os"
)

func main() {

	var logger = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)

	logger.V(1)

}
