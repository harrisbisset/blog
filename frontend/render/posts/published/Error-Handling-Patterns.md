---
title: "Error Handling Patterns"
publishedAt: "2024-02-21"
summary: "A look into error handling patterns for language design"
order: 2
---

# Error Handling Patterns

> What makes a good error?

Popular thought suggests:

- succint handling
- easy to test
- extendable
- checked exceptions are bad

What got me thinking about this was, 'if i could design a language, what would i do differently'; What are the things I dislike when writing/designing code and tests.

And the first thing I thought of is error handling, since I've been using Java due to uni coursework. So firstly, I'd like to say, that with the _exception_ of advocates of Java/C# OOP, people don't like exceptions. And I agree with people.

I've personally gone through three iterations of error handling with Java, firstly just throwing an error. However, it's kinda crazy since multiple things throw the same error, and they want to be handled differently. I also dislike try/catch.

Secondly, I created a custom generic error tuple that could return a new tuple with it's value or error. If .GetErr() == null, then I could .GetVal(), else handle. However, this is very verbose, very crazy, and not really how Java is supposed to be written.

Lastly, I've created custom exceptions for each function which take in objects (used to group data like error messages, and error enum contants for the exact point of failure), and throw them. Slightly less crazy than the two other ones, but requires a bunch of boilerplate, and is verbose.

As a result of this negative experience, I'd like to add:

- some form of stack trace through multiple functions
- should use enums for testing (string comparison is bad, and to detect if that exact error was thrown)
- Rust's result type is a good start, but not perfect
  s
- multiple error return values are good

## Implementation

Here's an example of something I might like.

```
func parentFunc() {
    risk subFunc():
        err -> err.Log -> err.StackTrace -> err.Exit
        res => isGood

    println(isGood)
}

func subFunc() (boolean, error) {
    println("pick a number between 1 and 10")

    risk scanner.read():
        err -> err.AddStack -> match {
            CLOSED  -> customFunction() -> fail
            default -> fail
        }
        res => strNum

    risk functionCouldErr(strNum):
        err -> err.AddStack -> fail
        res => num

    if (num < 5) {
        return false, nil
    }

    return true, nil
}

enum<Errors> functionCouldErrErrors() {
    BAD_INPUT
}

func functionCouldErr(strNum string) (int, error) {
    risk convertToInt(strNum):
        err -> fail
        res => num

    if (num < 1 || num > 10) {
        return 0, functionCouldErrErrors.BAD_INPUT("input must be between 1 and 10")
    }

    return num, nil
}
```

More Explained:

- -> : if true then
- => : assigns value
- res => var : local to function, not risk
- res and err : local to risk
- fail : return nil, err
- risk : denotes whether function returns error

### Tests:

```
func TestInput() {
    risk functionCouldErr("b"):
        err => expected
        res => noRes

    assert(convertToInt.VALUE_STRING, expected)
    assert(nil, noRes)
}

func TestBadInput() {
    risk functionCouldErr("0"):
        err => expected
        res => noRes

    assert(functionCouldErrErrors.BAD_INPUT, expected)
    assert(nil, noRes)
}
```

You might be able to tell that it's heavily inspired by golang.

Here's some reasoning:

- I want to test against enum constants
- I want to declare strings at point of error called, not inside enum, so it's centralised. This may be optional behaviour.
- You could assign the error value to another variable and return that as well as the result. To get a truthy behaviour.
- match statements on error enums seem very powerful.
- fail seems like a good, explict word for when an error causes the function to return early.
- risk: + indetation so not too bracket heavy.

Anyway, just some thoughts.
