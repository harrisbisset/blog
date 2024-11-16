<h1>Language Design Part 1</h1>
<p>Recently, i've been looking into OCaml and as much as I like some features of the language, it's not without what I think are faults.</p>
<h2>Good, The Bad &amp; The Ugly</h2>
<h3>The Good</h3>
<p>Here's a pretty nice example of a fizzbuzz implementation I cooked up 🍳</p>
<pre><code>let rec fbuzz n =
  if n &gt; 0
  then (
    match [ n mod 3; n mod 5 ] with
    | [ 0; 0 ] -&gt; print_endline &quot;fizzbuzz&quot;
    | [ _; 0 ] -&gt; print_endline &quot;buzz&quot;
    | [ 0; _ ] -&gt; print_endline &quot;fizz&quot;
    | _ -&gt; print_endline @@ string_of_int n);
  fbuzz (n - 1)
;;

fbuzz (int_of_string Sys.argv.(1))
</code></pre>
<p>You can feel how powerful the pattern matching is by how <em>it condenses knowledge about the program, while maintaing readability</em>!</p>
<h3>The Bad</h3>
<p>An example snippet from a <a href="https://medium.com/@aryangodara_19887/udp-client-and-server-in-ocaml-e203116a997c">udp server implementation</a>:</p>
<pre><code>let rec handle_request (server_socket : Lwt_unix.file_descr Lwt.t) : unit Lwt.t =
  let buffer = Bytes.create 1024 in
  server_socket
  &gt;&gt;= fun server_socket -&gt;
  Lwt_unix.recvfrom server_socket buffer 0 1024 []
  &gt;&gt;= fun (num_bytes, client_address) -&gt;
  let message = Bytes.sub buffer 0 num_bytes in
  let reply = handle_order message in
  match reply with
  | &quot;quit&quot; -&gt;
    Lwt_unix.sendto
      server_socket
      (Bytes.of_string reply)
      0
      (String.length reply)
      []
      client_address
    &gt;&gt;= fun _ -&gt; Lwt.return ()
  | _ -&gt; ...
</code></pre>
<p>While this doesn't have any syntax highlighting, I hope you can agree that it's difficult to read unless you have a lot of prior knowledge of OCaml.</p>
<p>Confusing operators (&gt;&gt;=), crazy syntax for passing lots of parameters, and difficult control flow (when you don't know what you're looking at).</p>
<h3>The Ugly</h3>
<p>As you've probably already seen, unless you're familiar with OCaml (and like the formatting), it's all rather insane.</p>
<p>Now something can be said about a good formatter leaves nobody happy, but that doesn't mean it's right. Although a lot of the problems, at least I feel, are a result of the syntax.</p>
<blockquote>
<p>The provided OCaml examples do use the 'janestreet' profile, however, you're welcome to test the formatting out with 'default' &amp; 'ocamlformat', which I also find disagreeable.</p>
</blockquote>
