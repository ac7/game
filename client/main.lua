
local socket = require "socket"
local msgpack = require "MessagePack"
local tcp

function love.load(arg)
	tcp = socket.tcp()
	if not tcp:connect("localhost", 1030) then
		error("Unable to connect to localhost:1030")
	end

	tcp:settimeout(nil)

	_, errmsg = tcp:send("handshake")
	if not result then
	end

	result, errmsg = tcp:receive("*a")
	if not result then
		error("Could not recieve from socket: '" .. tostring(errmsg) .. "'")
	end
	if result ~= "handshake_part_two" then
		error("Invalid handshake response from server '" .. tostring(result) .. "'")
	end

	print("Recieved valid handshake.")

	tcp:close()
end

