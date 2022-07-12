package core

type Session struct {
	session map[string]interface{}
}

func (s *Session) Get(name string) interface{} {
	return s.session[name]
}

func (s *Session) Set(name string, value interface{}) *Session {

	s.session[name] = value
	return s
}

func NewSession() *Session {

	log.Debug("Create new session")
	return &Session{session: make(map[string]interface{})}

}
