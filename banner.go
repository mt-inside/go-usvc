package usvc

import "github.com/go-logr/logr"

func BannerLog(log logr.Logger, name string) {
	log.Info(name, "version", "?")
}
