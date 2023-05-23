## Basic Data Types

### 1. Numbers
1. int8 - 8 bit signed integer
2. int16 - 16 bit signed integer
3. int32 - 32 bit signed integer
4. int64 - 64 bit signed integer
5. uint8 - 8 bit unsigned integer
6. uint16 - 16 bit unsigned integer
7. unit32 - 32 bit unsigned integer
8. uint64 - 64 bit unsigned integer
9. int - Both int and uint contain same size, either 32 or 64 bit.
10. uint - Both int and uint contain same size, either 32 or 64 bit.
11. rune - It is a synonym of int32 and also represent Unicode code points.
12. byte - synonym of uint8
13. uintptr - It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
14. float32 - 32-bit IEEE 754 floating-point number
15. float64 - 64-bit IEEE 754 floating-point number
16. complex64 
17. complex128

### 2. Boolean

### 3. Strings

### 4. [[Arrays]]

### 5. [[Slices]]

### 6. [[Maps]]

### 7. [[Structs]]

### 8. [[Pointers]]

## Declaration
### Method 1
```
var <VariableName> <DataType>;
```
### Method 2(using walrus operator)
```
<VariableName> := <Data>
```
### Method 3(implicit type)
```
var <VariableName> = <Data>
```