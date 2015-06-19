package drivers

import (
	"errors"
	"fmt"
	"sort"
	"os/exec"

	"github.com/paz-sh/paz.sh/ssh"
	"github.com/paz-sh/paz.sh/command"
)

// Define what a driver is, much like in docker/machine, but with a new
// take
type Driver interface {
	
	// Create a host
	Create() error

	// The name of the Driver
	DriverName() string

	SetConfigFromFlags(flags DriverOptions) error
	GetCommand() *cli.Command

	// GetIp
	GetIP() (string, error)

	// Gets the machines name, 
	GetMachineName() (string, error)

	GetFleetUrl() (string, error)

	GetFleetPort() (int, error)

	// GetProviderType returns whether the instance is local/remote
	GetProviderType() int

	// GetSSHHostname returns hostname for use with ssh
	GetSSHHostname() (string, error)

	//
	// Methods that help with ssh and all.
	//

	// GetSSHKeyPath returns key path for use with ssh
	GetSSHKeyPath() string

	// GetSSHPort returns port for use with ssh
	GetSSHPort() (int, error)

	// GetSSHUsername returns username for use with ssh
	GetSSHUsername() string

	//
	// Some methods to handle machines
	//

	// Start a host
	Start() error

	// Stop a host gracefully
	Stop() error

	// Restart a host. This may just call Stop(); Start() if the provider does not
	// have any special restart behaviour.
	Restart() error

	// Remove a host
	Remove() error

	// Kill stops a host forcefully
	Kill() error

}

type DriverOptions interface {
	String(key string) string
	Int(key string) int
	Bool(key string) bool
}

// From docker/machine
// RegisteredDriver is used to register a driver with the Register function.
// New: a function that returns a new driver given a path to store host
//   configuration in
type RegisteredDriver struct {
	New            func(machineName string) (Driver, error)
}

var ErrHostIsNotRunning = errors.New("host is not running")

var (
	drivers map[string]*RegisteredDriver
)

func init() {
	drivers = make(map[string]*RegisteredDriver)
}

// Register a driver
func Register(name string, registeredDriver *RegisteredDriver) error {
	if _, exists := drivers[name]; exists {
		return fmt.Errorf("Name already registered %s", name)
	}

	drivers[name] = registeredDriver
	return nil
}

// NewDriver creates a new driver of type "name"
func NewDriver(name string, machineName string) (Driver, error) {
	driver, exists := drivers[name]
	if !exists {
		return nil, fmt.Errorf("hosts: Unknown driver %q", name)
	}
	return driver.New(machineName)
}

// GetDriverNames returns a slice of all registered driver names
func GetDriverNames() []string {
	names := make([]string, 0, len(drivers))
	for k := range drivers {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

func GetSSHCommandFromDriver(d Driver, args ...string) (*exec.Cmd, error) {
	host, err := d.GetSSHHostname()
	if err != nil {
		return nil, err
	}

	port, err := d.GetSSHPort()
	if err != nil {
		return nil, err
	}

	user := d.GetSSHUsername()
	keyPath := d.GetSSHKeyPath()

	return ssh.GetSSHCommand(host, port, user, keyPath, args...), nil
}
