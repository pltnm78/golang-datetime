# golang-datetime
*golang-datetime* is a simple library in golang, which handles the concept of date and time easily.


## Purpose
How do you code when you get *time* specified the date and time with raw golang?
```
t, e := time.Parse("2006-01-02 15:04:05 UTC", "2018-08-01 12:30:00 UTC")
```

Is there a need for the first argument?

*golang-datetime* must remove that redundancy.


## Install

Install the package with:

```bash
go get github.com/pltnm78/golang-datetime
```

Import it with:

```go
import "github.com/pltnm78/golang-datetime"
```


## Example
```
package main

import (
    "github.com/pltnm78/golang-datetime"
)

func main() {
    t, e := datetime.GetDatetime("2018-08-01 12:30:00", time.Local)
}
```

## License
MIT License