-- checkme is a module to check somthing for function and others
-- for now, checkme just simple code for use, but we'll keep updating
-- and then, make code stability and beautiful, wait us


checkme = {}

checkme.constant = "checkme is write for error checking"

function checkme.show_error(msg)
    error(msg)
end

type_table = {"nil", "number", "boolean", "string", "table", "function", "userdata", "thread"}

function checkme.type_checking(...)
    for k, v in ipairs{...} do
        if type(v) == type_table[k] then
            print("Give args type is ", type(v))
        else
            print("Wrong type, please check again!")
        end
    end
end

function checkme.find_in_func(func)
    fname = debug.getinfo(func, "n")
    if pcall(func) then
        print("Nice job, no errors in given function: ", fname)
    else
        print("Something wrong in given function: ", fname)
    end
end

function checkme.debugme()
    print(debug.traceback("Staring stack trace..."))
    print(debug.getinfo(1))
    print("Stack trace ending...")
    debug.debug()
end

return checkme
