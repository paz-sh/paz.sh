[![Build Status](https://travis-ci.org/paz-sh/paz.sh.svg?branch=master)](https://travis-ci.org/paz-sh/paz.sh)

# paz.sh
The Paz command-line interface

### Building

paz must be built with Go 1.3+ on a \*nix machine. Simply run `./build` and then copy the binaries out of bin/ onto each of your machines. 

If you're on a machine without Go 1.3+ but you have Docker installed, run `./build-docker` to compile the binaries instead.

### Testing

Run `./test` to trigger unit and integration tests. [BATS](https://github.com/sstephenson/bats) runs the integration tests against the compiled binary. It should be installed from that link and placed into PATH, like so:

```
$ git clone https://github.com/sstephenson/bats.git
$ cd bats/
$ ./install.sh /usr/local
Installed Bats to /usr/local/bin/bats
$ which bats
/usr/local/bin/bats
```

You will also need to install "vet":
```
$ go get code.google.com/p/go.tools/cmd/vet
```

...which in turn depends on [hg](http://mercurial.selenic.com/wiki/Download); you may need to install that too if you haven't done Go builds on your machine before.

### License

paz is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
