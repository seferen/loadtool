package core

import "github.com/seferen/logger"

var log = logger.GetLoger("core")

type scenario struct {
	name    string
	session *Session
	funcs   []func()
}

func Scenario(name string) *scenario {
	log.Debug("Init scenario", name)

	return &scenario{name: name, session: NewSession()}
}

func (s *scenario) Exec(act Actionbuilder) *scenario {
	log.Debug("add fuction to", s.name)
	s.funcs = append(s.funcs, func() {
		act.Act()
		log.Debug(s.session)
	})

	return s
}

func (s *scenario) ExecFn(fn func(session *Session)) *scenario {
	log.Debug("add function from ExecFn", s.name)
	s.funcs = append(s.funcs, func() { fn(s.session) })

	return s
}
