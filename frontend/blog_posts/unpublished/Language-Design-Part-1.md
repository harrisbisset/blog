---
title: "Language Design Part 1"
publishedAt: "2024-11-15"
summary: "Some thoughts on language design"
order: 4
---

# Language Design Part 1

Recently, i've been looking into OCaml and as much as I like some features of the language, it's not without what I think are faults.

Here's a little example of a fizzbuzz implementation I cooked up 🍳

```ocaml
let rec fbuzz n =
  if n > 0
  then (
    match [ n mod 3; n mod 5 ] with
    | [ 0; 0 ] -> print_endline "fizzbuzz"
    | [ _; 0 ] -> print_endline "buzz"
    | [ 0; _ ] -> print_endline "fizz"
    | _ -> print_endline @@ string_of_int n);
  fbuzz (n - 1)
;;

fbuzz (int_of_string Sys.argv.(1))
```

If there's anything you can takeaway from OCaml, let it be its pattern matching. Being able to match against different indexes in a construct at once is great.

What is also cool is how you can apply functions with a prefix value (eg. x mod y).

However, whitespacing is great, brackets can make things clearer, and early returns can simpilfy things a bit.
Not only should types be important, but also functions. We want clarity.

```ocaml
fn fbuzz (n int) {
  if n < 0 { return }

  match [ n mod 3, n mod 5 ] {
    [ 0, 0 ] > print_endline "fizzbuzz"
    [ _, 0 ] > print_endline "buzz"
    [ 0; _ ] > print_endline "fizz"
    _ > print_endline (string_of_int n)
  }

  return fbuzz (n - 1)
}

fbuzz (int_of_string Sys.argv[1])
```
