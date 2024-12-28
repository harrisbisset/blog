<h1>Error Handling Patterns</h1>
<blockquote>
<p>What makes a good error?</p>
</blockquote>
<p>Popular thought suggests:</p>
<ul>
<li>succint handling</li>
<li>easy to test</li>
<li>extendable</li>
<li>checked exceptions are bad</li>
</ul>
<p>What got me thinking about this was, 'if i could design a language, what would i do differently'; What are the things I dislike when writing/designing code and tests.</p>
<p>And the first thing I thought of is error handling, since I've been using Java due to uni coursework. So firstly, I'd like to say, that with the <em>exception</em> of advocates of Java/C# OOP, people don't like exceptions. And I agree with people.</p>
<p>I've personally gone through three iterations of error handling with Java, firstly just throwing an error. However, it's kinda crazy since multiple things throw the same error, and they want to be handled differently. I also dislike try/catch.</p>
<p>Secondly, I created a custom generic error tuple that could return a new tuple with it's value or error. If .GetErr() == null, then I could .GetVal(), else handle. However, this is very verbose, very crazy, and not really how Java is supposed to be written.</p>
<p>Lastly, I've created custom exceptions for each function which take in objects (used to group data like error messages, and error enum contants for the exact point of failure), and throw them. Slightly less crazy than the two other ones, but requires a bunch of boilerplate, and is verbose.</p>
<p>As a result of this negative experience, I'd like to add:</p>
<ul>
<li>some form of stack trace through multiple functions</li>
<li>should use enums for testing (string comparison is bad, and to detect if that exact error was thrown)</li>
<li>Rust's result type is a good start, but not perfect
s</li>
<li>multiple error return values are good</li>
</ul>
<h2>Implementation</h2>
<p>Here's an example of something I might like.</p>
<pre><code>func parentFunc() {
    risk subFunc():
        err -&gt; err.Log -&gt; err.StackTrace -&gt; err.Exit
        res =&gt; isGood

    println(isGood)
}

func subFunc() (boolean, error) {
    println(&quot;pick a number between 1 and 10&quot;)

    risk scanner.read():
        err -&gt; err.AddStack -&gt; match {
            CLOSED  -&gt; customFunction() -&gt; fail
            default -&gt; fail
        }
        res =&gt; strNum

    risk functionCouldErr(strNum):
        err -&gt; err.AddStack -&gt; fail
        res =&gt; num

    if (num &lt; 5) {
        return false, nil
    }

    return true, nil
}

enum&lt;Errors&gt; functionCouldErrErrors() {
    BAD_INPUT
}

func functionCouldErr(strNum string) (int, error) {
    risk convertToInt(strNum):
        err -&gt; fail
        res =&gt; num

    if (num &lt; 1 || num &gt; 10) {
        return 0, functionCouldErrErrors.BAD_INPUT(&quot;input must be between 1 and 10&quot;)
    }

    return num, nil
}
</code></pre>
<p>More Explained:</p>
<ul>
<li>-&gt; : if true then</li>
<li>=&gt; : assigns value</li>
<li>res =&gt; var : local to function, not risk</li>
<li>res and err : local to risk</li>
<li>fail : return nil, err</li>
<li>risk : denotes whether function returns error</li>
</ul>
<h3>Tests:</h3>
<pre><code>func TestInput() {
    risk functionCouldErr(&quot;b&quot;):
        err =&gt; expected
        res =&gt; noRes

    assert(convertToInt.VALUE_STRING, expected)
    assert(nil, noRes)
}

func TestBadInput() {
    risk functionCouldErr(&quot;0&quot;):
        err =&gt; expected
        res =&gt; noRes

    assert(functionCouldErrErrors.BAD_INPUT, expected)
    assert(nil, noRes)
}
</code></pre>
<p>You might be able to tell that it's heavily inspired by golang.</p>
<p>Here's some reasoning:</p>
<ul>
<li>I want to test against enum constants</li>
<li>I want to declare strings at point of error called, not inside enum, so it's centralised. This may be optional behaviour.</li>
<li>You could assign the error value to another variable and return that as well as the result. To get a truthy behaviour.</li>
<li>match statements on error enums seem very powerful.</li>
<li>fail seems like a good, explict word for when an error causes the function to return early.</li>
<li>risk: + indetation so not too bracket heavy.</li>
</ul>
<p>Anyway, just some thoughts.</p>
