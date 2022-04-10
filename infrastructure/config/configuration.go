package config

import "os"

var (
	Tables = []string{}
)

type ServerConfig struct {
	ServerAddress string
	ServerPort    string
	debug         bool
	AccessSecret  string
	RefreshSecret string
	MaxFileSize   uint64
	dbConfig      *databaseConfig
	cacheConfig   *cacheConfig
	keyConfig     *keyConfig
	tokenConfig   *tokenConfig
}

func NewServerConfig() *ServerConfig {
	serverConfig := ServerConfig{}
	if os.Getenv("eazyweigh_server") != "" || len(os.Getenv("eazyweigh_server")) != 0 {
		serverConfig.ServerAddress = os.Getenv("eazyweigh_server")
	} else {
		serverConfig.ServerAddress = defaultsettings["server_address"]
	}
	if os.Getenv("eazyweigh_server_port") != "" || len(os.Getenv("eazyweigh_server_port")) != 0 {
		serverConfig.ServerPort = os.Getenv("eazyweigh_server_port")
	} else {
		serverConfig.ServerPort = defaultsettings["server_port"]
	}
	serverConfig.debug = debug
	serverConfig.dbConfig = NewDatabaseConfig()
	serverConfig.cacheConfig = NewCacheConfig()
	serverConfig.keyConfig = NewKeyConfig()
	serverConfig.tokenConfig = NewTokenConfig()
	return &serverConfig
}

func (conf *ServerConfig) IsDebug() bool {
	return conf.debug
}

func (conf *ServerConfig) GetDatabaseConfig() *databaseConfig {
	return conf.dbConfig
}

func (conf *ServerConfig) GetCacheConfig() *cacheConfig {
	return conf.cacheConfig
}

func (conf *ServerConfig) GetKeyConfig() *keyConfig {
	return conf.keyConfig
}

func (conf *ServerConfig) GetTokenConfig() *tokenConfig {
	return conf.tokenConfig
}

type databaseConfig struct {
	DbHost            string
	DbPort            string
	DbName            string
	DbUser            string
	DbPassword        string
	WarehouseHost     string
	WarehouseDBName   string
	WarehouseUser     string
	WarehousePassword string
}

func NewDatabaseConfig() *databaseConfig {
	dbConf := databaseConfig{}
	if os.Getenv("eazyweigh_database_server") != "" || len(os.Getenv("eazyweigh_database_server")) != 0 {
		dbConf.DbHost = os.Getenv("eazyweigh_database_server")
	} else {
		dbConf.DbHost = defaultsettings["db_host"]
	}
	if os.Getenv("eazyweigh_database_server_port") != "" || len(os.Getenv("eazyweigh_database_server_port")) != 0 {
		dbConf.DbPort = os.Getenv("eazyweigh_database_server_port")
	} else {
		dbConf.DbPort = defaultsettings["db_port"]
	}
	if os.Getenv("mysql_username") != "" || len(os.Getenv("mysql_username")) != 0 {
		dbConf.DbUser = os.Getenv("mysql_username")
	} else {
		dbConf.DbUser = defaultsettings["db_user"]
	}
	if os.Getenv("mysql_password") != "" || len(os.Getenv("mysql_password")) != 0 {
		dbConf.DbPassword = os.Getenv("mysql_password")
	} else {
		dbConf.DbPassword = defaultsettings["db_pass"]
	}
	if os.Getenv("eazyweigh_database_name") != "" || len(os.Getenv("eazyweigh_database_name")) != 0 {
		dbConf.DbName = os.Getenv("eazyweigh_database_name")
	} else {
		dbConf.DbName = defaultsettings["db_name"]
	}
	if os.Getenv("warehouse_database_server") != "" || len(os.Getenv("warehouse_database_server")) != 0 {
		dbConf.WarehouseHost = os.Getenv("warehouse_database_server")
	} else {
		dbConf.WarehouseHost = "localhost"
	}
	if os.Getenv("warehouse_username") != "" || len(os.Getenv("warehouse_username")) != 0 {
		dbConf.WarehouseUser = os.Getenv("warehouse_username")
	} else {
		dbConf.WarehouseUser = ""
	}
	if os.Getenv("warehouse_password") != "" || len(os.Getenv("warehouse_password")) != 0 {
		dbConf.WarehousePassword = os.Getenv("warehouse_password")
	} else {
		dbConf.WarehousePassword = ""
	}
	if os.Getenv("warehouse_database_name") != "" || len(os.Getenv("warehouse_database_name")) != 0 {
		dbConf.WarehouseDBName = os.Getenv("warehouse_database_name")
	} else {
		dbConf.WarehouseDBName = ""
	}
	return &dbConf
}

type cacheConfig struct {
	CacheHost     string
	CachePort     string
	CachePassword string
}

func NewCacheConfig() *cacheConfig {
	cacheConf := cacheConfig{}
	if os.Getenv("cache_server") != "" || len(os.Getenv("cache_server")) != 0 {
		cacheConf.CacheHost = os.Getenv("cache_server")
	} else {
		cacheConf.CacheHost = defaultsettings["cache_host"]
	}
	if os.Getenv("cache_port") != "" || len(os.Getenv("cache_port")) != 0 {
		cacheConf.CachePort = os.Getenv("cache_port")
	} else {
		cacheConf.CachePort = defaultsettings["cache_port"]
	}
	if os.Getenv("cache_password") != "" || len(os.Getenv("cache_password")) != 0 {
		cacheConf.CachePassword = os.Getenv("cache_password")
	} else {
		cacheConf.CachePassword = defaultsettings["db_password"]
	}
	return &cacheConf
}

type keyConfig struct {
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
}

func NewKeyConfig() *keyConfig {
	return &keyConfig{
		AccessTokenPrivateKeyPath:  "./access-private.pem",
		AccessTokenPublicKeyPath:   "./access-public.pem",
		RefreshTokenPrivateKeyPath: "./refresh-private.pem",
		RefreshTokenPublicKeyPath:  "./refresh-public.pem",
	}
}

type tokenConfig struct {
	JWTExpiration     int // in minutes
	RefreshExpiration int // in days
}

func NewTokenConfig() *tokenConfig {
	return &tokenConfig{
		JWTExpiration:     1,
		RefreshExpiration: 7,
	}
}
