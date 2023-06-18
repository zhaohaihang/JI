package backproc

type BackProc interface {
	start() error
	stop()
}

type BackProcServer struct {
	BackProcPool []BackProc
}

func NewBackProcServer(esp *EsSyncProc , rmp *RemindMailProc) *BackProcServer {
	var backProcServer BackProcServer
	backProcServer.addBackProc(esp)
	backProcServer.addBackProc(rmp)
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
