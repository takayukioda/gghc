Go GitHub Client
=====

This is a GitHub client implemented by Go for personal usage.

Requirements
-----

- Language
    - Go 1.8.3
- Package Management
    - dep
        - revision: 167adc2b67763423775ab0affc7fa5d2deb4fd53

Build the app
-----

```shell
$ dep ensure
$ go build
$ ./gghc
```

Command structure
-----

```shell
$ gghc [general options] <resource> <action> [action options]
```

Options
-----

- General
    - u | user
    - r | repo
    - t | token

Resources & Actions
-----

- labels
    - list
- milestones
    - list
