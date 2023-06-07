package backproc

import "github.com/google/wire"

var BackProcServerProviderSet = wire.NewSet(
	NewBackProcServer,
	NewEsSyncProc,
)
