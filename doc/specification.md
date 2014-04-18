
# Specification

The server will exchange packets with the client about the state of the
game. Each packet will be delimeted with a newline byte.

Units can be serialized to a representation like this:

	example: Marine;12;24;0;-12.4;18.3
	layout: NAME; HEALTH; MAX_HEALTH; TAGS; XPOS; YPOS

Tags are used to represent attributes of the objects, like `TAG_BIO`
for biological units and `TAG_MECH` for mechanical objects. The full
tag listing can be found in iunit.go (once the list is finalized, the
list will be replicated here). Tags are stored as bit flags and can be
combined with a bitwise-or operation.

