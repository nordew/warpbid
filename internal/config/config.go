package config

type Config struct {
	Cockroach CockroachConfig `env-prefix:"COCKROACH_"`
	Dragonfly DragonflyConfig `env-prefix:"DRAGONFLY_"`
	Scylla    ScyllaConfig    `env-prefix:"SCYLLA_"`
	NATS      NATSConfig      `env-prefix:"NATS_"`
}

type CockroachConfig struct {
	Host     string `env:"HOST,required"`
	Port     int    `env:"PORT" env-default:"26257"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE,required"`
	SSLMode  string `env:"SSLMODE" env-default:"disable"`
}

type DragonflyConfig struct {
	Host     string `env:"HOST,required"`
	Port     int    `env:"PORT" env-default:"6379"`
	Password string `env:"PASSWORD"`
}

type ScyllaConfig struct {
	Hosts    []string `env:"HOSTS,required" env-delim:","`
	Port     int      `env:"PORT" env-default:"9042"`
	Keyspace string   `env:"KEYSPACE,required"`
	User     string   `env:"USER"`
	Password string   `env:"PASSWORD"`
}

type NATSConfig struct {
	URL string `env:"URL" env-default:"nats://localhost:4222"`
}
