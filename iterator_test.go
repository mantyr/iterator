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
        fmt.Printl(key, value)
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

func TestItemIndex(t *testing.T) {
    items := Data()

    st := new(interface{})

    *st = "value2_update"
    items.Insert("test2", st)

    val, ok := items.Get("test2")
    if !ok {
        t.Errorf("Error value OK, %q", val)
    }
    if val != "value2_update" {
        t.Errorf("Error Ptr value, %q", val)
    }

    *st = "value2_update_2"

    val, ok = items.Get("test2")
    if !ok {
        t.Errorf("Error value OK, %q", val)
    }
    if val != "value2_update_2" {
        t.Errorf("Error Update items.Items[key] = value, %q", val)
    }
}
func TestNewPointerObject(t *testing.T) {
    items := Data()

    st := new(interface{})
    *st = "value"

    st2 := NewPointer(st)
    items.Insert("test2", st2)

    val, _ := items.Select("test2")
    if *val != "value" {
        t.Errorf("Error value, %q", *val)
    }
}
func TestNewPointerNil(t *testing.T) {
    items := Data()

    st := NewPointer(nil)
    *st = "value"
    items.Insert("test2", st)

    val, _ := items.Select("test2")
    if *val != "value" {
        t.Errorf("Error value, %q", *val)
    }
}
func TestItemIndexGet(t *testing.T) {
    items := Data()

    st := NewPointer(nil)
    *st = "value2_update"

    items.Insert("test2", st)

    val, _ := items.Select("test2")
    *val = "value2_update_2"

    val2, _ := items.Select("test2")

    if val2 != val {
        t.Errorf("Error %q", *val2)
    }
}