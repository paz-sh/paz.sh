package vagrant

import (
	"bytes"
	"fmt"
	"errors"
	"os/exec"
	"strings"

	"github.com/paz-sh/paz.sh/log"
)

var (
	ErrMachineExist    = errors.New("machine already exists")
	ErrMachineNotExist = errors.New("machine does not exist")
	ErrVagrantNotFound = errors.New("Vagrant not found")
)

type Client struct {
	Command          string
}

func (c *Client) vagrant(args ...string) error {
	_, _, err := c.vagrantErr(args...)
	return err
}

func (c *Client) vagrantOut(args ...string) (string, error) {
	stdout, _, err := c.vagrantErr(args...)
	return stdout, err
}

func (c *Client) vagrantErr(args ...string) (string, string, error) {

	cmd := exec.Command(c.Command, args...)
	log.Debugf("executing: %v %v", c.Command, strings.Join(args, " "))
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	stderrStr := stderr.String()
	log.Debugf("STDOUT: %v", stdout.String())
	log.Debugf("STDERR: %v", stderrStr)
	if err != nil {
		if ee, ok := err.(*exec.Error); ok && ee == exec.ErrNotFound {
			err = ErrVagrantNotFound
		}
	} else {
		if strings.Contains(stderrStr, "error:") {
			err = fmt.Errorf("%v %v failed: %v", c.Command, strings.Join(args, " "), stderrStr)
		}
	}
	return stdout.String(), stderrStr, err

}

func getVagrantPath() string {

	var cmd = "vagrant"

	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}

	return ""
}

func newClient() (*Client, error) {
	var path = getVagrantPath()
	return &Client{Command: path}, nil
}
