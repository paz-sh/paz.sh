package digitalocean

import (
	"time"

	"github.com/digitalocean/godo"
	"code.google.com/p/goauth2/oauth"
	"github.com/paz-sh/paz.sh/machine/drivers"
	"github.com/paz-sh/paz.sh/log"
	// "github.com/paz-sh/paz.sh/ssh"
)

type Driver struct {
	AccessToken       string
	DropletID         int
	DropletName       string
	Image             string
	MachineName       string
	IPAddress         string
	Region            string
	SSHKeyID          int
	SSHUser           string
	SSHPort           int
	Size              string
	IPv6              bool
	Backups           bool
	PrivateNetworking bool
	CaCertPath        string
	PrivateKeyPath    string
	DriverKeyPath     string
}

func init() {
	drivers.Register("digitalocean", &drivers.RegisteredDriver{
		New:            NewDriver,
	})
}

func (d *Driver) DriverName() error {
	return nil
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

func (d *Driver) getClient() *godo.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: d.AccessToken},
	}

	return godo.NewClient(t.Client())
}

func (d *Driver) Create() error {
	// log.Infof("Creating SSH key...")

	// key, err := d.createSSHKey()
	// if err != nil {
	// 	return err
	// }

	// d.SSHKeyID = key.ID

	log.Infof("Creating Digital Ocean droplet...")

	client := d.getClient()

	createRequest := &godo.DropletCreateRequest{
		Image:             d.Image,
		Name:              d.MachineName,
		Region:            d.Region,
		Size:              d.Size,
		IPv6:              d.IPv6,
		PrivateNetworking: d.PrivateNetworking,
		Backups:           d.Backups,
		SSHKeys:           []interface{}{d.SSHKeyID},
	}

	newDroplet, _, err := client.Droplets.Create(createRequest)
	if err != nil {
		return err
	}

	d.DropletID = newDroplet.Droplet.ID

	for {
		newDroplet, _, err = client.Droplets.Get(d.DropletID)
		if err != nil {
			return err
		}
		for _, network := range newDroplet.Droplet.Networks.V4 {
			if network.Type == "public" {
				d.IPAddress = network.IPAddress
			}
		}

		if d.IPAddress != "" {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Debugf("Created droplet ID %d, IP address %s",
		newDroplet.Droplet.ID,
		d.IPAddress)

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
