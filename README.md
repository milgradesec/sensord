# Agrosensor BLE daemon

## Building

```shell
make build
```

## Install systemd service

Install `sensord` binary

```shell
sudo cp sensord /usr/bin/sensord
```

Copy unit file to `/lib/systemd/system/`

```shell
sudo cp systemd/sensord.service /lib/systemd/system/sensord.service
```

Reload systemd units

```shell
sudo systemctl daemon-reload
```

Enable `sensord.service` and start the service

```shell
sudo systemctl enable sensord.service
sudo systemctl start sensord.service
```

Check if running properly
```shell
sudo systemctl status sensord
```

## Install/Upgrade Go

```shell
# Set Go version to install
export GO_VERSION=1.19

# Download & install
wget https://go.dev/dl/go$GO_VERSION.linux-armv6l.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go$GO_VERSION.linux-armv6l.tar.gz

# Add to $PATH
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/home/pi/go/bin
```
