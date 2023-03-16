package config

import (
	configEntity "crow-blog-backend/src/config/entity"
	"crow-blog-backend/src/entity"
	panicUtil "crow-blog-backend/src/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

var envConfig = &configEntity.EnvConfig{}
var db *gorm.DB
var globalLogger *zap.SugaredLogger

func initEnvConfig() {
	file, err := os.ReadFile("./application.yml")
	if err != nil {
		panicUtil.CustomPanic("Failed to open file", err)
	}
	err = yaml.Unmarshal(file, envConfig)
	if err != nil {
		panicUtil.CustomPanic("Failed unmarshal envConfig", err)
	}
}

func initDataSource() {
	dsn := envConfig.DataSource.MysqlDsn

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 是否使用彩色打印
		})

	dbInstance, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if dbErr != nil {
		panicUtil.CustomPanic("Failed to init database instance", dbErr)
	}

	if err := dbInstance.AutoMigrate(
		&entity.User{},
		&entity.Link{},
		&entity.UserAuth{},
		&entity.About{},
	); err != nil {
		panicUtil.CustomPanic("Failed to AutoMigrate", err)
	}

	db = dbInstance
}

func getLogWriter(filePath string, outConsole bool) io.Writer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   false,
	}
	// 控制info级别日志输出到控制台和文件
	if outConsole {
		return io.MultiWriter(lumberJackLogger, os.Stdout)
	}
	return lumberJackLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func initGlobalLogger() {
	encoder := getEncoder()
	infoWriter := getLogWriter("./logs/blog.info.log", true)
	warnWriter := getLogWriter("./logs/blog.warn.log", false)
	errorWriter := getLogWriter("./logs/blog.error.log", false)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	globalLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func GetDatabaseInstance() *gorm.DB {
	return db
}

func GetEnvConfig() *configEntity.EnvConfig {
	return envConfig
}

func GetGlobalLogger() *zap.SugaredLogger {
	return globalLogger
}

func InitConfig() {
	initEnvConfig()
	initDataSource()
	initGlobalLogger()
}
