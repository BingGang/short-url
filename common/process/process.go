package process

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Init create pid file, set working dir, setgid and setuid.
func Init(pidFile string) error {
	if err := ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d\n", os.Getpid())), 0644); err != nil {
		return err
	}
	return nil
}
