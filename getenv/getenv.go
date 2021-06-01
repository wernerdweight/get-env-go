package getenv

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const (
	defaultEnvFile = ".env.local"

	NoValue                = ""
	CwdFailureError        = 1
	CantAccessEnvFileError = 2
	NoEnvFileError         = 3
	CantLoadEnvFileError   = 4
	NoSuchEnvVarError      = 5
)

var errorMap = map[int]string{
	CwdFailureError:        "can not determine current working directory",
	CantAccessEnvFileError: "error accessing an existing env file '%s'",
	NoEnvFileError:         "no %s file exists in any of parent directories",
	CantLoadEnvFileError:   "can't load %s file",
	NoSuchEnvVarError:      "non-existing ENV variable %s requested",
}

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
		Err:  errors.New(fmt.Sprintf(errorMap[code], args...)),
	}
}

func getCwd() (error, string) {
	cwd, err := os.Getwd()
	if nil != err {
		return createError(CwdFailureError), NoValue
	}
	return nil, cwd
}

func pathExists(path string) (error, bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false
		}
		return createError(CantAccessEnvFileError, path), false
	}
	return nil, true
}

func getParentPath(path string) string {
	return regexp.MustCompile("^(.*/)([^/]+)/([^/]+)$").ReplaceAllString(path, "$1$3")
}

func getEnvFilePath(filename string) (error, string) {
	var err, cwd = getCwd()
	if nil != err {
		return err, NoValue
	}
	var path = fmt.Sprintf("%s/%s", cwd, filename)
	for true {
		if _, exists := pathExists(path); true == exists {
			break
		}
		parentPath := getParentPath(path)
		if parentPath == path {
			return createError(NoEnvFileError, filename), NoValue
		}
		path = parentPath
	}
	return nil, path
}

func initialize(filename string) error {
	envErr, envFile := getEnvFilePath(filename)
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
func GetEnv(key string) (error, string) {
	value, exists := os.LookupEnv(key)
	if true != exists {
		return createError(NoSuchEnvVarError, key), NoValue
	}
	return nil, value
}
