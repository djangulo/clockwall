# clockwall

GOPL, exercise 8.1.

It has three packages:

- `timezone`, whose only purpose is to validate timezones according to the IANA database (taken usually from `/usr/share/timezone` in *nix systems).
- `clock`, creates a fake clock server.
- `main.go` (`clockwall` itself), a client that reads from multiple servers concurrently.


## Run

Assuming you have `Go` installed, `init.sh` will:

- Build the `clock` package into a binary called `clok`.
- Build the `main.go` into a binary called `clockwall`
- Setup some time servers by calling `clok` multiple times
- Run `clockwall`, reading time for each of the time servers and displaying them in a table.