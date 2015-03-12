package host

import (
	"errors"
	"regexp"
	
	"github.com/paz-sh/paz.sh/machine/drivers"
)

type Host struct {
	Name           string `json:"-"`
	DriverName     string
	Driver         drivers.Driver
	CaCertPath     string
	ServerCertPath string
	ServerKeyPath  string
	PrivateKeyPath string
	ClientCertPath string
}

var (
	validHostNameChars       = `[a-zA-Z0-9\-\.]`
	validHostNamePattern     = regexp.MustCompile(`^` + validHostNameChars + `+$`)
	ErrInvalidHostname       = errors.New("Invalid hostname specified")
	ErrUnknownHypervisorType = errors.New("Unknown hypervisor type")
)

func NewHost(name string, driverName string) (*Host, error) {
	driver, err := drivers.NewDriver(driverName, name)
	if err != nil {
		return nil, err
	}
	return &Host{
		Name:           name,
		DriverName:     driverName,
		Driver:         driver,
	}, nil
}

func (h *Host) Start() error {
	return h.Driver.Start()
}

func (h *Host) Stop() error {
	return h.Driver.Stop()
}

func (h *Host) Remove(force bool) error {
	if err := h.Driver.Remove(); err != nil {
		if !force {
			return err
		}
	}
	return nil
}

func ValidateHostName(name string) (string, error) {
	if !validHostNamePattern.MatchString(name) {
		return name, ErrInvalidHostname
	}
	return name, nil
}

func (h *Host) Create(name string) error {
	name, err := ValidateHostName(name)
	if err != nil {
		return err
	}

	// create the instance
	if err := h.Driver.Create(); err != nil {
		return err
	}

	return nil
}
