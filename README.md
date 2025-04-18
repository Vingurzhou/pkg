# pkg

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

基于Golang的软件包



## Getting Started <a name = "getting_started"></a>


### Installing


```shell
go get github.com/Vingurzhou/pkg
```
## Usage <a name = "usage"></a>

```go 
import "github.com/Vingurzhou/pkg/db"

db.NewGormDB(mysql.open())
```