# [handy](https://godoc.org/github.com/diamondburned/handy)

A Golang binding for [libhandy-1](https://gnome.pages.gitlab.gnome.org/libhandy/).

Some parts of this library are uncovered or may contain bugs (mainly leaks) as
this binding is generated without complete testing (TODO).

## Missing

```
Avatar.GetLoadableIcon
TabBar.GetExtraDragDestTargets
```

## Regenerating

`go generate && go build`
