[View code on GitHub](https://github.com/duality-labs/duality/osmoutils/generic_helper.go)

The `osmoutils` package contains a function called `MakeNew` that creates a new instance of a generic type `T`. The purpose of this function is to provide a way to create new instances of any type, including pointers to structs. 

The function first creates a variable `v` of type `T`. It then uses reflection to determine if `T` is a pointer or not. If `T` is a pointer, the function creates a new instance of the underlying struct using reflection, and then returns a pointer to it. If `T` is not a pointer, the function allocates memory for a new instance of `T` using the `new` keyword and returns it.

This function can be useful in a variety of scenarios where dynamic creation of new instances of a type is required. For example, it could be used in a factory pattern to create new objects of different types based on some input. 

Here is an example of how `MakeNew` could be used to create a new instance of a struct:

```
type Person struct {
    Name string
    Age int
}

func main() {
    p := MakeNew[Person]()
    p.Name = "John"
    p.Age = 30
    fmt.Println(p) // prints "{John 30}"
}
```

In this example, `MakeNew` is used to create a new instance of the `Person` struct. The returned value is then assigned to the variable `p`, which is a pointer to a new instance of `Person`. The `Name` and `Age` fields of `p` are then set and printed to the console.
## Questions: 
 1. What is the purpose of the `MakeNew` function?
   - The `MakeNew` function is used to create a new instance of a generic type `T`.
2. What happens if `T` is a pointer type?
   - If `T` is a pointer type, the function creates a new instance of the underlying struct using reflection and returns a pointer to it.
3. Why is `reflect` used in the function?
   - `reflect` is used to create a new instance of the underlying struct when `T` is a pointer type, since the type of `T` is not known at compile time.