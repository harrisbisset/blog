---
title: "Language Design Part 1"
publishedAt: "2024-11-15"
summary: "Some thoughts on language design"
order: 4
---

# Language Design Part 1

Recently, i've been looking into OCaml and as much as I like some features of the language, it's not without what I think are faults.

## Good, The Bad & The Ugly

### The Good

Here's a pretty nice example of a fizzbuzz implementation I cooked up 🍳

```
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

You can feel how powerful the pattern matching is by how _it condenses knowledge about the program, while maintaing readability_!

### The Bad

An example snippet from a [udp server implementation](https://medium.com/@aryangodara_19887/udp-client-and-server-in-ocaml-e203116a997c):

```
let rec handle_request (server_socket : Lwt_unix.file_descr Lwt.t) : unit Lwt.t =
  let buffer = Bytes.create 1024 in
  server_socket
  >>= fun server_socket ->
  Lwt_unix.recvfrom server_socket buffer 0 1024 []
  >>= fun (num_bytes, client_address) ->
  let message = Bytes.sub buffer 0 num_bytes in
  let reply = handle_order message in
  match reply with
  | "quit" ->
    Lwt_unix.sendto
      server_socket
      (Bytes.of_string reply)
      0
      (String.length reply)
      []
      client_address
    >>= fun _ -> Lwt.return ()
  | _ -> ...
```

While this doesn't have any syntax highlighting, I hope you can agree that it's difficult to read unless you have a lot of prior knowledge of OCaml.

Confusing operators (\>>=), crazy syntax for passing lots of parameters, and difficult control flow (when you don't know what you're looking at).

### The Ugly

As you've probably already seen, unless you're familiar with OCaml (and like the formatting), it's all rather insane.

Now something can be said about a good formatter leaves nobody happy, but that doesn't mean it's right. Although a lot of the problems, at least I feel, are a result of the syntax.

> The provided OCaml examples do use the 'janestreet' profile, however, you're welcome to test the formatting out with 'default' & 'ocamlformat', which I also find disagreeable.
