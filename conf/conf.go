package conf

import (
	"code.google.com/p/gcfg"
	"runtime"
	"short_url/common/snowflake"
)

var (
	// global config object
	Conf *Config
)

type Config struct {
	Base struct {
		User      string   //`gcfg:"user"`
		MaxProc   int      `gcfg:"maxproc"`
		PidFile   string   `gcfg:"pid"`
		PprofAddr []string `gcfg:"pprof"`
		DirPath   string   `gcfg:"dir"`
		LogPath   string   `gcfg:"log"`
		Domain string
	}

	Http struct {
		Addr []string
		Ping string
		Port string
	}



	Redis map[string]*struct {
		Addr      string
		MaxActive int
		MaxIdle   int
		Timeout   int
		Password  string
	}

	Snowflake snowflake.Config


}

func init() {
	Conf = &Config{}
	Conf.Base.User = "nobody"
	Conf.Base.MaxProc = runtime.NumCPU()
	Conf.Base.PidFile = "/tmp/flume.pid"
	Conf.Base.DirPath = "/dev/null"

}

// Init init the configuration file.
func Init(confPath string) error {
	if err := gcfg.ReadFileInto(Conf, confPath); err != nil {
		return err
	}
	return nil
}

// Reload reload the configuration file.
func Reload(confPath string) error {
	tmp := &Config{}
	if err := gcfg.ReadFileInto(tmp, confPath); err != nil {
		return err
	}
	Conf = tmp
	return nil
}
