In Go language, a map is a powerful, ingenious, and versatile data structure. Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.

-   It is a reference to a hash table.
-   Due to its reference type it is inexpensive to pass, for example, for a 64-bit machine it takes 8 bytes and for a 32-bit machine, it takes 4 bytes.
-   In the maps, a key must be unique and always in the type which is comparable using == operator or the type which support _!=_ operator. So, most of the built-in type can be used as a key like an int, float64, rune, string, comparable array and structure, pointer, etc. The data types like slice and noncomparable arrays and structs or the custom data types which are not comparable don’t use as a map key.
-   In maps, the values are not unique like keys and can be of any type like int, float64, rune, string, pointer, reference type, map type, etc.
-   The type of keys and type of values must be of the same type, different types of keys and values in the same maps are not allowed. But the type of key and the type values can differ.
-   The map is also known as a hash map, hash table, unordered map, dictionary, or associative array.
-   In maps, you can only add value when the map is initialized if you try to add value in the uninitialized map, then the compiler will throw an error.

## Declaration

### Method 1
```
// An Empty map
map[Key_Type]Value_Type{}

// Map with key-value pair
map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}
```

```
var mymap map[int]string
```

### Method 2
```
make(map[Key_Type]Value_Type, initial_Capacity)
make(map[Key_Type]Value_Type)
```

## delete()
Use this to delete a key from the map
