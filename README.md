![screenshot](https://github.com/nomura-lab/dps/blob/images/screenshot.png)

# How to use

Download binaly file and execute it.  
```
./dps
```

When you get blow error,

```
panic: Error response from daemon: client version <x.xx> is too new. Maximum supported API version is <y.yy>
```

you have to use `--apiversion` flag.

```
./dps --apiversion <y.yy>
```



