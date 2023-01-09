package usvc

import (
	"fmt"
	golog "log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/mattn/go-isatty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"
)

var (
	// Global is a globally-scoped logger than can be used if you have no other choice
	Global logr.Logger

	zapCfg zap.Config
)

func init() {
	spew.Config.DisableMethods = true
	spew.Config.DisablePointerMethods = true

	Global = GetLogger(false).WithName("GLOBAL")
}

// GetLogger returns a zap-based zapr Logger, typed as a logr.Logger
func GetLogger(devMode bool, options ...int) logr.Logger {
	var devLevel, prodLevel zapcore.Level
	var err error

	if len(options) == 0 {
		devLevel = -10
		prodLevel = zap.InfoLevel
	} else {
		devLevel = zapcore.Level(-1 * options[0])
		prodLevel = zapcore.Level(-1 * options[0])
		if devLevel > 0 {
			panic(fmt.Errorf("logging level must be +ve"))
		}
	}

	if isatty.IsTerminal(os.Stdout.Fd()) || devMode {
		zapCfg = zap.NewDevelopmentConfig()
		zapCfg.EncoderConfig.EncodeCaller = nil
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapCfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("15:04:05"))
		}
		zapCfg.Level.SetLevel(devLevel)
	} else {
		zapCfg = zap.NewProductionConfig()
		zapCfg.Level.SetLevel(prodLevel)
	}
	zapLog, err := zapCfg.Build()
	if err != nil {
		panic(err.Error())
	}

	zr := zapr.NewLogger(zapLog)

	if devMode {
		zr.Info("Logging in dev mode; remove --dev flag for structured json output")
	}

	/* == Intercept pkg log == */

	golog.SetFlags(0) // don't add date and timestamps to the message, as the zapr writer will do that
	golog.SetOutput(zaprWriter{zr.WithValues("source", "go log")})

	/* Intercept klog == */

	/* k8s stuff uses klog which is both an interface (with text and
	* structured methods) and an impl (which prints only as text, not
	* json). Can rewire its impl to something that outputs json (this is
	* typed as a logr, so eg zapr), which makes all log-lines
	* machine-parsable. However most of the k8s stuff doesn't use the
	* structured part of the interface, so you just get msg="blob". */
	klog.SetLogger(zr)

	return zr
}

// SetLevel sets the level of the entire tree of loggers returned from GetLogger
func SetLevel(l int) {
	zapCfg.Level.SetLevel(zapcore.Level(-1 * l))
}

type zaprWriter struct{ log logr.Logger }

func (w zaprWriter) Write(data []byte) (n int, err error) {
	w.log.Info(string(data))
	return len(data), nil
}
