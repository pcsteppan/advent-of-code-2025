Wow only 12 days this year! I appreciate that, since I usually only get through 20 days anyhow... And the holidays are probably better spent not staring at a terminal :)
During this Advent of Code, as usual, I'm trying out some new languages, starting with Zig!

I chose zig because I didn't want to write C (jk!), and I have already done a fair amount of rust for AoC. lately, I've been hearing more about zig because the Roc language changed over to zig for their compiler, and it's been interesting to hear about. Last year I was doing a lot of functional programming languages, so I figured I'd try some "systems"-y languages this year.

**Day 1 - Zig**
- first time using zig
- how do you import a dang file in this language... As it turns out `@embedFile` was the easiest answer ([helpful blog post](https://kristoff.it/blog/advent-of-code-zig/))
- basics: parsing, loops, tokenizers

**Day 2 - Zig**
- string manipulation with `std.fmt.bufPrint`
- performance impact of buffer placement

**Day 3 - Zig**
- saturating subtraction (`-|`) for safe index math

**Day 4 - Zig**
- memory ownership: `@embedFile` data is read-only
- `@constCast` doesn't make memory writable, need `@memcpy`
- struct methods
- optionals (`?`)

**Day 5 - Zig**
- more custom struct work, but this was mostly a logic + off-by-one kind of puzzle
- created a "range array" data structure
- struct format - zig's .ToString() equivalent, helped with debugging

**Day 3 - Go**
- went back and did, what I think was the easiest day so far, in go
- first time using go, getting familiar with basics of language
- much much easier to get started w/, and seems nice to write
- parsing, vars + assignments, make(), for loop (the only loop), pointers and references

**Day 6 - Go**
- this problem comes down to parsing
- as a result, getting familiar with 'strings' package and helpers

**Day 7 - Go**
- go is really easy and quick to write!
- maybe my fastest day so far
- work with maps/sets
