# golog

A simple unified logging framework.

## Usage

```go
import (
    "os"

    "github.com/lyticaa-public/golog"
)

logger := golog.CreateStdout()

logger.Info("Starting HTTP application.", golog.Attributes{"port": os.Getenv("PORT")})
```
