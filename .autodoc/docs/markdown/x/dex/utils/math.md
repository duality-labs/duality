[View code on GitHub](https://github.com/duality-labs/duality/utils/math.go)

The code in this file provides utility functions for the duality project, focusing on mathematical operations and conversions. These functions are designed to work with the Cosmos SDK, a framework for building blockchain applications in Golang.

1. `BasePrice()`: Returns the base value for price as a decimal, which is 1.0001. This function can be used to set a default price value in the project.

2. `Abs(x int64)`: Calculates the absolute value of an int64 input and returns it as a uint64.

3. `MaxInt64(a, b int64)` and `MinInt64(a, b int64)`: Return the maximum and minimum values, respectively, between two int64 inputs.

4. `MinDec(a, b sdk.Dec)` and `MaxDec(a, b sdk.Dec)`: Return the minimum and maximum values, respectively, between two sdk.Dec (decimal) inputs.

5. `MinIntArr(vals []sdk.Int)` and `MaxIntArr(vals []sdk.Int)`: Return the minimum and maximum values, respectively, from an array of sdk.Int inputs.

6. `Uint64ToSortableString(i uint64)`: Converts a uint64 input to a string that sorts lexicographically in integer order. This can be useful for sorting large numbers as strings.

7. `SafeUint64(in uint64)` and `MustSafeUint64(in uint64)`: Safely convert a uint64 input to an int64 output. The first function returns the converted value and a boolean indicating if an overflow occurred. The second function returns the converted value and panics if an overflow occurs.

These utility functions can be used throughout the duality project to perform common mathematical operations and conversions, ensuring consistency and reducing the need for repetitive code. For example, when comparing two prices in the project, one can use the `MaxDec` and `MinDec` functions to easily determine the higher and lower prices.
## Questions: 
 1. **Question:** What is the purpose of the `BasePrice` function and what does it return?
   **Answer:** The `BasePrice` function returns the base value for price, which is 1.0001. It is used to provide a constant value for price calculations in the duality project.

2. **Question:** How does the `Uint64ToSortableString` function work and what is its use case?
   **Answer:** The `Uint64ToSortableString` function converts a uint64 integer to a string that sorts lexicographically in integer order. This can be useful when you need to store or compare uint64 values as strings while maintaining their numerical order.

3. **Question:** What is the purpose of the `SafeUint64` and `MustSafeUint64` functions, and how do they handle overflow situations?
   **Answer:** The `SafeUint64` function attempts to safely cast a uint64 value to an int64 value, returning the result and a boolean indicating if an overflow occurred. The `MustSafeUint64` function does the same, but instead of returning a boolean, it panics if an overflow occurs. These functions are used to handle situations where casting between uint64 and int64 types is necessary while ensuring that overflow errors are properly handled.