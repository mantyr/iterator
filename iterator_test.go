package iterator

import (
    "testing"
    "strconv"
)

func Data() (items *Items) {
    items = New()

    items.Add("test1", "value1")
    items.Add("test2", "value2")
    items.Add("test2", "value2")
    items.Add("test3", "value3")
    items.Add("test4", "value4")
    return
}

func TestIter(t *testing.T) {
    items := Data()

    i := 0
    for item := range items.Iter() {
        if item.Index != i {
            t.Errorf("Position error, %q, %q", strconv.Itoa(i), strconv.Itoa(item.Index))
        }
        i++
    }

    /* Default range:
    for key, value := items.Items {
        fmt.Println(key, value)
    }
    */
}

func TestItem(t *testing.T) {
    items := Data()

    val, ok := items.Get("test1")
    if !ok {
        t.Errorf("Error value OK, %q", val)
    }
    if val != "value1" {
        t.Errorf("Error value, %q", val)
    }
}

func TestObj(t *testing.T) {
    items := Data()

    obj := "value2_update"

    items.Add("test2", &obj)

    obj = "value2_update_2"

    obj2, _ := items.Get("test2")
    obj3 := obj2.(*string)

    if *obj2.(*string) != obj {
        t.Errorf("Error, %q", *obj2.(*string))
    }
    if *obj3 != obj {
        t.Errorf("Error, %q", *obj3)
    }

    *obj3 = "test_true"
    if obj != "test_true" {
        t.Errorf("Error, %q", obj)
    }
}