black = { r = 0, g = 0, b = 0, a = 255 }
white = { r = 255, g = 255, b = 255, a = 255 }

function load()
    print("Hello, world!")
end

function frame()
    draw.clear(black)
    draw.text("Hello, world", 16, 16, 20, white)
end
