# serveDir
Simple Http Server for static files

## Motivation
Sometimes we need and easy way to share files between different computers and different operating systems. If the other device is connected to the same network you can run a simple http server and share the URL to the other person.

## Installation
If you have a Go development environment you can type the following command in your command line:

```sh
go install github.com/jempe/servedir@latest
```
## Usage
Open the terminal and change the directory to the directory that you want to share, then type:
```sh
servedir 3000
```
Change **3000** with the port number that you want to use

The following message will appear:

```sh
2018/07/27 10:54:26 Starting web server at http://192.168.100.62:3000
```

Open the url of your web server in the other device

## License
GPL © [Jempe](https://www.jempe.org/)
