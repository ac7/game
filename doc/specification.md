
# Specification

The server will exchange packets with the client about the state of the
game. Each packet will be delimeted with a newline byte.

Units are serialized with the "msgpack" protocol, a common interface
system like JSON that has implementations in every modern language. The
server expects to send and recieve string-to-value maps. The maps look
like this currently:

	name -> unit name             (string)
	hp -> unit health             (int)
	maxhp -> unit maximum health  (int)
	tags -> unit tags             (int)
	x -> unit x position          (float)
	y -> unit y position          (float)

Eventually the server will support partial updates (for example, just
position or just health updates).

Tags are used to represent attributes of the objects, like `TAG_BIO`
for biological units and `TAG_MECH` for mechanical objects. The full
tag listing can be found in iunit.go (once the list is finalized, the
list will be replicated here). Tags are stored as bit flags and can be
combined with a bitwise-or operation.

