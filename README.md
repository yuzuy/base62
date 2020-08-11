# yuzuy/uuid62
UUID based a unique ID generator without hyphens inspired by [shibukawa/uuid62](https://github.com/shibukawa/uuid62).

## Installation
```
go get -u github.com/yuzuy/uuid62
```

## Usage
```go
package main

import (
    "fmt"

    "github.com/yuzuy/uuid62"
)

func main() {
    fmt.Println(uuid62.MustNew())
}
```

## Differences from [shibukawa/uuid62](https://github.com/shibukawa/uuid62)
- You can decide whether to handle the error by choosing New or MustNew
- This uses [google/uuid](https://github.com/google/uuid) (shibukawa/uuid62 uses [gofrs/uuid](https://github.com/gofrs/uuid))
