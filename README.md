# Introduction

This is a golang implementation of the Emerald Cloud Lab coding challange. 

## Usage

``` 
go run parser.go -f <input data file> -n <the highest N entries>
```

There is an optional `-d` flag for more verbose consol logging.

ofcourse if you would like to build the binary and relocate it to `/usr/local/bin` or add `$GOPATH/bin` to your `$PATH` you would not have to invoke the go compiler everytime to run the app.

## Example session

```
go run parser.go -f data/score_recs.data -n 7
```

## Testing

A few automated unit test cases are included. The test cases utilize input files with various vaild and invalid input scenarios. The input files are located under the `data` directory.

To run the tests:

```
go test
``` 