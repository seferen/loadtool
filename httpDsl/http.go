package httpDsl

import "github.com/seferen/logger"

type http struct {
	name string
}

var log = logger.GetLoger("httpDsl")

func Http(name string) *http {
	log.Debug("Create new htpp request", name)
	return &http{name: name}

}

func (h *http) Act() error {
	log.Debug("Start action")
	return nil

}
