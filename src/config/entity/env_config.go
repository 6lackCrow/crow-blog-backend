package configEntity

type EnvConfig struct {
	DataSource struct {
		MysqlDsn string `yaml:"mysql-dsn" required:"true"`
	} `yaml:"data-source"`

	Server struct {
		Port string `required:"true" yaml:"port"`
	} `yaml:"server"`
}
