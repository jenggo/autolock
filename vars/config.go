package vars

var Config Configs

type Configs struct {
	Path struct {
		Image  string `yaml:"image" env-default:"/tmp/lastsnap.jpg`
		Models string `yaml:"model" env-default:"/etc/models"`
	} `yaml:"path"`

	Threshold   int    `yaml:"threshold" env-default:"7"`
	LockCommand string `yaml:"lock_command" env-default:"betterlockscreen --lock --span --off 5"`
	Verbose     bool   `yaml:"verbose" env-default:"false"`
}
