### Hashmap

Implementation of a Hash map with dynamic types for keys and values in Golang.

Hash collision resolved by separate chaining.

### Download and install

```bash
$ go get github.com/suncat2000/hashmap
```
### Interface

```go
type Key interface{}

// HashMaper interface
type HashMaper interface {
	Set(key Key, value interface{}) error
	Get(key Key) (value interface{}, err error)
	Unset(key Key) error
	Count() int
	Iter() <-chan map[string]interface{}
}

func NewHashMap(blockSize int, fn ...func(blockSize int, key Key) (hashKey uint, bucketIdx uint)) HashMaper {}
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/suncat2000/hashmap"
)

type Key struct {
	key interface{}
}
type Value struct {
	value interface{}
}

func main() {
    // create hashmap
    hashMap := hashmap.NewHashMap(16)

    // set int:int
    err := hashMap.Set(1, 1231)
    // set float32:float32
    err := hashMap.Set(1.123, 1231.12312)
    // set string:int
    err := hashMap.Set("key", 12312)
    // set string:string
    err := hashMap.Set("key", "value")
    // setting array:array
    err := hashMap.Set([1]int{1}, []int{1})
    // setting slice:slice
    err := hashMap.Set([]string{"key"}, []string{"value"})
    // setting map:map
    err := hashMap.Set(map[string]string{"key": "value"}, map[string]string{"key": "value"})
    // set struct:struct
    err := hashMap.Set(Key{"key"}, Value{"value"})
    ...
    // get
    value, err := hashMap.Get("key")
    value, err := hashMap.Get([]string{"key"})
    value, err := hashMap.Get(Key{"key"})
    ...
    // unset
    err := hashMap.Unset("key")
    err := hashMap.Unset([]string{"key"})
    err := hashMap.Unset(Key{"key"})
    ...
    // iterate
    for r := range hashMap.Iter() {
        fmt.Sprintf("%s: %s, ", r["key"], r["value"])
    }
}
```

### Test & Benchmark

```
$ go test hashmap_test.go hashfunc.go hashmap.go
```

```
$ go test hashmap_bench_test.go hashfunc.go hashmap.go -bench=. -benchmem -cpu 1
BenchmarkSet16                 	 1000000	      1024 ns/op	     293 B/op	       6 allocs/op
BenchmarkSet64                 	 1000000	      1033 ns/op	     293 B/op	       6 allocs/op
BenchmarkSet128                	 1000000	      1001 ns/op	     293 B/op	       6 allocs/op
BenchmarkSet1024               	 2000000	      1375 ns/op	     293 B/op	       6 allocs/op
BenchmarkStringSet1024         	 1000000	      1645 ns/op	     309 B/op	       7 allocs/op
BenchmarkSliceSet1024          	 1000000	      2759 ns/op	     405 B/op	      10 allocs/op
BenchmarkMapSet16              	  500000	      4654 ns/op	     925 B/op	      15 allocs/op
BenchmarkStructSet16           	 1000000	      2682 ns/op	     389 B/op	      10 allocs/op
BenchmarkGet16                 	 3000000	       582 ns/op	     127 B/op	       2 allocs/op
BenchmarkGet64                 	 3000000	       594 ns/op	     127 B/op	       2 allocs/op
BenchmarkGet128                	 3000000	       517 ns/op	     127 B/op	       2 allocs/op
BenchmarkGet1024               	 3000000	       546 ns/op	     127 B/op	       2 allocs/op
BenchmarkStringGet16           	 1000000	      1139 ns/op	     132 B/op	       3 allocs/op
BenchmarkSliceGet16            	 1000000	      1950 ns/op	     224 B/op	       8 allocs/op
BenchmarkStructGet16           	 1000000	      2413 ns/op	     208 B/op	       6 allocs/op
BenchmarkMapGet16              	  500000	      4008 ns/op	     744 B/op	      18 allocs/op
BenchmarkUnset16               	 3000000	       528 ns/op	     127 B/op	       2 allocs/op
BenchmarkUnset64               	 3000000	       511 ns/op	     127 B/op	       2 allocs/op
BenchmarkUnset128              	 3000000	       523 ns/op	     127 B/op	       2 allocs/op
BenchmarkUnset1024             	 3000000	       562 ns/op	     127 B/op	       2 allocs/op
BenchmarkIntNativeMapSet       	 5000000	       247 ns/op	      70 B/op	       0 allocs/op
BenchmarkStringNativeMapSet    	 2000000	       581 ns/op	      89 B/op	       2 allocs/op
BenchmarkStringNativeMapGet    	 3000000	       340 ns/op	      45 B/op	       1 allocs/op
BenchmarkStringNativeMapDelete 	 5000000	       230 ns/op	      28 B/op	       1 allocs/op
```