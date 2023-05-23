Slices are used very often in Go. Think of them as dynamic arrays
## Declaration

### Method 1
```
var fruitList = []string{"mango","apple","banana"}
```

### Method 2
```
fruitList = make([]string,4)
```
Note that when you use the second method, you can assign only 4 values based on the index. But you can go further and add more values by using the append method.

```
fruitList = append(fruitList,"kiwi");
```
