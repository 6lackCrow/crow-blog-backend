package globalLogger

import "crow-blog-backend/src/config"

func Info(args ...interface{}) {
	config.GetGlobalLogger().Info(args...)
}
func Infof(template string, args ...interface{}) {
	config.GetGlobalLogger().Infof(template, args...)
}
func Warn(args ...interface{}) {
	config.GetGlobalLogger().Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	config.GetGlobalLogger().Warnf(template, args...)
}
func Error(args ...interface{}) {
	config.GetGlobalLogger().Error(args...)
}
func Errorf(template string, args ...interface{}) {
	config.GetGlobalLogger().Errorf(template, args...)
}
