# SigCgt

Command Line Tool To Get Signals Process Is Listening, write by golang,
inside it run a exec of `cat /proc/%n/status`, then parse the SigCgt field to human readable texts

## Install 
```
go get -u github.com/deoops-net/sigcgt
```


## Usage

```
sigcgt -p n

```

`n` is the pid of which process you want to inspect


