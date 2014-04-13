
function love.conf(t)
	t.version = "0.9.1"
	t.console = true

	t.window.title = "client"
	t.window.width = 960
	t.window.height = 540
	t.window.resizeable = true
	t.window.fullscreentype = "desktop"

	t.modules.audio = false
	t.modules.joystick = false
	t.modules.physics = false
	t.modules.mouse = false
	t.modules.sound = false
end

