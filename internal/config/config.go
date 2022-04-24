package config

import (
	"github.com/joho/godotenv"
	"os"
	"path"
	"strconv"
)

var HTTPS bool // Whether to use https.
var TLSCertPath string // The path to the SSL/TLS certificate.
var TLSKeyPath string // The path to the SSL/TLS key.
var ListenAddress string // The address for the HTTP/HTTPS server to listen on.
var PostgresURL string // The PostgreSQL connection URL used for all database operations.

var ApplicationPath string // The current working directory of the running executable.

var BcryptCost int = 12 // The cost factor used in Bcrypt. Default to 12.
var BodyLimit int = 100 * 1024 * 1024  // The body limit of all HTTP/HTTPS requests. Default to 100 MB

var ClipSavePath string // The path to the directory used to store user uploaded clips.
var AllowSignup bool = true // Whether to allow signups.
var AllowUpload bool = true // Whether to allow uploads.

// LoadConfig loads the environment variables from .env into global config variables.
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

// CreateClipStorageDirectory creates a directory used to store user uploads if no directory was already specified.
func CreateClipStorageDirectory() (string,error) {

	storagePath := path.Join(ApplicationPath,"build/","uploads/")

	if err := os.Mkdir(storagePath,0777); err != nil {
		return "",err
	}

	return storagePath,nil
}