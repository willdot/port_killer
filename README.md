## Port Killer

For when you're debugging Visual Studio Code and it tells you that the port has already been bound, even though you have no debug session running.

## Usage

```
go get -u github.com/willdot/port_killer
```

```
port_killer
```

It will ask you for the port you wish to kill and then present to you information about the process for that port. Confirm you're happy to kill the process, by entering 'Y'.

