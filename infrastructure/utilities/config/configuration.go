package config

import "os"

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
	if os.Getenv("unza_server") != "" || len(os.Getenv("unza_server")) != 0 {
		serverConfig.ServerAddress = os.Getenv("unza_server")
	} else {
		serverConfig.ServerAddress = "localhost"
	}
	if os.Getenv("unza_server_port") != "" || len(os.Getenv("unza_server_port")) != 0 {
		serverConfig.ServerPort = os.Getenv("unza_server_port")
	} else {
		serverConfig.ServerPort = "8000"
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
	DbHost      string
	DbPort      string
	DbName      string
	DbUser      string
	DbPassword  string
	SQLHost     string
	SQLDBName   string
	SQLUser     string
	SQLPassword string
}

func NewDatabaseConfig() *databaseConfig {
	dbConf := databaseConfig{}
	if os.Getenv("unza_database_server") != "" || len(os.Getenv("unza_database_server")) != 0 {
		dbConf.DbHost = os.Getenv("unza_database_server")
	} else {
		dbConf.DbHost = "localhost"
	}
	if os.Getenv("unza_database_server_port") != "" || len(os.Getenv("unza_database_server_port")) != 0 {
		dbConf.DbPort = os.Getenv("unza_database_server_port")
	} else {
		dbConf.DbPort = "3306"
	}
	if os.Getenv("mysql_username") != "" || len(os.Getenv("mysql_username")) != 0 {
		dbConf.DbUser = os.Getenv("mysql_username")
	} else {
		dbConf.DbUser = ""
	}
	if os.Getenv("mysql_password") != "" || len(os.Getenv("mysql_password")) != 0 {
		dbConf.DbPassword = os.Getenv("mysql_password")
	} else {
		dbConf.DbPassword = ""
	}
	if os.Getenv("unza_database_name") != "" || len(os.Getenv("unza_database_name")) != 0 {
		dbConf.DbName = os.Getenv("unza_database_name")
	} else {
		dbConf.DbName = "unza"
	}
	if os.Getenv("sql_database_server") != "" || len(os.Getenv("sql_database_server")) != 0 {
		dbConf.SQLHost = os.Getenv("sql_database_server")
	} else {
		dbConf.SQLHost = "localhost"
	}
	if os.Getenv("sql_username") != "" || len(os.Getenv("sql_username")) != 0 {
		dbConf.SQLUser = os.Getenv("sql_username")
	} else {
		dbConf.SQLUser = ""
	}
	if os.Getenv("sql_password") != "" || len(os.Getenv("sql_password")) != 0 {
		dbConf.SQLPassword = os.Getenv("sql_password")
	} else {
		dbConf.SQLPassword = ""
	}
	if os.Getenv("sql_database_name") != "" || len(os.Getenv("sql_database_name")) != 0 {
		dbConf.SQLDBName = os.Getenv("sql_database_name")
	} else {
		dbConf.SQLDBName = "unza"
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
		cacheConf.CacheHost = "localhost"
	}
	if os.Getenv("cache_port") != "" || len(os.Getenv("cache_port")) != 0 {
		cacheConf.CachePort = os.Getenv("cache_port")
	} else {
		cacheConf.CachePort = "6379"
	}
	if os.Getenv("cache_password") != "" || len(os.Getenv("cache_password")) != 0 {
		cacheConf.CachePassword = os.Getenv("cache_password")
	} else {
		cacheConf.CachePassword = "rishabh2791"
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
		JWTExpiration:     10 * 2,
		RefreshExpiration: 7,
	}
}
