# Swift Data Structures and Algorithms

## Swift Language Caveats
### Comparing Structures and Classes [1]

Structures and classes in Swift have many things in common. Both can:

    Define properties to store values
    Define methods to provide functionality
    Define subscripts to provide access to their values using subscript syntax
    Define initializers to set up their initial state
    Be extended to expand their functionality beyond a default implementation
    Conform to protocols to provide standard functionality of a certain kind

For more information, see Properties, Methods, Subscripts, Initialization, Extensions, and Protocols.

Classes have additional capabilities that structures don’t have:

    Inheritance enables one class to inherit the characteristics of another.
    Type casting enables you to check and interpret the type of a class instance at runtime.
    Deinitializers enable an instance of a class to free up any resources it has assigned.
    Reference counting allows more than one reference to a class instance.

For more information, see Inheritance, Type Casting, Deinitialization, and Automatic Reference Counting.

The additional capabilities that classes support come at the cost of increased complexity. As a general guideline, prefer structures because they’re easier to reason about, and use classes when they’re appropriate or necessary. In practice, this means most of the custom data types you define will be structures and enumerations. For a more detailed comparison, see Choosing Between Structures and Classes.

## References
1. [Structures and Classes](https://docs.swift.org/swift-book/LanguageGuide/ClassesAndStructures.html)