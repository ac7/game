
local socket = require "socket"
local udp

function love.load()
	tcp = socket.tcp()
	tcp:connect("localhost", 1030)
	tcp:send("handshake")
	tcp:close()
end

