# fn-plugin-example

Example for adding third party plugins to fn project server.

This repo includes an example called `logspam`.

## Method 1 - for end users / operators

This will build a custom Docker image for you with the extensions you provide in a yaml file built in.

```yaml
extensions:
  - name: github.com/treeder/fn-ext-example/logspam
  - name: github.com/treeder/fn-ext-example/logspam2
```

Build it:

```sh
fn build-server -t imageuser/imagename
```

Now run your new server:

```sh
docker run --rm --name fnserver -it -v /var/run/docker.sock:/var/run/docker.sock -v $PWD/data:/app/data -p 8080:8080 imageuser/imagename
```


## Method 2 - for extension developers

This method is good for extension developers as it's fast builds for dev/testing.

Copy [main.go](https://github.com/fnproject/fn/blob/master/main.go) from `fn` repo, then add plugins. See [main.go](main.go) in this
repo for an example.

Then assuming you have the fn project in your GOPATH or you've vendored it here, it should build:

```sh
go build -o fnserver
./fnserver
```

Then deploy a function and you'll see the special spam output like this:

```txt
INFO[0000] Serving Functions API on address `:8080`
YO! This is an annoying message that will happen every time a function is called.
YO! And this is an annoying message that will happen AFTER every time a function is called.
```
