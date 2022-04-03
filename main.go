package main

import (
	"fmt"
	"github.com/candyMonster1/tools/logger"
	"go.uber.org/zap"
)

func main() {
	//logger,err :=  zap.Config{
	//	Encoding: "json",
	//	Level: zap.NewAtomicLevelAt(zapcore.WarnLevel),
	//	OutputPaths: []string{"stdout"},
	//	EncoderConfig: zapcore.EncoderConfig{
	//		MessageKey: "message",
	//		LevelKey:    "level",
	//		EncodeLevel: zapcore.CapitalLevelEncoder,// INFO
	//
	//		TimeKey:    "time",
	//		EncodeTime: zapcore.ISO8601TimeEncoder,
	//
	//		CallerKey:    "caller",
	//		EncodeCaller: zapcore.ShortCallerEncoder,
	//	},
	//}.Build()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer logger.Sync()
	//
	//logger.Warn("test")

	logger1, err := logger.LogLevel("info")
	if err != nil {
		fmt.Println(err)
	}

	log := zap.ReplaceGlobals(logger1)
	defer log()

	zap.L().Named("test").Info("test haha")
}