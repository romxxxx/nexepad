package common

import (
	"fmt"
	"os"
	"sync/atomic"
	"syscall"
	"testing"

	"github.com/romxxxx/nexepad/domain/dagconfig"
)

// RunnexepadForTesting runs nexepad for testing purposes
func RunnexepadForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	nexepadRunCommand, err := StartCmd("nexePAD",
		"nexepad",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("nexepad started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := nexepadRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("nexepad closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := nexepadRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("nexepad stopped")
	}
}
