# Varnish-NewRelic

This is a plugin for NewRelic that reads from `varnishstat` and reports the data to NewRelic.

[![Build Status](https://travis-ci.org/Nalum/vnr.svg?branch=master)](https://travis-ci.org/Nalum/vnr) [![Coverage Status](https://coveralls.io/repos/Nalum/vnr/badge.svg)](https://coveralls.io/r/Nalum/vnr)

# Config

This plugin reads from a JSON config file of the following format, the values here are the defaults:

```
{
    "key": null,
    "instances": [],
    "interval": 1
}
```

Available options:

* `key` : This is the NewRelic API Key for your account.
* `instances` : This is a list of Varnish instance names. Leave empty if you haven't specified your own name for Varnish instances.
* `interval` : The number of seconds to wait between each stat check. <br> **Note**: NewRelic only accepts data once a minute so this data will be processed and then once a minute the data will be sent.

This config file should be located in `/etc` and called `vnr.json` and be readable by the user `vnr`.

# Running the Plugin

You can install this plugin by downloading `install.sh` file from this repository and executing the file as the root user. This will download the latest release from github and set it up on your system, creating all the required files.

The install script can setup a `systemd` or `upstart` init file for you and have it start up with the your server. During installation, if you have not already created the config file, you will be asked for your NewRelic API Key, a list of Varnish instance names and an interval value in seconds, the config file will then be created for you in `/etc/vnr.json`. If the file already exists the install script will skip that part of the setup. Finally the install script will download the latest version and install it in `/usr/local/bin` as `vnr`.
