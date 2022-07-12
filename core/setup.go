package core

func SetUp(arg ...scenario) {
	log.Debug("Start SetUp function")
	for i, v := range arg {
		log.Debug(i, v.name)
		for _, f := range v.funcs {
			f()
		}
	}
	defer log.Debug("Stop SetUp function")
}
