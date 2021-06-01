package getenv_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wernerdweight/get-env-go/getenv"
	"testing"
)

func TestDefaultFile(t *testing.T) {
	assertion := assert.New(t)
	err := getenv.Init()
	assertion.NotNil(err)
	assertion.Equal(getenv.NoEnvFileError, err.(*getenv.Error).Code)

	err1, value1 := getenv.GetEnv("ENV_VAR_1")
	assertion.NotNil(err1)
	assertion.Equal(getenv.NoSuchEnvVarError, err1.(*getenv.Error).Code)
	assertion.Equal("", value1)

	err2, value2 := getenv.GetEnv("ENV_VAR_2")
	assertion.NotNil(err2)
	assertion.Equal(getenv.NoSuchEnvVarError, err2.(*getenv.Error).Code)
	assertion.Equal("", value2)

	err3, value3 := getenv.GetEnv("ENV_VAR_3")
	assertion.NotNil(err3)
	assertion.Equal(getenv.NoSuchEnvVarError, err3.(*getenv.Error).Code)
	assertion.Equal("", value3)
}

func TestCustomFile(t *testing.T) {
	assertion := assert.New(t)
	assertion.Nil(getenv.InitFromFile(".env.test"))

	err1, value1 := getenv.GetEnv("ENV_VAR_1")
	assertion.Nil(err1)
	assertion.Equal("value-5", value1)

	err2, value2 := getenv.GetEnv("ENV_VAR_2")
	assertion.Nil(err2)
	assertion.Equal("value-6", value2)

	err3, value3 := getenv.GetEnv("ENV_VAR_3")
	assertion.NotNil(err3)
	assertion.Equal(getenv.NoSuchEnvVarError, err3.(*getenv.Error).Code)
	assertion.Equal("", value3)
}

func TestNoFile(t *testing.T) {
	assertion := assert.New(t)
	err := getenv.InitFromFile("non-existing-file")
	assertion.NotNil(err)
	assertion.Equal(getenv.NoEnvFileError, err.(*getenv.Error).Code)
}
