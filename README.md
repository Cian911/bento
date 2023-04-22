# Blauberg Vento Go Client & Co2 Monitor

### Description

Bento is a Go client to communicate and control the set of [Blauberg Vento A*](https://blaubergventilatoren.de/en/series/vento-expert-a50-1-w) smart fans. In addition, this project has the capability to connect to an InfluxDB source, and automate your fans given queried from it. 

This project is made in conjunction with the [Flink Home](https://github.com/Cian911/flink-home) custom Co2 sensor project by [@Cian911](https://github.com/Cian911). If you're interested in building your own sensor, and setting up a real-time pipeline behind it, please visit the project.


### Installation

This project is ideally intended to be forked and used as reference in your own fan home automation project, but should you wish to download and test it out by using your own config file (please see the sample for reference) you can download and run a binary by doing the following.

**Manually**

```bash
curl https://github.com/Cian911/bento/releases/download/${VERSION}/${PACKAGE_NAME} -o ${PACKAGE_NAME}
sudo tar -xvf ${PACKAGE_NAME} -C /usr/local/bin/
sudo chmod +x /usr/local/bin/bento
```
**Docker**

```bash
docker pull ghcr.io/cian911/bento:${VERSION}

docker run -d ghcr.io/cian911/bento:${VERSION} -config config.yaml
```

### Quick Start

Please see the [sample config](./sample_config.yml) for reference. Once configured, you can pass in the config file and run the binary like so:

```bash
bento -config config.yml
```

### Bento Technical Details

![Fan Packet Structure](./images/packet-structure.png)

Each fan runs a UDP server which we can communicate with. For further details on how the fans operate and how we communicate with them, please visit our [wiki](https://github.com/Cian911/bento/wiki/Blauberg-Vento-Packet-Structure) details the packet structure and communication parameters.
