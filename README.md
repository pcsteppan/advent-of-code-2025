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

<details>
  <summary>Spoiler - approach</summary>
  
  - single pass over the data, keeping track of reduced state in a map
  - update state as you traverse - no tree or graph structure needed
  - made this one pretty quick to solve
  
</details>

**Day 8 - Go**
- custom `Graph` struct for tracking node groups/regions
- struct methods with pointer receivers for mutation (`*Graph`) vs value receivers for read-only (`Graph`)
- `Vec3` struct for 3D points + `Coupling` struct to pair nodes with their distance

<details>
  <summary>Spoiler - approach</summary>
  
  - nested loops are O(n x (n-1)/2) not O(n^2) - generate all unique pairs once, sort by distance, process in order
  - the "graph" doesn't store edges - just tracks which group each node belongs to via `map[Vec3]int`
  - `Connect(a, b)` handles 4 cases: both nodes already grouped (merge groups), one grouped (add other to that group), neither grouped (create new group)
  - track `NodeCount` and `GroupCount` to know when all nodes are in one connected component (part 2 answer)
  
  _addendum_:
  this is essentially a union-find (disjoint set) structure, but my merge is O(n) since I iterate all nodes to update group IDs. union-find uses pointers, so merge is O(1). for a larger input size, that optimization would be useful   
  
</details>

**Day 9 - Go**
- first time using Go modules + separate package (`daynine/lib`)
- ended up reaching for TDD approach after a lot of trial and error, since edge detection logic was tricky, or at least my implementation was tricky
- enums were useful here for my `EdgeType` (Up/Down/Both/None) and `ShapePosition` (In/Out/OnEdge)
- learned Go uses fat pointers for slices/maps - data passed by reference automatically, explicit `*` pointers less needed than expected. passing by value also indicates read-only intent vs. using pointer on type is kind of signaling `mut`
- learning sets in go are just maps to boolean values basically, but struct{} requires no memory whereas bool does, so the convention is to use `map[T]struct{}` rather than `map[T]bool`

<details>
  <summary>Spoiler - approach</summary>
  
  - compressed 2D representation: store only x-coordinates of edges per row, not full grid
  - binary search (`findInsertionPoint`) to check if a point is inside the shape
  
  _addendum_:
  I think there is another way to solve this one, where you track the corners around the corners you are adding in, based off your direction. you'd end up with a set of points that are 'left of the line' and a set of points that are 'right of the line', and one of those would be the set of points which are inside the shape -- and that could also be a valid way to solve the problem, but I ended up going with edge detection because it felt more familiar from a problem in previous years.
  
</details>
