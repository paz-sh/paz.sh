package main

import ()

var (
	flagTarget   string
	flagHosts    int
	flagList     bool
	cmdProvision = &Command{
		Name:    "provision",
		Summary: "Provision instances of Paz on the target provider",
		Usage:   "[-n=N|--number-of-hosts=N] [-l|--list-providers] <provider>",
		Description: `Provision n number of hosts for paz, this will create the instances,
install all of the required paz subsystems and initalise the Paz HTTP UI.
Paz will default to vagrant as the target provider if this value is ommited.

Provision 3 paz machines on vagrant:
    paz provision vagrant

Provision 1 paz machine on vagrant:
    paz provision --number-of-hosts 1 vagrant

Provision 1 paz machine on digitalocean
    paz provision --number-of-hosts 1 digitalocean
`,
		Run: runProvision,
	}
)

func init() {
	cmdProvision.Flags.IntVar(&flagHosts, "n", 3, "Shorthand for --number-of-hosts")
	cmdProvision.Flags.IntVar(&flagHosts, "number-of-hosts", 3, "Number of hosts to provision")
	cmdProvision.Flags.BoolVar(&flagList, "l", false, "Shorthand for --list")
	cmdProvision.Flags.BoolVar(&flagList, "list", false, "List the availible targets")
}

func runProvision(args []string) (exit int) {
	if len(args) != 1 {
		stderr("A target to provision to must be provided")
		return 1
	}

	return
}
