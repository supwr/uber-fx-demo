# Uber Fx by example

Testing several ways to use dependency injection with [fx](https://github.com/uber-go/fx) by creating a small app that consumes a simple [API](http://randomuser.me/api/) that returns a random user's data and prints it at the console.

There is a folder for each approach of using dependency injection with uber fx:

### v1

The vanilla way of injecting the dependencies, by using `func init()` in the `main.go` file to instantiate the needed object and then passing it to the struct.


To run this version:

```
$ cd v1
$ go build -o v1
$ ./v1
```

### v2

The way that i find the simplest of using uber-fx. Using `fx.Provide`, we can [provide](https://uber-go.github.io/fx/annotate.html#annotating-a-function) the constructor that instantiates the struct we need to inject.

To run this version:

```
$ cd v2
$ go build -o v2
$ ./v2
```

### v3

Still using `fx.Provide`, but this time adding the use of `fx.In` and `param objects`. More details [here](https://uber-go.github.io/fx/parameter-objects.html#using-parameter-objects)

To run this version:

```
$ cd v3
$ go build -o v3
$ ./v3
```

### v4

Mapping a struct to an interface, using `fx.Annotate`, as seen [here](https://uber-go.github.io/fx/annotate.html#casting-structs-to-interfaces)

To run this version:

```
$ cd v4
$ go build -o v4
$ ./v4
```

### v5

Using [result objects](https://uber-go.github.io/fx/result-objects.html#using-result-objects) to define what you intend to inject using `fx.Out`.

To run this version:

```
$ cd v4
$ go build -o v4
$ ./v4
```


Those are just some of the ways i found of using Uber Fx to make my life simple. But i surely will keep finding new ways and intend to keep this repo updated with all the new stuff i stumble upon.



