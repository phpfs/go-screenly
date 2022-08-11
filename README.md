# go-screenly

Golang API client that wraps [Screenly API](https://developer.screenlyapp.com) based upon their Swagger definition file.

### Installation

If you use `go mod`, simply import `github.com/phpfs/go-screenly/screenly` and run `go mod tidy`.


For other setups, pull the library via:
```bash
go get -u github.com/phpfs/go-screenly/screenly
```

### Usage

See the `examples` folder for usage details.

### Code generation

In case the `swagger.yml` in this repository should become obsolete, simply clone this repository, replace `swagger.yml` and run `make generate` to regenerate all source code.

## License etc.

This repository is based upon [`go-netbox`](https://github.com/netbox-community/go-netbox) and therefore available under [Apache 2.0 license](LICENSE.md) as well.

The `swagger.yml` does not belong to this repository but is only a copy of the [official api definition](https://developer.screenlyapp.com/?format=openapi) by [Screenly Inc.](https://www.screenly.io/about-us/).
