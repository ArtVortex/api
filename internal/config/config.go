package config

import (
	"artvortex-api/types"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

// in current code, this function is not being used
// Config func to get env value from key ---
func LoadConfig(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

// Default configurations for the NFT service
// Should be replaced by environment variables
var (
	DefaultDBHost     = "postgres"
	DefaultDBPort     = 5432
	DefaultDBUser     = "pguser"
	DefaultDBName     = "artvortex"
	DefaultDBPassword = "nft2023"
	DefaultHost       = "http://localhost"
	DefaultPort       = "8000"
	DefaultIPFSHost   = "ipfs"
	DefaultIPFSPort   = "5001"
)

// this function can not be called twice as it results error of redefining flags
// Parse the configurations from console flag or from the system environment variables.
func makeConfig() *types.NFTServiceConfig {
	var dbHost string
	var dbPort int
	var dbUser string
	var dbName string
	var dbPassword string
	var host string
	var port string
	var ipfsHost string
	var ipfsPort string

	flag.StringVar(&dbHost, "db_host", DefaultDBHost, "The host name or IP address of the database service")
	flag.IntVar(&dbPort, "db_port", DefaultDBPort, "The port of the database service")
	flag.StringVar(&dbUser, "db_user", DefaultDBUser, "The user of the database service")
	flag.StringVar(&dbName, "db_name", DefaultDBName, "The database name")
	flag.StringVar(&dbPassword, "db_password", DefaultDBPassword, "The password of the database service")
	flag.StringVar(&host, "host", DefaultHost, "The IP address of application")
	flag.StringVar(&port, "port", DefaultPort, "The IP address of application")
	flag.StringVar(&ipfsHost, "ipfs_host", DefaultIPFSHost, "The IP address of IPFS host")
	flag.StringVar(&ipfsPort, "ipfs_port", DefaultIPFSPort, "The IP address of IPFS port")

	flag.Parse()

	return &types.NFTServiceConfig{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBName:     dbName,
		DBPassword: dbPassword,
		Host:       host,
		Port:       port,
		IPFSHost:   ipfsHost,
		IPFSPort:   ipfsPort,
	}
}

func InitializeConfig() {
	NFTServiceConfig = makeConfig()
}

// global config
var NFTServiceConfig *types.NFTServiceConfig
