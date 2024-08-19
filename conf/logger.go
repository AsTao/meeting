package conf

import (
	"os"
	"path/filepath"
	"time"

	"github.com/AsTao/meeting/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	mode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		mode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), mode)
	global.Logger = zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	separator := string(filepath.Separator)
	rootiDir, _ := os.Getwd()
	logFilePath := rootiDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".txt"

	lumberjackSyncer := lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.Maxsize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   false, //  #是否开启压缩；
	}
	return zapcore.AddSync(&lumberjackSyncer)

}
