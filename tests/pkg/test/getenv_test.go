package getenv_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wernerdweight/get-env-go/getenv"
	"testing"
)

func TestDefaultFile(t *testing.T) {
	assertion := assert.New(t)
	assertion.Nil(getenv.Init())

	err1, value1 := getenv.GetEnv("ENV_VAR_1")
	assertion.Nil(err1)
	assertion.Equal("value-1", value1)

	err2, value2 := getenv.GetEnv("ENV_VAR_2")
	assertion.Nil(err2)
	assertion.Equal("value-2", value2)

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
