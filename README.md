# week-2-homework-1

## Weekly Contents

- Using Array, Slice and Map
- f, switch, break, continue, goto
- for, while
- Array Creation
- Range
- Functions, Syntax, Naming
- Function vs Method, Parameter vs Argument, Multiple Return
- Error Fundamentals, NIL, Error Type

## Chapter 4

- Arrays
    - They are homogeneous and fixed size. 
    - Any changes that the function makes to array elements affect only the copy, not original.
    - Other than special cases like SHA256's fixed-size hash, arrays are seldom used as function parameters or results;instead, we use slices.
- Slices
    - Lightweight data structure that gives access to subsequence of the elements of an array, which is known as the slices's underlying array. It has 3 component; a pointer, a length, a capacity.
    - Slices are not comparable, so cannot use == to test. (can use bytes.Equal function)
    - Slice elements are indirect;
        - 1-Making it possible for a slice to contain itself. 
        - 2-Fixed slice value may contain diff. element at diff. times.
    - make() create an unnamed array variable and returns a slice of it.
    - copy() : copies elements from one slice to another of the same type.
- Maps

- Structs
- JSON
- Text and HTML Templates

## Chapter 5

- Function Declarations
- Recursion
- Multiple Return Values
- Errors
- Function Values
- Anonymous Functions
- Variadic Functions
- Deferred Function Calls
- Panic
- Recover

## İnanç Gümüş's Repo [Github Repo](https://github.com/inancgumus/learngo)

- [Exercise Solutions](./learngo_exercises)

## Exercises

- An exercise for arithmetic operations [Exercise1](./exercises/aritmetic)
- An exercise for conversions among basic types [Exercise2](./exercises/conv_basic_types)
- An exercise for conversions among string and basic types using strconv package [Exercise3](./exercises/string_types)
- An exercise for iota [Exercise4](./exercises/itoa)
- An exercise for scope [Exercise5](./exercises/scope)
