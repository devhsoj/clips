package config

import (
	"github.com/joho/godotenv"
	"os"
	"path"
	"strconv"
)

var HTTPS bool
var TLSCertPath string
var TLSKeyPath string
var ListenAddress string
var PostgresURL string

var ApplicationPath string

var BcryptCost int = 12 // Default to 12, so we don't have slow signups/logins
var BodyLimit int = 100 * 1024 * 1024  // Default to 100 MB

var ClipSavePath string
var AllowSignup bool = true
var AllowUpload bool = true

// LoadConfig loads the environment variables from .env into the above variables
func LoadConfig() error {
	var err error

	if err = godotenv.Load(); err != nil {
		return err
	}

	HTTPS = os.Getenv("HTTPS") == "true"

	if HTTPS {
		TLSCertPath = os.Getenv("TLS_CERT")
		TLSKeyPath = os.Getenv("TLS_KEY")
	}

	ListenAddress = os.Getenv("LISTEN_ADDRESS")

	if ListenAddress == "" {
		ListenAddress = "0.0.0.0:3000"
	}

	PostgresURL = os.Getenv("POSTGRES_URL")

	if PostgresURL == "" {
		PostgresURL = "postgres://postgres:postgres@localhost/clips?sslmode=disable"
	}

	if os.Getenv("BCRYPT_COST") != "" {
		BcryptCost,err = strconv.Atoi(os.Getenv("BCRYPT_COST"))

		if err != nil {
			return err
		}
	}

	if os.Getenv("BODY_LIMIT_MB") != "" {
		BodyLimit,err = strconv.Atoi(os.Getenv("BODY_LIMIT_MB"))

		BodyLimit = BodyLimit * 1024 * 1024

		if err != nil {
			return err
		}
	}

	ClipSavePath = os.Getenv("CLIP_SAVE_PATH")

	AllowSignup = os.Getenv("ALLOW_SIGNUP") == "true" || os.Getenv("ALLOW_SIGNUP") == ""
	AllowUpload = os.Getenv("ALLOW_UPLOAD") == "true" || os.Getenv("ALLOW_UPLOAD") == ""

	ApplicationPath,err = os.Getwd()

	if err != nil {
		return err
	}

	if ClipSavePath == "" {
		ClipSavePath,err = CreateClipStorageDirectory()
	}

	return err
}


func CreateClipStorageDirectory() (string,error) {

	storagePath := path.Join(ApplicationPath,"build/","uploads/")

	if err := os.Mkdir(storagePath,0777); err != nil {
		return "",err
	}

	return storagePath,nil
}