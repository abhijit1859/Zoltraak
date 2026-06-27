package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger


func InitLogger(env string) error{
	var config zap.Config

	if env=="production"{
		config=zap.NewProductionConfig()
		config.EncoderConfig.TimeKey="timestamp"
		config.EncoderConfig.EncodeTime=zapcore.ISO8601TimeEncoder

	}else{
		config=zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel=zapcore.CapitalColorLevelEncoder

	}
	config.OutputPaths=[]string{"stdout"}
	config.ErrorOutputPaths=[]string{"stderr"}

	logger,err:=config.Build()
	if err!=nil{
		return  err
	}

	Log=logger

	zap.ReplaceGlobals(logger)
	return nil
}


type Field=zap.Field


func Info(msg string,fields ...Field){
	Log.Info(msg,fields...)
}



func Error(msg string,fields ...Field){
	Log.Error(msg,fields...)
}


func Fatal(msg string,fields ...Field){
	Log.Fatal(msg,fields...)
}


func Debug(msg string,fields ...Field){
	Log.Debug(msg,fields...)
}