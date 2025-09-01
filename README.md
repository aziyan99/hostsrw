# hostsrw

`hostsrw` is a simple cli app / package that able to read and write windows hosts file.

## Installation

Download the latest release from the release page or clone this repository and build the project using `Make` command. The executeable will be available inside the `build` directory.

To build the hostsrw without using `Make`:
```
go build -mod=readonly -ldflags "-s -w" -o .\build\hostsrw.exe .\cmd\hostsrw\main.go
```

## Usage CLI

```
hostsrw add [entry]        : Add a new entry.
hostsrw rm  [entry]        : Remove an existng entry.
hostsrw exists [entry]     : Determine if entry is exists.
```

## Usage Programmatically
Programmatically/package usage will need elevated admin permission.

```go
allHosts := hostsrw.All()

existsHosts := hostsrw.Exists("foo.local")

hostsrw.Add("foo.local")

hostsrw.Remove("foo.local")
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License

[MIT](https://github.com/aziyan99/hostsrw/blob/main/LICENSE)
