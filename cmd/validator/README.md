# OpenPlaybooks/validator

The CACAO validator is a command line tool for validating CACAO cyber security
playbooks(CACAO playbooks) with the Go (Golang) programming language.

## Version 
0.1.1

## Installation

This package can be installed with the go get command:

```
go get github.com/openplaybooks/libcacao/cmd/validator
make
```

## Using the validator

This tool expects CACAO JSON data to be sent in via a pipe "|". For example:

```
cat cacaoplaybook1.json | validator
```


## Help

```
./validator --help
```

## License

This is free software, licensed under the Apache License, Version 2.0.
[Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for
a summary.


## Copyright

Copyright 2021 Bret Jordan, All rights reserved.
