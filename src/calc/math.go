package calc

import (
    "fmt"
    "reflect"
)

func Add(a, b interface{})(interface{}, error) {
    value_a := reflect.ValueOf(a)
    value_b := reflect.ValueOf(b)

    //fmt.Println(value_a)
    //fmt.Println(value_b)
    //if value_a.Kind() != value_b.Kind() {
    //    return nil, fmt.Errorf("Different kinds, cannot add them")
    //}
    
    var sum interface{}
    switch value_a.Kind() {
        case  reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            sum = value_a.Int() + value_b.Int()
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            sum = value_a.Uint() + value_b.Uint()
        case reflect.Float32, reflect.Float64:
            sum = value_a.Float() + value_b.Float()
        case reflect.String:
            sum = value_a.String() + value_b.String()
        default:
            return nil, fmt.Errorf("Type does not support addition")
    }

    return sum, nil
}
