package bundler

import (
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/logger"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/logger/middlewares"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/trace"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/log"
	"bitbucket.org/HeilaSystems/log/bzerolog"
	"bitbucket.org/HeilaSystems/sharedlib/consts"
	"go.uber.org/fx"
)

func LoggerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(func(config configuration.Config) log.Builder {
			defaultLogSetting := bzerolog.LogSettings{
				LogToFile:         false,
				FileJsonFormat:    false,
				LogToConsole:      true,
				ConsoleJsonFormat: false,
				CompressLogsFile:  false,
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
