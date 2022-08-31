
-- This is lua.
-- [[
-- lua hello.lua to run it.
-- really funny.
-- ]]

function hello()
    print('hello lua')
end

io.write('Please enter something...like hello?\n')
local line = io.read()
if line == 'hello' then
    hello()
else
    print('Your Grace!')
end
