# go-allowed-list

[![godoc][godoc-badge]][godoc]

> A helper library for creating allowed lists using either a blacklist or
> whitelist strategy.

## Example

```go
import (
  "log"

  "github.com/dominicbarnes/go-allowed-list"
)

func main() {
  blacklist := allowedList.NewBlacklist()
  blacklist.Allowed("a") // true, because an empty list is permissive
  blacklist.Add("a", "b", "c")
  blacklist.Allowed("a") // false, because "a" is present
  blacklist.Allowed("d") // true, because "d" is absent

  whitelist := allowedList.NewWhitelist()
  whitelist.Allowed("a") // true, because an empty list is permissive
  whitelist.Add("a", "b", "c")
  whitelist.Allowed("a") // true, because "a" is present
  whitelist.Allowed("d") // false, because "d" is absent
}
```


[godoc-badge]: https://godoc.org/github.com/dominicbarnes/go-allowed-list?status.svg
[godoc]: https://godoc.org/github.com/dominicbarnes/go-allowed-list
