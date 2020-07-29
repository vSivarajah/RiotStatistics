package config

import "os"

const (
	// riot
	keyRiotKey string = "RIOTAPI_KEY"

	// kafka
	keyKafkaBootstrapServers = "KAFKA_BOOTSTRAP_SERVERS"

	// db
	keyDbHost = "DB_HOST"
	keyDbPort = "DB_PORT"
	keyDbUser = "DB_USER"
	keyDbPass = "DB_PASS"
)

func New() (*Config, string, error) {
	c := new(Config)

	if missingField, err := c.setupRiot(); err != nil {
		return nil, missingField, err
	}

	if missingField, err := c.setupKafka(); err != nil {
		return nil, missingField, err
	}

	// db config if required
	return c, "", nil
}


type Config struct {
	DB    db
	Riot  riot
	Kafka kafka
}

type riot struct {
	ApiKey string
}

type kafka struct {
	BootstrapServers string
}

type db struct {
	Host string
	Port string
	User string
	Pass string
}

func (c *Config) setupRiot() (string, error) {
	r := riot{}

	if r.ApiKey = os.Getenv(keyRiotKey); r.ApiKey == "" {
		return keyRiotKey, errEnvNotSet
	}

	c.Riot = r
	return "", nil
}

func (c *Config) setupKafka() (string, error) {
	k := kafka{}

	if k.BootstrapServers = os.Getenv(keyKafkaBootstrapServers); k.BootstrapServers == "" {
		return keyKafkaBootstrapServers, errEnvNotSet
	}

	c.Kafka = k
	return "", nil
}

func (c *Config) setupDB() (string, error) {
	d := db{}

	if d.Host = os.Getenv(keyDbHost); d.Host == "" {
		return keyDbHost, errEnvNotSet
	}

	if d.Port = os.Getenv(keyDbPort); d.Port == "" {
		return keyDbPort, errEnvNotSet
	}

	if d.User = os.Getenv(keyDbUser); d.User == "" {
		return keyDbUser, errEnvNotSet
	}

	if d.Port = os.Getenv(keyDbPass); d.Pass == "" {
		return keyDbPass, errEnvNotSet
	}

	c.DB = d
	return "", nil
}
