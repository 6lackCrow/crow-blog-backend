package configEntity

type EnvConfig struct {
	DataSource struct {
		MysqlDsn string `yaml:"mysql-dsn" required:"true"`
	} `yaml:"data-source"`

	Server struct {
		Port string `required:"true" yaml:"port"`
	} `yaml:"server"`

	SysUser struct {
		Account  string `yaml:"account"`
		Password string `yaml:"password"`
		Email    string `yaml:"email"`
		Nickname string `yaml:"nickname"`
	} `yaml:"sys-user"`

	Cache struct {
		Use   bool `yaml:"use"`
		Redis struct {
			Address  string `yaml:"address"`
			Password string `yaml:"password"`
			Db       int    `yaml:"db"`
		} `yaml:"redis"`
	} `yaml:"cache"`
}
