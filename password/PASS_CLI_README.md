# Pass-check

```sh
pass-check get
pass-check validate --p "<Strong/weak password>"
```

> Note: `validate` needs a flag password.

To get a strong password run:

```shell
pass-check get
```

To validate the password

````shell
pass-check validate --p "<<strong password>>"
````

For building the modules

```shell
go build .
```

Installing the CLI in the system

```shell
go install
```