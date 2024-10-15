// Package configs provides utility functions to load and retrieve environment variables. These variables configure the application, such as database URIs, login credentials,
// server ports, and limits. It ensures proper defaults or errors if environment variables are missing.
package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/UTDNebula/nebula-api/api/common/log"

	_ "github.com/joho/godotenv/autoload"
)

// GetPortString retrieves the port number for the application from the environment variables and returns the port number as a string in the format
// ":<port>" to be used when configuring the server, if the "Port" environment variable is not found, it defaults to port 8080.
func GetPortString() string {

	portNumber, exist := os.LookupEnv("Port")
	if !exist {
		portNumber = "8080"
	}

	portString := fmt.Sprintf(":%s", portNumber)

	return portString
}

// GetEnvMongoURI retrieves the MongoDB URI from the environment variables if the "MONGODB_URI" variable is missing, it logs an error and terminates the program.
func GetEnvMongoURI() string {

	uri, exist := os.LookupEnv("MONGODB_URI")
	if !exist {
		log.WriteErrorMsg("Error loading 'MONGODB_URI' from the .env file")
		os.Exit(1) // Exit the program if the MongoDB URI is missing
	}

	return uri
}

// GetEnvLogin retrieves the login credentials (NetID and password) from the environment variables and it returns both them as strings for use in authentication.
// If either "LOGIN_NETID" or "LOGIN_PASSWORD" is missing, it logs an error message and terminates.
func GetEnvLogin() (netID string, password string) {

	netID, exist := os.LookupEnv("LOGIN_NETID")
	if !exist {
		log.WriteErrorMsg("Error loading 'LOGIN_NETID' from the .env file")
		os.Exit(1) // Exit the program if NetID is missing
	}
	password, exist = os.LookupEnv("LOGIN_PASSWORD")
	if !exist {
		log.WriteErrorMsg("Error loading 'LOGIN_PASSWORD' from the .env file")
		os.Exit(1) // Exit the program if password is missing
	}

	return netID, password
}

// GetEnvLimit retrieves the limit for a specific configuration from the environment variables.
// If the "LIMIT" environment variable is not found or is not a valid integer, it defaults to a limit of 20 and returns it.
func GetEnvLimit() int64 {

	const defaultLimit int64 = 20

	limitString, exist := os.LookupEnv("LIMIT")
	if !exist {
		return defaultLimit // Return default if LIMIT is not set
	}

	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		return defaultLimit // Return default if the value is not a valid integer
	}

	return limit
}
