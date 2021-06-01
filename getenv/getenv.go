package getenv

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const (
	defaultEnvFile = ".env.local"

	// NoValue is used if requested ENV var doesn't exist
	NoValue = ""
	// CwdFailureError signals that the current working directory can not be determined
	CwdFailureError = 1
	// CantAccessEnvFileError signals that the env file exists but can not be accessed
	CantAccessEnvFileError = 2
	// NoEnvFileError signals that no env file exists in any of the parent directories
	NoEnvFileError = 3
	// CantLoadEnvFileError signals that the env file exists and can be accessed but not parsed and exported
	CantLoadEnvFileError = 4
	// NoSuchEnvVarError signals that the requested ENV var doesn't exist
	NoSuchEnvVarError = 5
)

var errorMap = map[int]string{
	CwdFailureError:        "can not determine current working directory",
	CantAccessEnvFileError: "error accessing an existing env file '%s'",
	NoEnvFileError:         "no %s file exists in any of parent directories",
	CantLoadEnvFileError:   "can't load %s file",
	NoSuchEnvVarError:      "non-existing ENV variable %s requested",
}

// Error is a custom error type which provides useful error codes
type Error struct {
	Code int
	Err  error
}

func (err *Error) Error() string {
	return fmt.Sprintf("#%d: %v", err.Code, err.Err)
}

func createError(code int, args ...interface{}) *Error {
	return &Error{
		Code: code,
		Err:  fmt.Errorf(errorMap[code], args...),
	}
}

func getCwd() (string, error) {
	cwd, err := os.Getwd()
	if nil != err {
		return NoValue, createError(CwdFailureError)
	}
	return cwd, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, createError(CantAccessEnvFileError, path)
	}
	return true, nil
}

func getParentPath(path string) string {
	return regexp.MustCompile("^(.*/)([^/]+)/([^/]+)$").ReplaceAllString(path, "$1$3")
}

func getEnvFilePath(filename string) (string, error) {
	var cwd, err = getCwd()
	if nil != err {
		return NoValue, err
	}
	var path = fmt.Sprintf("%s/%s", cwd, filename)
	for true {
		if exists, _ := pathExists(path); true == exists {
			break
		}
		parentPath := getParentPath(path)
		if parentPath == path {
			return NoValue, createError(NoEnvFileError, filename)
		}
		path = parentPath
	}
	return path, nil
}

func initialize(filename string) error {
	envFile, envErr := getEnvFilePath(filename)
	if nil != envErr {
		return envErr
	}
	err := godotenv.Load(envFile)
	if nil != err {
		return createError(CantLoadEnvFileError, filename)
	}
	return nil
}

// InitFromFile looks for custom ENV file and extracts ENV variables from it.
func InitFromFile(filename string) error {
	return initialize(filename)
}

// Init looks for ENV file and extracts ENV variables from it.
func Init() error {
	return initialize(defaultEnvFile)
}

// GetEnv returns requested ENV variable.
func GetEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if true != exists {
		return NoValue, createError(NoSuchEnvVarError, key)
	}
	return value, nil
}
