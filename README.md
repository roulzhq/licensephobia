# Licensephobia - Don't be afraid of software licenses anymore!

[Licensephobia](https://licensephobia.com) is a tool that let's you easily search for NPM (and soon pip) packages to view the licenses. You can also upload your package files and see what restrictions apply to your project.

## Build the executable for Docker

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
