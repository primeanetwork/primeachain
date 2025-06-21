# clone-and-build Role

## Purpose

This role is part of the PrimeaChain Ansible pipeline.

It clones the PrimeaChain GitHub repository and compiles the `geth` binary using `make`.

## Tasks

- Clones https://github.com/primeanetwork/primeachain into /opt/primea/primeachain
- Builds the geth binary using `make geth`
- Verifies the compiled version using `geth version`

## Requirements

- Go 1.19+ (handled in Phase 1)
- `make` and build tools (handled in Phase 1)

## Notes

No blockchain configuration or data setup is done here — only code cloning and binary compilation.
