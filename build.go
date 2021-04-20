package usvc

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/go-logr/logr"
)

// BannerStdout prints a short banner showing the programme name and version
func BannerStdout(name, version, buildTime string) {
	fmt.Printf("%s %s\n", name, version)
}

// BannerStdoutLong prints a long banner detailing things about the programme's build and runtime environment
func BannerStdoutLong(name, version, buildTime string) {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		`%s %s
Built at %s from %s by %s
%s running on %s under %s %s
Binary %s
Running as %s in %s
`,
		name, version,
		buildTime, "source dir", "build UID",
		runtime.Version(), host, runtime.GOOS, runtime.GOARCH,
		exe,
		user.Name, cwd,
	)
}

// BannerLog logs a few pieces of information about the programme; name, version, build time
func BannerLog(log logr.Logger, name, version, buildTime string) {
	log.Info(name, "version", version, "build time", buildTime)
}

// GetExeDir returns the directory in which the running binary is located
func GetExeDir(log logr.Logger) string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exe)
	cwd, _ := os.Getwd() // just for info
	log.V(1).Info("Executable directory", "exeDir", exeDir, "CWD", cwd)
	return exeDir
}
