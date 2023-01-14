package flag

import "flag"

const DefaultModePath = "dev"

type FlagCategory struct {
	Name    string
	Default string
	Usage   string
}

var (
	ServerConfigFlag = &FlagCategory{
		Name:    "server",
		Default: "./was/config/server/config.toml",
		Usage:   "toml file to use for server configuration",
	}

	LogConfigFlag = &FlagCategory{
		Name:    "log",
		Default: "./was/config/log/config.toml",
		Usage:   "toml file to use for log configuration",
	}

	DatabaseFlag = &FlagCategory{
		Name:    "db",
		Default: "./was/config/" + DefaultModePath + "/db/config.toml",
		Usage:   "toml file to use for database configuration",
	}

	JWTFlag = &FlagCategory{
		Name:    "jwt",
		Default: "./was/config/" + DefaultModePath + "/jwt/config.toml",
		Usage:   "toml file to use for jwt configuration",
	}

	RedisFlag = &FlagCategory{
		Name:    "redis",
		Default: "./was/config/" + DefaultModePath + "/redis/config.toml",
		Usage:   "toml file to use for redis configuration",
	}
)

func (f *FlagCategory) Load() *string {
	return flag.String(f.Name, f.Default, f.Usage)
}
