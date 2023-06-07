package backproc

type BackProc interface {
	start() error
	stop()
}

type BackProcServer struct {
	BackProcPool []BackProc
}

func NewBackProcServer(esp *EsSyncProc) *BackProcServer {
	var backProcServer BackProcServer
	backProcServer.addBackProc(esp)
	return &backProcServer
}

func (bps *BackProcServer) Start() {
	for _, backProc := range bps.BackProcPool {
		backProc.start()
	}
}

func (bps *BackProcServer) Stop() {
	for _, backProc := range bps.BackProcPool {
		backProc.stop()
	}
}

func (bps *BackProcServer) addBackProc(bp BackProc) {
	bps.BackProcPool = append(bps.BackProcPool, bp)
}
