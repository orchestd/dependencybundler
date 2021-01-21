package bundler

import (
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/logger"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/logger/middlewares"
	"bitbucket.org/HeilaSystems/dependencybundler/constructors/trace"
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/log"
	"bitbucket.org/HeilaSystems/log/bzerolog"
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
			}
			if fileLog, err := config.Get(depBundler.LogToFile).Bool(); err == nil {
				defaultLogSetting.LogToFile = fileLog
			}
			if fileJsonFormat, err := config.Get(depBundler.FileJsonFormat).Bool(); err == nil {
				defaultLogSetting.FileJsonFormat = fileJsonFormat
			}
			if consoleLog, err := config.Get(depBundler.LogToConsole).Bool(); err == nil {
				defaultLogSetting.LogToConsole = consoleLog
			}

			if consoleJsonFormat, err := config.Get(depBundler.ConsoleJsonFormat).Bool(); err == nil {
				defaultLogSetting.ConsoleJsonFormat = consoleJsonFormat
			}
			
			if compressLogs, err := config.Get(depBundler.CompressLogs).Bool(); err == nil {
				defaultLogSetting.CompressLogsFile = compressLogs
			}

			return bzerolog.DefaultZeroLogBuilder(defaultLogSetting)
		}),
		fx.Provide(fx.Annotated{Group: depBundler.FxGroupLoggerContextExtractors, Target: middlewares.DefaultLoggerIncomingContextExtractor}),

		trace.TraceInfoContextExtractorFxOption(),

		fx.Provide(logger.DefaultLogger),

	)
}
