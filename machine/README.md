#Paz Machine (for go)

paz-machine should allow you to setup your paz cluster, in a variety of cloud providers or on your local environment, by taking care of all the different interactions you might have with the cluster.

## Operations

It should allow for the following operations:

* Create clusters and machines
* Destroy clusters and machines
* Managing lifecycle (starting, stopping) for clusters
* Inspect clusters
* Check cluster status
* SSH into a given machine

## Drivers

paz-machine should be able to setup its units (clusters and/or machines) in different environments, be it cloud providers or on local environments.

To support that, paz-machine should implement a variety of drivers that would allow it to interface with said environments.

###The requirements for each driver are:

* Built around a provider that offers CoreOS support
* Consume an HTTP API offered by the provider service
* Setup the machines so that a user can access them through SSH

A driver should also provide a set of commands that go in line with the operations that paz-machine should expose to its users

## Operations

### Create

Creates a new cluster or a machine inside a cluster, setting it up with the necessary configurations for it to start working immediately.

### Destroy

Destroys a machine or a cluster, making sure that every single artifact and/or configuration created in the process of setting it up is properly disposed of.

### Starting

Starts a new machine inside a cluster, making sure that the machine is prepared to be ssh'ing into and ready to run paz units.

### Stopping

Stop a machine inside a cluster, making sure that the given machine was stopped correctly

### Inspect

Check machines that are running in a specific cluster and their status (might also give insight over what units are being executed where)

### Check Status

See what clusters the user has configured and their current status (e.g.: running, stopped, destroyed)

### SSH

Allows for simple ssh'ing into a specific machine
