package config

import (
	configEntity "crow-blog-backend/src/config/entity"
	authType "crow-blog-backend/src/consts/auth_type"
	"crow-blog-backend/src/entity"
	encryptUtil "crow-blog-backend/src/utils/encrypt"
	panicUtil "crow-blog-backend/src/utils/painc"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/redis/go-redis/v9"
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
	"regexp"
	"runtime"
	"strings"
	"time"
)

var envConfig = &configEntity.EnvConfig{}
var db *gorm.DB
var globalLogger *zap.SugaredLogger
var redisDb *redis.Client

var app *iris.Application

func SetApp(appInstance *iris.Application) {
	app = appInstance
}

func GetApp() *iris.Application {
	return app
}

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
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func initGlobalLogger() {
	encoder := getEncoder()
	// 方便测试时观察日志打印，打包后不需要控制台输出
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
	globalLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0)).Sugar()
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

func initSysUser() {

	accountRegexp, err := regexp.Compile("^[a-zA-Z][a-zA-Z0-9_]{4,15}$")
	if err != nil {
		panicUtil.CustomPanic("Failed to compile account regexp", err)
	}
	passwordRegexp, err := regexp.Compile("^[a-zA-Z]\\w{5,17}$")
	if err != nil {
		panicUtil.CustomPanic("Failed to compile password regexp", err)
	}
	emailRegexp, err := regexp.Compile("^[a-zA-Z0-9]+([-_.][A-Za-zd]+)*@([a-zA-Z0-9]+[-.])+[A-Za-zd]{2,5}$")
	if err != nil {
		panicUtil.CustomPanic("Failed to compile email regexp", err)
	}

	var userCount int64
	db.Model(&entity.User{}).Count(&userCount)
	if userCount != 0 {
		return
	}
	//用户数为零 判断信息是否均不为空且符合需求
	sysUser := envConfig.SysUser
	if !accountRegexp.MatchString(sysUser.Account) ||
		!passwordRegexp.MatchString(sysUser.Password) ||
		!emailRegexp.MatchString(sysUser.Email) ||
		len(strings.Trim(sysUser.Nickname, " ")) == 0 {
		return
	}

	if trErr := db.Transaction(func(trDb *gorm.DB) error {
		user := &entity.User{
			Nickname: sysUser.Nickname,
		}
		db.Create(user)

		// 加密密码
		encryptPassword, encryptErr := encryptUtil.EncryptPassword(sysUser.Password)
		if encryptErr != nil {
			return encryptErr
		}

		accountUserAuth := entity.UserAuth{
			Type:       authType.Account,
			UserId:     user.ID,
			Identifier: sysUser.Account,
			Credential: encryptPassword,
			CreatedAt:  time.Now(),
		}

		emailAuth := entity.UserAuth{
			Type:       authType.Email,
			UserId:     user.ID,
			Identifier: sysUser.Email,
			Credential: encryptPassword,
			CreatedAt:  time.Now(),
		}
		auths := make([]entity.UserAuth, 0)
		auths = append(auths, accountUserAuth, emailAuth)
		return db.CreateInBatches(auths, len(auths)).Error
	}); trErr != nil {
		globalLogger.Error("创建系统用户失败")
	}
}

func initCache() {
	cache := envConfig.Cache
	if !cache.Use {
		return
	}
	cpuNum := runtime.NumCPU()

	fmt.Println(cache.Redis)

	rdb := redis.NewClient(&redis.Options{
		Addr:         cache.Redis.Address,
		Password:     cache.Redis.Password, // 没有密码，默认值
		DB:           cache.Redis.Db,       // 默认DB 0
		MaxIdleConns: cpuNum * 2,
		MinIdleConns: 0,
	})
	redisDb = rdb

}

func GetRedisClient() *redis.Client {
	return redisDb
}

func InitConfig() {
	initEnvConfig()
	initDataSource()
	initGlobalLogger()
	initSysUser()
	initCache()

	fmt.Println(envConfig)
}
