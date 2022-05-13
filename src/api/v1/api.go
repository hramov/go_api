package api

func Init(prefix string) {
	initRestApi(prefix)
	initGrpcApi()
}

func Start() {
	go createRestServer()
	go createGRPCServer()
}
