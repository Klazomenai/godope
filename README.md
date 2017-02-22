# godope [![Build Status](https://travis-ci.org/Klazomenai/godope.svg)](https://travis-ci.org/Klazomenai/godope)

## About
Go project for querying DigitalOcean droplet private LANs and upating /etc/hosts (local to where godope has been run from). This is just a small implementation to help bootstrapping an environment in DigitalOcean when no services such as DNS or Consul have been configured to yet.

Heavily based on https://github.com/tam7t/droplan

## Usage
``` sh
DO_KEY=... ./godope
```
