package main

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func main() {
	// 创建 logger
	logger := otelzap.New(
		zap.NewExample(),                    // zap实例，按需配置
		otelzap.WithMinLevel(zap.InfoLevel), // 指定日志级别
		otelzap.WithTraceIDField(true),      // 在日志中记录 traceID
	)
	defer logger.Sync()
	// logger.Debug()
	logger.Ctx(context.TODO()).Debug("...")

	// 替换全局的logger
	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	// zap.L().
	otelzap.L().Info("replaced zap's global loggers")        // 记录日志
	otelzap.Ctx(context.TODO()).Info("... and with context") // 从ctx中获取traceID并记录
}
