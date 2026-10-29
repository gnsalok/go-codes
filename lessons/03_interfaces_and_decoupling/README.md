# Lesson 03: Interfaces, Method Sets, and Decoupling

## Learning goals

- Describe an interface value as a dynamic type and dynamic value.
- Determine the method sets of `T` and `*T` without compiling the program.
- Place interfaces at consumer boundaries and avoid interface pollution.
- Explain when a function value is a simpler dependency than an interface.

## Mental model

Interfaces are satisfied implicitly. The method set of `T` contains methods with
receiver `T`; the method set of `*T` contains methods with receiver `T` or `*T`.
The compiler may take an address for an ordinary call on an addressable value, but
that convenience does not change which interfaces the value implements.

Choose receiver semantics from the meaning of the type. Use pointer receivers when
methods mutate the value, copying is unsafe or expensive, or identity matters. Use
value receivers for small values whose methods naturally operate on copies. Keep
the choice consistent across a type's method set.

## Runnable examples

```bash
go run ./lessons/03_interfaces_and_decoupling/examples/interfaces
go run ./lessons/03_interfaces_and_decoupling/examples/method_value
go run ./lessons/03_interfaces_and_decoupling/examples/receiver_semantics
go run ./lessons/03_interfaces_and_decoupling/examples/embedding
go test ./lessons/03_interfaces_and_decoupling/examples/...
```

The interface-pollution progression compares an unnecessary package-owned
interface with a concrete return type:

```bash
go run ./lessons/03_interfaces_and_decoupling/progression/interface_pollution/problem
go run ./lessons/03_interfaces_and_decoupling/progression/interface_pollution/solution
```

## Intentional compiler failures

- `compile_failures/addressability` calls a pointer-receiver method on a conversion
  result, which is not addressable.
- `compile_failures/method_set` passes a `user` value where the interface requires
  a method declared only on `*user`.

```bash
go run ./lessons/03_interfaces_and_decoupling/compile_failures/addressability
go run ./lessons/03_interfaces_and_decoupling/compile_failures/method_set
```

## Production tradeoffs and common mistakes

- An interface with no concrete need for substitution adds indirection without
  decoupling anything.
- A typed nil pointer stored in an interface makes the interface itself non-nil.
- Returning a package-owned interface can hide useful concrete capabilities and
  makes future API evolution harder.
- Large interfaces force unrelated implementations to depend on unrelated methods.
- A method value with a value receiver captures a copy when the method value is
  created; a pointer receiver captures access to the original value.

## Interview prompts

1. Why can `value.Method()` compile while `value` still fails an interface check?
2. When is a one-method interface better than a function parameter, and when is it
   not?
3. How do you detect and avoid a typed-nil interface bug?
4. What is lost when a constructor returns an interface instead of a concrete type?

## Practice

Refactor a service that owns a large repository interface so each consumer declares
only the behavior it calls. Use a fake implementation in a test and explain why the
interface belongs to the consumer.
