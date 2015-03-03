# Varnish-NewRelic

This is a plugin for NewRelic that reads from `varnishstat` and reports the data to NewRelic. Optionally it can also read from `varnishncsa`.

[![Build Status](https://travis-ci.org/Nalum/vnr.svg?branch=master)](https://travis-ci.org/Nalum/vnr) [![Coverage Status](https://coveralls.io/repos/Nalum/vnr/badge.svg)](https://coveralls.io/r/Nalum/vnr)

# Config

This plugin reads from a JSON config file of the following format:

```
{
    "key": null,
    "instances": [],
    "interval": 1
}
```

Available options:

* `key` : This is the NewRelic API Key for your account. Default: `null`
* `instances` : This is a list of varnish instance names. Leave empty if you haven't specified you own name for varnish instances. Default: `[]`
* `delay` : The number of seconds to wait between each stat check. Default: `1`

# Running the Plugin

You can install this plugin by downloading `install.sh` file from this repository and executing the file as the root user. This will download the latest release from github and set it up on your system, creating all the required files.

The install script will setup a `systemd` or `upstart` init file for you, if you pass the option `--init` and the value of `systemd` or `upstart`, and have it start up with the your server. During installation you will be asked for your NewRelic API Key and the config file will be created for you in `/etc/vnr.json`. You can make any changes you need to here. If the file already exists the install script will skip that part of the setup. Finally the install script will download the latest version and install it in `/usr/local/bin` as `vnr`.
