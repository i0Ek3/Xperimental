# bop

> bit operation.

When I use lua to code something, I found lua haven't `>>=` and `<<=`, so we made this, and call it `bop`. Different with other programming languages, lua have it's own features, but we really need something to make our work much easily, that's all. bop have two methods we called `LME` and `RME` to represent `<<=` and `>>=`, you can use it directly by import bop model.

## Usage

Add code `require("bop")` into your source file, like this:

```lua
-- test.lua

require("bop")

bop.lme(1, 2) -- 4
bop.rme(1, 2) -- 0
```

## YSK

> You should know.

bop is a simple wrapper of bit operation, so you shouldn't expect much more, we just ensure bop can work normally. But we'll find a way to update bop's code to support more features on bit operation for use easily.

