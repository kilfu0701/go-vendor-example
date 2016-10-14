# go-vendor-example
An example that using "govendor" into your project.

## How to use
  Step 1: Clone repo and setup GOPATH.
```
git clone https://github.com/kilfu0701/go-vendor-example.git
cd go-vendor-example
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
```

  Setp 2: Install `govender`
```
go get -u github.com/kardianos/govendor
```

  Step 3: Download vendor packages
```
cd src/app
govendor sync
```

  Step 4: Run application, and check `localhost:8081`
```
cd $GOPATH
go run src/main.go
```

## Support version
  - Go 1.5 (run with `GO15VENDOREXPERIMENT=1 go run src/main.go`)
  - Go 1.6
  - Go 1.7

## Contributors
kilfu0701 (kilfu0701@gmail.com)

## License
MIT
