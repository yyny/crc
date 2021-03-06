package shell

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	defer func(shell string) { os.Setenv("SHELL", shell) }(os.Getenv("SHELL"))
	os.Setenv("SHELL", "")

	shell, err := detect()

	assert.Contains(t, supportedShell, shell)
	assert.NoError(t, err)
}

func TestGetNameAndItsPpidOfCurrent(t *testing.T) {
	shell, shellppid, err := getNameAndItsPpid(os.Getpid())

	assert.Equal(t, "shell.test.exe", shell)
	assert.Equal(t, os.Getppid(), shellppid)
	assert.NoError(t, err)
}

func TestGetNameAndItsPpidOfParent(t *testing.T) {
	shell, _, err := getNameAndItsPpid(os.Getppid())

	assert.Equal(t, "go.exe", shell)
	assert.NoError(t, err)
}
