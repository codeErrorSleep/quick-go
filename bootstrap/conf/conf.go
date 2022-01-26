package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Env *viper.Viper

	ErrorLogger *zap.Logger
	DebugLogger *zap.Logger
	InfoLogger  *zap.Logger
	CallLogger  *zap.Logger
)
