# simple-covid-grouping

The first interview assessment for Golang Developer position at LMWN

## Getting Started

Run the following command to start serve HTTP API.

```sh
$./covid serve-api
```

or

```sh
$go run main.go serve-api
```

This will require go 1.17 or later installed in your computer.

## Configuration

Feel free to adjust and play with the [configuration.](./config/config.yaml)

You can also use your own configuration file by passing file path flag like this.

```$
$./covid serve-api --config {./path/to/your/config.yaml}
```

Default config file is `./config/config,yaml`
