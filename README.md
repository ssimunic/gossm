# gossm
Server status monitor written in Go

## Overview

gossm performs checks if servers can be reached every *t* seconds and notifies when unsuccessful

### Web Interface

![Dashboard HTTP server](https://i.imgur.com/IkdmQ4k.png)

## Instruction

### Build and run

Run from terminal:

```bash
go get github.com/ssimunic/gossm/cmd/gossm
go build github.com/ssimunic/gossm/cmd/gossm
./gossm -config configs/myconfig.json
```

This will build and run program with configuration file `configs/myconfig.json`.

### Command line arguments

Following arguments are available:

`config` configuration file path (default "configs/default.json")

`log` log file path (default "logs/current-date.log")

`http` address for http server (default ":8080")

`logfilter` text to filter log by (both console and file)

`nolog` presence of this argument will disable logging to file only

Example: `./gossm -config configs/myconfig.json -logfilter tcp -http :1337`. Web interface will be available at `localhost:1337`.


You can also use `./gossm -help` for help.

## Docker

An example Dockerfile is located in the project root. Currently there is no offical build on Docker Hub, but that can be addressed if more people show interest.

### docker-compose

An example docker-compose is also found in our project root.

```yaml
version: "2"

services:
  gossm:
    build: ./
    ports:
      - "8067:8080"
    volumes:
      - ./configs:/configs
      - ./logs:/var/log/gossum
```

Getting started with `docker-compose` is as simple as having Docker and Docker Compose installed on your machine, and typing:

```bash
docker-compose build
docker-compose up
```

Please note that the example config found at `configs/default.json` is invalid JSON, so we will need to fix that before bringing up the container.

## Configuration

JSON structure is used for configuration. Example can be found in `configs/default.json`.

```json
{
    "settings": {
        "notifications": {
            "email": [
                {
                    "smtp": "smtp.gmail.com",
                    "port": 587,
                    "username": "silvio.simunic@gmail.com",
                    "password": "...",
                    "from": "silvio.simunic@gmail.com",
                    "to": [
                        "silvio.simunic@gmail.com"
                    ]
                }
            ],
            "telegram": [
                {
                    "botToken": "123456:ABC-DEF1234...",
                    "chatId": "12341234"
                }
            ],
            "pushover": [
                {
                    "userKey": "user_key",
                    "appToken": "app_token"
                }
            ],
            "webhook": [
                {
                    "url": "url",
                    "method": "GET"
                }
            ],
            "sms": [
                {
                    "sms": "todo"
                }
            ]
        },
        "monitor": {
            "checkInterval": 15,
            "timeout": 5,
            "maxConnections": 50,
            "exponentialBackoffSeconds": 5
        }
    },
    "servers": [
        {
            "name":"Local Webserver 1",
            "ipAddress":"192.168.20.168",
            "port": 80,
            "protocol": "tcp",
            "checkInterval": 5,
            "timeout": 5
        },
        {
            "name":"Test server 1",
            "ipAddress":"162.243.10.151",
            "port": 80,
            "protocol": "tcp",
            "checkInterval": 5,
            "timeout": 5
        },
        {
            "name":"Test server 2",
            "ipAddress":"162.243.10.151",
            "port": 8080,
            "protocol": "tcp",
            "checkInterval": 5,
            "timeout": 5
        }
    ]
}
```

### Global

`checkInterval` check interval for each server in seconds

`timeout` check connection timeout in seconds

`maxConnections` maximum concurrent connections 

`exponentialBackoffSeconds`

After each notification, time until next notification is available is increased exponetially. On first unsuccessful server reach, notification will always be sent immediately. If, for example, `exponentialBackoffSeconds` is set to `5`, then next notifications will be available after 5, 25, 125... seconds. On successful server reach after downtime, this will be reset.  

### Servers

`name` server name for identification

`ipAddress` server ip address

`port` server port

`protocol` network protocol (tcp, udp)

`checkInterval`  check interval for each server in seconds (this will override global settings)

`timeout` check connection timeout in seconds (this will override global settings)

### Notifications

There can be multiple email or sms notification settings.

#### Email

`smtp` smtp server address

`port` smtp server port

`username` login email

`password` login password

`from` email that notifications will be sent from

`to` array of recipients 

#### Telegram

`botToken` Telegram Bot token obtained via the BotFather.

`chatId` ChatID of the user to message (Can also be a group id).

#### Pushover

`appToken` your Pushover application's API token

`userKey` the user/group key of your Pushover user

#### Webhook

`url` url to make request to

`method` method to use (`GET` or `POST`)

Server information will be stored in `server` parameter.

#### SMS

TODO

## API

JSON of current status is available at `/json` endpoint.

Example is given below.

```json
{
    "tcp 162.243.10.151:80": [
        {
            "time": "2018-03-06T19:57:33.633712261+01:00",
            "online": true
        }
    ],
    "tcp 176.32.98.166:80": [
        {
            "time": "2018-03-06T19:57:33.650150286+01:00",
            "online": true
        }
    ]
}
```