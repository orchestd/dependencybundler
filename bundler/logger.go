package bundler

import (
	"github.com/orchestd/dependencybundler/constructors/logger"
	"github.com/orchestd/dependencybundler/constructors/logger/middlewares"
	"github.com/orchestd/dependencybundler/constructors/trace"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/log"
	"github.com/orchestd/log/bzerolog"
	"github.com/orchestd/sharedlib/consts"
	"go.uber.org/fx"
)

func LoggerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(func(config configuration.Config) log.Builder {
			defaultLogSetting := bzerolog.LogSettings{
				LogToFile:           false,
				FileJsonFormat:      false,
				LogToConsole:        true,
				ConsoleJsonFormat:   false,
				CompressLogsFile:    false,
				DisableConsoleColor: false,
			}
			if fileLog, err := config.Get(consts.LogToFile).Bool(); err == nil {
				defaultLogSetting.LogToFile = fileLog
			}
			if fileJsonFormat, err := config.Get(consts.FileJsonFormat).Bool(); err == nil {
				defaultLogSetting.FileJsonFormat = fileJsonFormat
			}
			if consoleLog, err := config.Get(consts.LogToConsole).Bool(); err == nil {
				defaultLogSetting.LogToConsole = consoleLog
			}

			if consoleJsonFormat, err := config.Get(consts.ConsoleJsonFormat).Bool(); err == nil {
				defaultLogSetting.ConsoleJsonFormat = consoleJsonFormat
			}

			if compressLogs, err := config.Get(consts.CompressLogs).Bool(); err == nil {
				defaultLogSetting.CompressLogsFile = compressLogs
			}
			if disableConsoleColor, err := config.Get(consts.DisableConsoleColor).Bool(); err == nil {
				defaultLogSetting.DisableConsoleColor = disableConsoleColor
			}

			return bzerolog.DefaultZeroLogBuilder(defaultLogSetting)
		}),
		fx.Provide(fx.Annotated{Group: consts.FxGroupLoggerContextExtractors, Target: middlewares.DefaultLoggerIncomingContextExtractor}),

		trace.TraceInfoContextExtractorFxOption(),

		fx.Provide(logger.DefaultLogger),
	)
}
