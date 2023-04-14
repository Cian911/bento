# Blauberg Vento Go Client & Co2 Monitor

### Description

Bento is a Go client to communicate and control the set of [Blauberg Vento A*](https://blaubergventilatoren.de/en/series/vento-expert-a50-1-w) smart fans. In addition, this project has the capability to connect to an InfluxDB source, and automate your fans given queried from it. 

This project is made in conjunction with the [Flink Home](https://github.com/Cian911/flink-home) custom Co2 sensor project by [@Cian911](https://github.com/Cian911). If you're interested in building your own sensor, and setting up a real-time pipeline behind it, please visit the project.


### Installation

@TODO Setup goreleaser
@TODO Setup docker

### Quick Start

### Bento Technical Details

Each fan runs a UDP server which we can communicate with. For further details on how the fans operate and how we communicate with them, please visit our [wiki](https://github.com/Cian911/bento/wiki/Blauberg-Vento-Packet-Structure) details the packet structure and communication parameters.
