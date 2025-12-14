#!/usr/local/bin/lua

local chk=require("checkme")

chk.type_checking(chk)

--chk.show_error("wrong here")

function hello()
    local he = "this is test word: hello"
    print(he)
end

chk.find_in_func(hello)
chk.debugme()
