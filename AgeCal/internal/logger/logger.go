package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init() error {
	I, err := zap.NewProduction()
	if err != nil {
		return err
	}
	Log = I
	Log.Info("logger Initialized")
	return nil
}
