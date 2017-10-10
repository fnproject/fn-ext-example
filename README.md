# fn-plugin-example

Example for adding third party plugins to fn project server.

The first is `logspam`.

## Method 1

Copy main.go, add listeners, see main.go here.

Then assuming you have the fn project in your go path of you've vendored it here, it should build:

```sh
go build -o fn-server
./fn-server
```

Then deploy a function and you'll see the special spam output like this:

```
INFO[0000] Serving Functions API on address `:8080`
YO! This is an annoying message that will happen every time a function is called.
YO! And this is an annoying message that will happen AFTER every time a function is called.
```

## Method 2

Docker coolness. First copy main.go and add your extensions, see main.go in this dir.

Then make a Dockerfile with only this in it:

```
FROM fnproject/functions
```

Then run:

```sh
docker build --build-arg REPO=github.com/treeder/fn-plugin-example -t treeder/fn-custom .
```

replace DIR value with the import name for this repo and -t with what you want this image to be called.

Then start it with:

```sh
docker run --rm -it --name functions -v /var/run/docker.sock:/var/run/docker.sock -v $PWD/data:/app/data -p 8080:8080 treeder/fn-custom
```
