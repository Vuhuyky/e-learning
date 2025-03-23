package connection

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{
		"log/app.log",
		"stdout",
	}

	config.ErrorOutputPaths = []string{
		"log/error.log",
		"stderr",
	}

	config.EncoderConfig = zapcore.EncoderConfig{
		LevelKey:     "LEVEL",
		TimeKey:      "TIME",
		CallerKey:    "CALLER",
		FunctionKey:  "FUNC",
		MessageKey:   "MESSAGE",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		log.Fatalln("error init logger: ", err)
		return
	}
}
