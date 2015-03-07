# paz.sh
The Paz command-line interface

### Building

paz must be built with Go 1.3+ on a \*nix machine. Simply run `./build` and then copy the binaries out of bin/ onto each of your machines. The tests can similarly be run by simply invoking `./test`.

If you're on a machine without Go 1.3+ but you have Docker installed, run `./build-docker` to compile the binaries instead.

### Testing

Run `./test` to trigger unit and integration tests. [BATS](https://github.com/sstephenson/bats) runs the integration tests against the compiled binary. It should be installed from that link and placed into PATH.

### License

paz is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
