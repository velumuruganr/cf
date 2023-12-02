# CF

**cf** is CLI application created using Go. You can build the project and use the command to create a new folder or empty file.


## Usage

    cf flag file/foldername

#### Available flags:
    
    no-flag     : Creates new file if the argument contains dot(.) notation otherwise it will create a new folder
    -f, --file  : Used to create new file
    -d, --folder: Used to create new folder

## Configuration

1. Clone the github repository

        git clone https://github.com/velumuruganr/cf.git

2. Change directory to the project folder

        cd cf

3. Install cobra cli

        go get -u github.com\spf13\cobra@latest

4. Build the go project

        go build main.go

5. Add the binary to your GOPATH

        go install

## Examples

Create a empty file using the cf command

    cf Sample.txt

    cf -f Dockerfile

    cf --file Dockerfile

Create a new folder using the cf command

    cf SampleFolder

    cf -d SampleFolder.com

    cf --folder gihub.com

## Contribution

You are free to make any changes to the project. You are welcome to contribute adding new features.