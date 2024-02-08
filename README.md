# Fiber

[![Doc](https://pkg.go.dev/badge/github.com/go-unity/fiber)](https://pkg.go.dev/github.com/go-unity/fiber)
[![Go](https://img.shields.io/github/go-mod/go-version/go-unity/fiber)](https://go.dev/)
[![Release](https://img.shields.io/github/release/go-unity/fiber.svg)](https://github.com/go-unity/fiber/releases)
[![Test](https://github.com/go-unity/fiber/actions/workflows/test.yml/badge.svg)](https://github.com/go-unity/fiber/actions)
[![Report Card](https://goreportcard.com/badge/github.com/go-unity/fiber)](https://goreportcard.com/report/github.com/go-unity/fiber)
[![Codecov](https://codecov.io/gh/go-unity/fiber/branch/master/graph/badge.svg)](https://codecov.io/gh/go-unity/fiber)
![License](https://img.shields.io/github/license/go-unity/fiber)

Fiber http driver for go-unity.

## Version

| go-unity/fiber | go-unity/framework |
|---------------|---------------------|
| v1.0.x        | v1.0.x              |

## Install

1. Add package

```
go get -u github.com/go-unity/fiber
```

2. Register service provider

```
// config/app.go
import "github.com/go-unity/fiber"

"providers": []foundation.ServiceProvider{
    ...
    &fiber.ServiceProvider{},
}
```

3. Add fiber config to `config/http.go` file

```
// config/http.go
import (
    fiberfacades "github.com/go-unity/fiber/facades"
    "github.com/gofiber/template/html/v2"
    "github.com/gofiber/fiber/v2"
)

"default": "fiber",

"drivers": map[string]any{
    "fiber": map[string]any{
        // prefork mode, see https://docs.gofiber.io/api/fiber/#config
        "prefork": false,
        // Optional, default is 4096 KB
        "body_limit": 4096,
        "route": func() (route.Route, error) {
            return fiberfacades.Route("fiber"), nil
        },
        // Optional, default is "html/template"
        "template": func() (fiber.Views, error) {
            return html.New("./resources/views", ".tmpl"), nil
        },
    },
},
```

## Testing

Run command below to run test:

```
go test ./...
```
