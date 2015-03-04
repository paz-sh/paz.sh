// Copyright 2015 YLD Ltd.
// Copyright 2014 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  "os"
  "flag"
  "fmt"
  "strings"
  "text/tabwriter"

  "github.com/paz-sh/paz.sh/log"
  "github.com/paz-sh/paz.sh/ssh"
)

const (
  cliName        = "paz"
  cliDescription = "paz is a command-line interface to paz-sh, in-house service platform with a PaaS-like workflow."
)

var (
  out           *tabwriter.Writer
  globalFlagset = flag.NewFlagSet("paz", flag.ExitOnError)

  // set of top-level commands
  commands []*Command

  // flags used by all commands
  globalFlags = struct {
    Debug   bool
    Version bool
    Help    bool

    KeyFile  string
    CertFile string
    CAFile   string

    RequestTimeout        float64
    Tunnel                string
    KnownHostsFile        string
    StrictHostKeyChecking bool
    SSHTimeout            float64
    SSHUserName           string
  }{}
)

func init() {
  // call this as early as possible to ensure we always have timestamps
  // on paz logs
  log.EnableTimestamps()

  globalFlagset.BoolVar(&globalFlags.Help, "help", false, "Print usage information and exit")
  globalFlagset.BoolVar(&globalFlags.Help, "h", false, "Print usage information and exit")

  globalFlagset.BoolVar(&globalFlags.Debug, "debug", false, "Print out more debug information to stderr")
  globalFlagset.BoolVar(&globalFlags.Version, "version", false, "Print the version and exit")

  globalFlagset.StringVar(&globalFlags.KeyFile, "key-file", "", "Location of TLS key file used to secure communication with the fleet API or etcd")
  globalFlagset.StringVar(&globalFlags.CertFile, "cert-file", "", "Location of TLS cert file used to secure communication with the fleet API or etcd")
  globalFlagset.StringVar(&globalFlags.CAFile, "ca-file", "", "Location of TLS CA file used to secure communication with the fleet API or etcd")

  globalFlagset.StringVar(&globalFlags.KnownHostsFile, "known-hosts-file", ssh.DefaultKnownHostsFile, "File used to store remote machine fingerprints. Ignored if strict host key checking is disabled.")
  globalFlagset.BoolVar(&globalFlags.StrictHostKeyChecking, "strict-host-key-checking", true, "Verify host keys presented by remote machines before initiating SSH connections.")
  globalFlagset.Float64Var(&globalFlags.SSHTimeout, "ssh-timeout", 10.0, "Amount of time in seconds to allow for SSH connection initialization before failing.")
  globalFlagset.StringVar(&globalFlags.Tunnel, "tunnel", "", "Establish an SSH tunnel through the provided address for communication with fleet and etcd.")
  globalFlagset.Float64Var(&globalFlags.RequestTimeout, "request-timeout", 3.0, "Amount of time in seconds to allow a single request before considering it failed.")
  globalFlagset.StringVar(&globalFlags.SSHUserName, "ssh-username", "core", "Username to use when connecting to CoreOS instance.")
}

type Command struct {
  Name        string       // Name of the Command and the string to use to invoke it
  Summary     string       // One-sentence summary of what the Command does
  Usage       string       // Usage options/arguments
  Description string       // Detailed description of command
  Flags       flag.FlagSet // Set of flags associated with this command

  Run func(args []string) int // Run a command with the given arguments, return exit status

}

func init() {
  out = new(tabwriter.Writer)
  out.Init(os.Stdout, 0, 8, 1, '\t', 0)
  commands = []*Command{
    cmdHelp,
    //cmdSSH,
    cmdVersion,
  }
}

func getAllFlags() (flags []*flag.Flag) {
  return getFlags(globalFlagset)
}

func getFlags(flagset *flag.FlagSet) (flags []*flag.Flag) {
  flags = make([]*flag.Flag, 0)
  flagset.VisitAll(func(f *flag.Flag) {
    flags = append(flags, f)
  })
  return
}

func maybeAddNewline(s string) string {
  if !strings.HasSuffix(s, "\n") {
    s = s + "\n"
  }
  return s
}

func stderr(format string, args ...interface{}) {
  fmt.Fprintf(os.Stderr, maybeAddNewline(format), args...)
}

func stdout(format string, args ...interface{}) {
  fmt.Fprintf(os.Stdout, maybeAddNewline(format), args...)
}

func main() {
  // parse global arguments
  globalFlagset.Parse(os.Args[1:])

  var args = globalFlagset.Args()

  getFlagsFromEnv(cliName, globalFlagset)

  if globalFlags.Debug {
    log.EnableDebug()
  }

  if globalFlags.Version {
    args = []string{"version"}
  } else if len(args) < 1 || globalFlags.Help {
    args = []string{"help"}
  }

  var cmd *Command

  // determine which Command should be run
  for _, c := range commands {
    if c.Name == args[0] {
      cmd = c
      if err := c.Flags.Parse(args[1:]); err != nil {
        stderr("%v", err)
        os.Exit(2)
      }
      break
    }
  }

  if cmd == nil {
    stderr("%v: unknown subcommand: %q", cliName, args[0])
    stderr("Run '%v help' for usage.", cliName)
    os.Exit(2)
  }

  if cmd.Name != "help" && cmd.Name != "version" {
    var err error
    //cAPI, err = getClient()
    if err != nil {
      stderr("Unable to initialize client: %v", err)
      os.Exit(1)
    }
  }

  os.Exit(cmd.Run(cmd.Flags.Args()))

}

// getFlagsFromEnv parses all registered flags in the given flagset,
// and if they are not already set it attempts to set their values from
// environment variables. Environment variables take the name of the flag but
// are UPPERCASE, have the given prefix, and any dashes are replaced by
// underscores - for example: some-flag => PREFIX_SOME_FLAG
func getFlagsFromEnv(prefix string, fs *flag.FlagSet) {
  alreadySet := make(map[string]bool)
  fs.Visit(func(f *flag.Flag) {
    alreadySet[f.Name] = true
  })
  fs.VisitAll(func(f *flag.Flag) {
    if !alreadySet[f.Name] {
      key := strings.ToUpper(prefix + "_" + strings.Replace(f.Name, "-", "_", -1))
      val := os.Getenv(key)
      if val != "" {
        fs.Set(f.Name, val)
      }
    }
  })
}
