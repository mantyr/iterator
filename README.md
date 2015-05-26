# Golang Iterator for ordered map

1. The order of the elements in the map.

2. Transparent use *interface{} as value.

## Installation

  $ go get github.com/mantyr/iterator

## Examples

```Go
package main

import (
    "github.com/mantyr/iterator"
    "fmt"
    "strconv"
)

type Obj struct {
    param1 int
    param2 string
}

func main() {
    items := iterator.New()

    items.Add("test1", "value1")
    items.Add("test2", "value2")
    items.Add("test2", "value2")
    items.Add("test3", "value3")
    items.Add("test4", "value4")

    fmt.Println("Ordered map:")
    for item := range items.Iter() {
        fmt.Printf("Position %q, key %q, value %q \r\n", strconv.Itoa(item.Index), item.Key, item.Value)
    }

    fmt.Println("Default map:")
    for key, value := range items.Items {
        fmt.Printf("Key %q, value %q \r\n", key, value)
    }

    // for *obj
    obj := new(Obj)
    obj.param1 = 10
    obj.param2 = "test1"

    items = iterator.New()
    items.Add("test5", obj)

    fmt.Println("Ordered map *:")
    for item := range items.Iter() {
        fmt.Printf("Position %q, key %q, value param2 %q \r\n", strconv.Itoa(item.Index), item.Key, item.Value.(*Obj).param2)
    }

    fmt.Println("Default map *:")
    for key, value := range items.Items {
        fmt.Printf("Key %q, value param2 %q \r\n", key, value.(*Obj).param2)
    }
}
/*  out
Ordered map:
Position "0", key "test1", value "value1" 
Position "1", key "test2", value "value2" 
Position "2", key "test3", value "value3" 
Position "3", key "test4", value "value4" 
Default map:
Key "test2", value "value2" 
Key "test3", value "value3" 
Key "test4", value "value4" 
Key "test1", value "value1" 
Ordered map *:
Position "0", key "test5", value param2 "test1" 
Default map *:
Key "test5", value param2 "test1"
*/
```

## More Examples

    $ cat iterator_test.go

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
