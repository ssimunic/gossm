# gossm
Server status monitor written in Go

## Overview

gossm performs checks if servers can be reached every *t* seconds and notifies when unsuccessful

### Web Interface

![Dashboard HTTP server](https://i.imgur.com/IkdmQ4k.png)

## Instruction

### Build and run

Run from terminal:

`go build github.com/ssimunic/gossm/cmd/gossm && ./gossm -config configs/myconfig.json`

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

#### SMS

TODO
