# SigCgt

Command line tool to get signals that process is listening on, writen by golang,
inside the code, it runs a exec call of `cat /proc/%n/status` and then parse the SigCgt field to human readable texts

## Install 

```shell
go get -u github.com/deoops-net/sigcgt
```


## Usage

```shell
sigcgt -p n

```

`n` is the pid of which process you want to inspect


