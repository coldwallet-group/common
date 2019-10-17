package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var(
	log *zap.Logger
)

const (
	EnvDevelopment = "dev"
	EnvProd = "prod"
	LvlDebug = "debug"
	LvlInfo = "info"
	LvlWarn = "warn"
	LvlError = "error"
	LvlDPanic = "dpanic"
	LvlPanic = "panic"
	LvlFatal = "fatal"
	TxtJson = "json"
	TxtConsole = "console"
)

type Logcfg struct {
	Level            string
	Env              string
	TxtType          string
	TimeKey          string
	LevelKey         string
	NameKey          string
	CallerKey        string
	MessageKey       string
	StacktraceKey    string
	OutputPaths      []string
	ErrorOutputPaths []string
}


//func sample() {
//	//_log,_ := zap.NewProduction(zap.AddCaller())
//	//defer _log.Sync()
//	InitLog()
//	Debug("debug,最灵繁的人也看不见自己的背脊")
//	Info("info,最困难的事情就是认识自己。")
//	Warn("warn,有勇气承担命运这才是英雄好汉")
//	Error("error,与肝胆人共事，无字句处读书。")
//	DPanic("dpanic,阅读使人充实，会谈使人敏捷，写作使人精确。")
//	//Panic("panic,最大的骄傲于最大的自卑都表示心灵的最软弱无力。")
//	//Fatal("fatal,自知之明是最难得的知识。")
//	Debugf("debugf,勇气通往天堂，怯懦通往地狱。")
//	Infof("infof,有时候读书是%s一种巧%s妙地避开思考%s的方法。","test","demo","done")
//	Warnf("warnf,阅读%s一切好书%s如同和过去%s最杰出的人谈话。","test","demo","done")
//	Errorf("errorf,越是%s没有本领%s的就越加%s自命不凡。","test","demo","done")
//	DPanicf("dpanicf,越是%s无能的人，%s越喜欢挑剔%s别人的错儿。","test","demo","done")
//	//Panicf("panicf,知人者智%s，自知者明%s。胜人者有力%s，自胜者强。","test","demo","done")
//	//Fatalf("fatalf,意志坚强%s的人能把%s世界放在手中像泥块%s一样任意揉捏。","test","demo","done")
//}

func InitLog(cfg *Logcfg)  {
	var (
		lvl = zap.DebugLevel
		isDev = true
		txtType = TxtJson
		outputPaths = []string{"stderr"}
		erroroutputPaths = []string{"stderr"}
	)
	switch cfg.Level {
	case LvlDebug:
		lvl= zap.DebugLevel
	case LvlInfo:
		lvl = zap.InfoLevel
	case LvlWarn:
		lvl = zap.WarnLevel
	case LvlError:
		lvl = zap.ErrorLevel
	case LvlDPanic:
		lvl = zap.DPanicLevel
	case LvlPanic:
		lvl = zap.PanicLevel
	case LvlFatal:
		lvl = zap.FatalLevel
	default:
		lvl = zap.DebugLevel
	}

	if cfg.TxtType == TxtJson {
		txtType = TxtJson
	} else {
		txtType = TxtConsole
	}

	if cfg.Env == EnvProd {
		isDev=false
	} else {
		isDev=true
	}

	if len(cfg.OutputPaths) > 0 {
		outputPaths = cfg.OutputPaths
	} else {
		outputPaths = []string{"stderr"}
	}

	if len(cfg.ErrorOutputPaths) > 0 {
		erroroutputPaths = cfg.ErrorOutputPaths
	} else {
		erroroutputPaths = []string{"stderr"}
	}

	log,_ = zap.Config {
		Level:       zap.NewAtomicLevelAt(lvl),
		Development: isDev,
		Sampling: &zap.SamplingConfig {
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         txtType,
		EncoderConfig:    zapcore.EncoderConfig{
			TimeKey:        cfg.TimeKey,
			LevelKey:       cfg.Level,
			NameKey:        cfg.NameKey,
			CallerKey:      cfg.CallerKey,
			MessageKey:     cfg.MessageKey,
			StacktraceKey:  cfg.StacktraceKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,//zapcore.SecondsDurationEncoder,//zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      outputPaths,
		ErrorOutputPaths: erroroutputPaths,
	}.Build(zap.AddCaller())
}

//func Debug(msg string) {
//	log.Debug(msg)
//}
//
//func Info(msg string) {
//	log.Info(msg)
//}
//
//func Warn(msg string) {
//	log.Warn(msg)
//}
//
//func Error(msg string) {
//	log.Error(msg)
//}
//
//func DPanic(msg string) {
//	log.DPanic(msg)
//}
//
//func Panic(msg string) {
//	log.Panic(msg)
//}
//
//func Fatal(msg string) {
//	log.Fatal(msg)
//}

func Debug(args ...interface{}) {
	log.Debug(fmt.Sprint(args...))
}

func Info(args ...interface{}) {
	log.Info(fmt.Sprint(args...))
}

func Warn(args ...interface{}) {
	log.Warn(fmt.Sprint(args...))
}

func Error(args ...interface{}) {
	log.Error(fmt.Sprint(args...))
}

func DPanic(args ...interface{}) {
	log.DPanic(fmt.Sprint(args...))
}

func Panic(args ...interface{}) {
	log.Panic(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	log.Fatal(fmt.Sprint(args...))
}

func Debugf(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Debug(msg)
}

func Infof(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Info(msg)
}

func Warnf(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Warn(msg)
}

func Errorf(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Error(msg)
}

func DPanicf(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.DPanic(msg)
}

func Panicf(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Panic(msg)
}

func Fatalf(format string, args ...interface{}) {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	log.Fatal(msg)
}

func Sync() {
	if log == nil {
		return
	}
	log.Sync()
}


func Debugw(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.Debug(msg,fs...)
		return
	}
	log.Debug(msg)
}

func Infow(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.Info(msg,fs...)
		return
	}
	log.Info(msg)
}

func Warnw(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.Warn(msg,fs...)
		return
	}
	log.Warn(msg)
}

func Errorw(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.Error(msg,fs...)
		return
	}
	log.Error(msg)
}

func DPanicw(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.DPanic(msg,fs...)
		return
	}
	log.DPanic(msg)
}

func Panicw(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.Panic(msg,fs...)
		return
	}
	log.Panic(msg)
}

func Fatalw(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0, len(fields))
		for k,v:=range fields {
			fs=append(fs,zap.Any(k,v))
		}
		log.Fatal(msg,fs...)
		return
	}
	log.Fatal(msg)
}

