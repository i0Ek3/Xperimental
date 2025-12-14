--
--    bop - bit operation
--        - lme, left move equal
--        - rme, right move equal
--
--    In each function, parameter x and n means:
--        - x, the number you want to operate
--        - n, the bit you want to move

bop = {}
bop.constant = "bop is a wrapper for bit operation"

function bop.lme(x, n)
    x = x << n
    return x
end

function bop.rme(x, n)
    x = x >> n
    return x
end

return bop
