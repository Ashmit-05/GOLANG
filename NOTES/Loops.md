Go only has for loop

## Method 1
```
for index:=0;index<len(days);index++{
	if index%2 ==0{
		fmt.Println(days[index])
	}
}
```

## Method 2
```
for i := range days{
	fmt.Println(days[i])
}
```

## Method 3 (take note of goto, break and continue as well)
```
var index2 int64 = 1
for index2<10{
	fmt.Println("Value is: ", index2)
	// break
	if index2 ==5{
		break
	}
	// goto	
	if index2 ==4{
		goto hehe
	}
	// continue
	if index2 ==6{
		continue
	}
	index2++
	}
hehe:
fmt.Println("This is our goto label")
```
