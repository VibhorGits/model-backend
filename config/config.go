package config

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

// ServerConfig defines HTTP server configurations
type ServerConfig struct {
	PrivatePort int `koanf:"privateport"`
	PublicPort  int `koanf:"publicport"`
	HTTPS       struct {
		Cert string `koanf:"cert"`
		Key  string `koanf:"key"`
	}
	Edition string `koanf:"edition"`
	Usage   struct {
		Enabled    bool   `koanf:"enabled"`
		TLSEnabled bool   `koanf:"tlsenabled"`
		Host       string `koanf:"host"`
		Port       int    `koanf:"port"`
	}
	Debug  bool `koanf:"debug"`
	ItMode struct {
		Enabled bool `koanf:"enabled"`
	}
	MaxDataSize int `koanf:"maxdatasize"`
	Workflow    struct {
		MaxWorkflowTimeout int32 `koanf:"maxworkflowtimeout"`
		MaxWorkflowRetry   int32 `koanf:"maxworkflowretry"`
		MaxActivityRetry   int32 `koanf:"maxactivityretry"`
	}
}

// DatabaseConfig related to database
type DatabaseConfig struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Name     string `koanf:"name"`
	Version  uint   `koanf:"version"`
	TimeZone string `koanf:"timezone"`
	Pool     struct {
		IdleConnections int           `koanf:"idleconnections"`
		MaxConnections  int           `koanf:"maxconnections"`
		ConnLifeTime    time.Duration `koanf:"connlifetime"`
	}
}

// TritonServerConfig related to Triton server
type TritonServerConfig struct {
	GrpcURI    string `koanf:"grpcuri"`
	ModelStore string `koanf:"modelstore"`
}

// MgmtBackendConfig related to mgmt-backend
type MgmtBackendConfig struct {
	Host        string `koanf:"host"`
	PrivatePort int    `koanf:"privateport"`
	HTTPS       struct {
		Cert string `koanf:"cert"`
		Key  string `koanf:"key"`
	}
}

// CacheConfig related to Redis
type CacheConfig struct {
	Redis struct {
		RedisOptions redis.Options `koanf:"redisoptions"`
	}
	Model bool `koanf:"model"`
}

// ControllerConfig related to controller
type ControllerConfig struct {
	Host        string `koanf:"host"`
	PrivatePort int    `koanf:"privateport"`
	HTTPS       struct {
		Cert string `koanf:"cert"`
		Key  string `koanf:"key"`
	}
}

type InitModelConfig struct {
	Path    string `koanf:"path"`
	Enabled bool   `koanf:"enabled"`
}

// MaxBatchSizeConfig defines the maximum size of the batch of a AI task
type MaxBatchSizeConfig struct {
	Unspecified          int `koanf:"unspecified"`
	Classification       int `koanf:"classification"`
	Detection            int `koanf:"detection"`
	Keypoint             int `koanf:"keypoint"`
	Ocr                  int `koanf:"ocr"`
	InstanceSegmentation int `koanf:"instancesegmentation"`
	SemanticSegmentation int `koanf:"semanticsegmentation"`
	TextGeneration       int `koanf:"textgeneration"`
}

// TemporalConfig related to Temporal
type TemporalConfig struct {
	HostPort   string `koanf:"hostport"`
	Namespace  string `koanf:"namespace"`
	Ca         string `koanf:"ca"`
	Cert       string `koanf:"cert"`
	Key        string `koanf:"key"`
	ServerName string `koanf:"servername"`
}

// LogConfig related to logging
type LogConfig struct {
	External      bool `koanf:"external"`
	OtelCollector struct {
		Host string `koanf:"host"`
		Port string `koanf:"port"`
	}
}

// AppConfig defines
type AppConfig struct {
	Server                 ServerConfig       `koanf:"server"`
	Database               DatabaseConfig     `koanf:"database"`
	TritonServer           TritonServerConfig `koanf:"tritonserver"`
	MgmtBackend            MgmtBackendConfig  `koanf:"mgmtbackend"`
	Cache                  CacheConfig        `koanf:"cache"`
	MaxBatchSizeLimitation MaxBatchSizeConfig `koanf:"maxbatchsizelimitation"`
	Temporal               TemporalConfig     `koanf:"temporal"`
	Controller             ControllerConfig   `koanf:"controller"`
	InitModel              InitModelConfig    `koanf:"initmodel"`
	Log                    LogConfig          `koanf:"log"`
}

// Config - Global variable to export
var Config AppConfig

// Init - Assign global config to decoded config struct
func Init() error {
	k := koanf.New(".")
	parser := yaml.Parser()

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fileRelativePath := fs.String("file", "config/config.yaml", "configuration file")
	flag.Parse()

	if err := k.Load(file.Provider(*fileRelativePath), parser); err != nil {
		log.Fatal(err.Error())
	}

	if err := k.Load(env.ProviderWithValue("CFG_", ".", func(s string, v string) (string, interface{}) {
		key := strings.Replace(strings.ToLower(strings.TrimPrefix(s, "CFG_")), "_", ".", -1)
		if strings.Contains(v, ",") {
			return key, strings.Split(strings.TrimSpace(v), ",")
		}
		return key, v
	}), nil); err != nil {
		return err
	}

	if err := k.Unmarshal("", &Config); err != nil {
		return err
	}

	return ValidateConfig(&Config)
}

// ValidateConfig is for custom validation rules for the configuration
func ValidateConfig(cfg *AppConfig) error {
	return nil
}
