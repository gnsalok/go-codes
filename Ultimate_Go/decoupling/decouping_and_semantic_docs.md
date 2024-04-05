# Decoupling and Language Semantic

### Contents

- [Value and Pointer semantic](#value-and-pointer-semantic`)
- [Methods](#methods)
- [Function](#function)
- [Inteface](#interface)

**Data -> Concrete type -> Decouple -> Behaviour**

> Note: Follow left to right, vise-versa creates problem

## Value and Pointer Semantic

- When to use  value receiver/method?
- When to use ptr receiver/method?

### Answer

- You have to choose the semantic before writing single line of code.
- First of all, This could be very wrong assumption that Mutating methods always have to pointer semantic.
- You can safely mutate and return copy of value semantic.
- **Read** : https://github.com/showa-93/go-mask/blob/13235444922be13823c4664919c53d66b1ad4861/mask.go#L54


### There are 3 classes of Types in Golang :

1. Built-in types : Numeric, String, bool
2. Reference type : Slice, Maps, Channel, Interface value, functions
3. User defined type : Struct

So,

- If you are working with Built-in types, use Value semantic (Including fields (ie. int, string) in struct). (there are some exception case)
- If you are working with Reference types, use Value semantic. (there are some exception case)
- If you have user defined type i.e. **struct**, we have to make a choise between Value and pointer semantic.

#### So, Exception for choosing Pointer semantic :-

- The only use of pointer semantic (in Time standard library) when you about to do Unmarshal and Decode operation. You need to do because it's mutating.
- **Read here** : https://pkg.go.dev/time#Time.UnmarshalBinary

### Important 

- You do not have right  to make COPY, where pointer points to.


### Best Practices
- Value semantic shares COPY, whereas Pointer semantic giving access.
- Don't change the semantics; From value semantic to ptr semantic, that can cause issue.
- Semantic consistency is Everything.



## Methods 

Methods are syntactical sugar which gives us believe system that piece of data has behaviour. 

```go 

type data struct{
	v int
}


func(d data) display(){
	
}

// When you call this piece of code :-
d := data{v: 1}

d.display()

//What Go actually do underneath
data.display(d)    // never ever right code like this.

```

## Function 

- Function in Golang is type.
- You can assign any function to any variable and call later.
- See example :- [func.go](func.go)

## Interface

- Use?? When the piece of data should give the behavious
- Interface types are not real. 
- They only define a method set of behaviour.
- Interface value are valueless.
- When you attach any concrete type it will behave according to the type it passed.


- [IMP] Interface Semantic : https://www.ardanlabs.com/blog/2017/07/interface-semantics.html