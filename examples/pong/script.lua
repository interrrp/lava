local player1Y = 140
local player2Y = 140
local ballX = 320
local ballY = 180
local ballSpeedX = 5
local ballSpeedY = 5
local score1 = 0
local score2 = 0

function load()
    draw.setFps(24)
end

function frame()
    -- Clear screen
    draw.clear({r=0, g=0, b=0, a=255})

    -- Move paddles
    if input.isKeyDown(87) and player1Y > 0 then -- W key
        player1Y = player1Y - 5
    end
    if input.isKeyDown(83) and player1Y < 280 then -- S key
        player1Y = player1Y + 5
    end
    if input.isKeyDown(265) and player2Y > 0 then -- Up arrow
        player2Y = player2Y - 5
    end
    if input.isKeyDown(264) and player2Y < 280 then -- Down arrow
        player2Y = player2Y + 5
    end

    -- Move ball
    ballX = ballX + ballSpeedX
    ballY = ballY + ballSpeedY

    -- Ball collisions
    if ballY <= 0 or ballY >= 350 then
        ballSpeedY = -ballSpeedY
    end

    -- Paddle collisions
    if ballX <= 30 and ballY > player1Y and ballY < player1Y + 80 then
        ballSpeedX = -ballSpeedX
    end
    if ballX >= 600 and ballY > player2Y and ballY < player2Y + 80 then
        ballSpeedX = -ballSpeedX
    end

    -- Score points
    if ballX <= 0 then
        score2 = score2 + 1
        ballX = 320
        ballY = 180
    end
    if ballX >= 640 then
        score1 = score1 + 1
        ballX = 320
        ballY = 180
    end

    -- Draw everything
    draw.rect(20, player1Y, 10, 80, {r=255, g=255, b=255, a=255})
    draw.rect(610, player2Y, 10, 80, {r=255, g=255, b=255, a=255})
    draw.rect(ballX, ballY, 10, 10, {r=255, g=255, b=255, a=255})
    draw.text(score1 .. " - " .. score2, 300, 20, 20, {r=255, g=255, b=255, a=255})
end
