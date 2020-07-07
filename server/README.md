# StolenCars

This is a stolencars app written in Go

Just run the following

## SETUP

1. Install Golang version >= 1.11

```
apt install golang-go
```

2. Verify version is >= 1.11

```
go version
```

3. Add the following to profile file ( the path where go is installed)

```
vi ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin
```

4. Make a go folder inside your home directory. Then run the following commands:

```
cd go
mkdir bin src
cd src
git clone {repo-url}
export GOPATH={path of go directory that you created}
go get
go install
```

5. We are using dep as a dependeny manager for golang. To install dep run:

```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

Once Installation is completed, add dep path to your environment variables. To do so, add the following to your .bashrc or ./bash_profile

```
export PATH=$PATH:$GOPATH/bin
```

6. Run the following command to install dependencies (in project root directory, i.e. where Gopkg.toml file is present)

```
dep ensure -v
```

And to update dependencies, run -

```
dep ensure -update -v
```

This will generate/update Gopkg.lock and Gopkg.toml files. These files should be committed to version control.

If you are running dep from a shared folder in Virtual machine setup, you may run into lock file creation error. Use the following command then:

```
env DEPNOLOCK=1 dep ensure -v
```

Update the latest configuration files config.go and service.go

7. To build and install the app (every time changes are made to the code inside the server repo):

```
go install && go run main.go
```

8. Run the server using the following command

```
go run main.go
```

This will start the Service on port 8081.

9. Make Sure your MongoDB is installed & running and make a DB named `SCLCDB`
