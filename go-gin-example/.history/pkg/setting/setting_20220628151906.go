package setting

import (
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg         *ini.File
	RunMode     string
	HTTPPort    int
	ReadTimeout time.Duration
)
