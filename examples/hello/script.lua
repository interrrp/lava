black = { r = 0, g = 0, b = 0, a = 255 }
white = { r = 255, g = 255, b = 255, a = 255 }

function load()
    print("Hello, world!")
end

function frame()
    draw.clear(black)
    draw.rect(16, 14, 4, 24, white)
    draw.text("Hello, world", 32, 16, 20, white)
end
