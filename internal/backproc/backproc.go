package backproc

type BackProc interface {
	start() error
	stop()
}

type BackProcServer struct {
	BackProcPool []BackProc
}

func NewBackProcServer(
	esp *EsSyncProc , 
	rmp *RemindMailProc,
	dup *DBUpdateProc,) *BackProcServer {
	var backProcServer BackProcServer
	backProcServer.addBackProc(esp)
	backProcServer.addBackProc(rmp)
	backProcServer.addBackProc(dup)
	return &backProcServer
}

func (bps *BackProcServer) Start() error {
	for _, backProc := range bps.BackProcPool {
		err := backProc.start()
		if err != nil {
			return err 
		}
	}
	return nil
}

func (bps *BackProcServer) Stop() {
	for _, backProc := range bps.BackProcPool {
		backProc.stop()
	}
}

func (bps *BackProcServer) addBackProc(bp BackProc) {
	bps.BackProcPool = append(bps.BackProcPool, bp)
}
