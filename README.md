# Agrosensor BLE daemon

## Building

```shell
make build
```

## Install systemd service

Install `sensord` binary

```shell
cp sensord /usr/bin/sensord
```

Copy unit file to `/lib/systemd/system/`

```shell
cp systemd/sensord.service /lib/systemd/system/sensord.service
```

Reload systemd units

```shell
systemctl daemon-reload
```

Enable `sensord.service` and start the service

```shell
systemctl enable sensord.service
systemctl start sensord.service
```

## Install/Upgrade Go

```shell
wget https://go.dev/dl/go1.18.linux-armv6l.tar.gz
sudo tar -C /usr/local -xzf go1.18.linux-armv6l.tar.gz
```
