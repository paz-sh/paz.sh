package vagrant

import (
	"github.com/paz-sh/paz.sh/machine/drivers"
	"github.com/paz-sh/paz.sh/command"
	"github.com/paz-sh/paz.sh/log"
	// "github.com/paz-sh/paz.sh/ssh"
)

type Driver struct {
	MachineName       string
	UserData          string
	SSHUser           string
	SSHPort           int
	Command           string
}

func init() {
	drivers.Register("vagrant", &drivers.RegisteredDriver{
		New:            NewDriver,
	})
}

func (d *Driver) DriverName() string {
	return "vagrant"
}

func (d *Driver) GetCommand() *cli.Command {
	var cmd =  &cli.Command{
		Name:    d.DriverName(),
		Summary: "Provision instances of Paz on vagrant",
		Usage:   "",
		Description: ``,
	}

	cmd.Run = d.Help
	cmd.Flags.StringVar(&d.UserData, "user-data", "https://raw.githubusercontent.com/paz-sh/paz/master/vagrant/user-data", "Paz's user-data file for vagrant")
	return cmd
}

func (d *Driver) Help (args []string) (exit int) {
	return 0
}

func (d *Driver) GetIP() (string, error) {
	return "0.0.0.0", nil
}

func (d *Driver) GetProviderType() int {
	return 0
}

func (d *Driver) GetFleetPort() (int, error) {
	return 0, nil
}

func (d *Driver) GetMachineName() (string, error) {
	return "paz", nil
}

func (d *Driver) GetFleetUrl() (string, error) {
	return "http://paz.sh", nil
}

func (d *Driver) GetSSHHostname() (string, error) {
	return "paz.sh", nil
}

func (d *Driver) GetSSHPort() (int, error) {
	return 22, nil
}

func (d *Driver) GetSSHUsername() string {
	return "core"
}

func (d *Driver) GetSSHKeyPath() string {
	return "~/.ssh/id_rsa.pub"
}

func NewDriver(machineName string) (drivers.Driver, error) {
	return &Driver{MachineName: machineName}, nil
}

func (d *Driver) Start() error {
	return nil
}

func (d *Driver) Stop() error {
	return nil
}

func (d *Driver) Restart() error {
	return nil
}

func (d *Driver) Remove() error {
	return nil
}

func (d *Driver) Kill() error {
	return nil
}

func (d *Driver) getClient() (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (d *Driver) Create() error {
	// log.Infof("Creating SSH key...")

	// key, err := d.createSSHKey()
	// if err != nil {
	// 	return err
	// }

	// d.SSHKeyID = key.ID

	log.Infof("Creating Vagrant droplet...")

	client, err := d.getClient()
	if err != nil {
		return nil
	}

	if err := client.vagrant("up"); err != nil {
		return err
	}

	// Maybe use fleetssh instead?
	// log.Infof("Waiting for SSH...")

	// if err := ssh.WaitForTCP(fmt.Sprintf("%s:%d", d.IPAddress, 22)); err != nil {
	// 	return err
	// }

	log.Info("Configuring Machine...")

	// log.Debugf("Setting hostname: %s", d.MachineName)
	// cmd, err := d.GetSSHCommand(fmt.Sprintf(
	// 	"echo \"127.0.0.1 %s\" | sudo tee -a /etc/hosts && sudo hostname %s && echo \"%s\" | sudo tee /etc/hostname",
	// 	d.MachineName,
	// 	d.MachineName,
	// 	d.MachineName,
	// ))

	// if err != nil {
	// 	return err
	// }
	// if err := cmd.Run(); err != nil {
	// 	return err
	// }

	return nil
}
