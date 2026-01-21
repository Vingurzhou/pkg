# pkg

## Table of Contents

- [pkg](#pkg)
  - [Table of Contents](#table-of-contents)
  - [About ](#about-)
    - [加扰](#加扰)
    - [基于 TLE 的卫星轨道递推方法（SGP4）](#基于-tle-的卫星轨道递推方法sgp4)
  - [Getting Started ](#getting-started-)
    - [Installing](#installing)
  - [Usage ](#usage-)

## About <a name = "about"></a>

基于Golang的软件包

### 加扰

### 基于 TLE 的卫星轨道递推方法（SGP4）

用 SGP4 得到 r / v → 反算瞬时六根数（振荡根数）

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
