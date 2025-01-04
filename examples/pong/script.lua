local COLORS = {
    BACKGROUND = {r=0, g=0, b=0, a=255},
    PADDLE = {r=255, g=255, b=255, a=255},
    BALL = {r=255, g=255, b=255, a=255},
    TEXT = {r=255, g=255, b=255, a=255}
}

local KEYS = {
    W = 87,
    S = 83,
    UP = 265,
    DOWN = 264
}

local gameState = {
    player1 = { y = 140, score = 0 },
    player2 = { y = 140, score = 0 },
    ball = { x = 320, y = 180, speedX = 5, speedY = 5 }
}

local function resetBall()
    gameState.ball.x = 320
    gameState.ball.y = 180
end

local function movePaddles()
    if input.isKeyDown(KEYS.W) and gameState.player1.y > 0 then
        gameState.player1.y = gameState.player1.y - 5
    end
    if input.isKeyDown(KEYS.S) and gameState.player1.y < 280 then
        gameState.player1.y = gameState.player1.y + 5
    end
    if input.isKeyDown(KEYS.UP) and gameState.player2.y > 0 then
        gameState.player2.y = gameState.player2.y - 5
    end
    if input.isKeyDown(KEYS.DOWN) and gameState.player2.y < 280 then
        gameState.player2.y = gameState.player2.y + 5
    end
end

local function moveBall()
    gameState.ball.x = gameState.ball.x + gameState.ball.speedX
    gameState.ball.y = gameState.ball.y + gameState.ball.speedY

    if gameState.ball.y <= 0 or gameState.ball.y >= 350 then
        gameState.ball.speedY = -gameState.ball.speedY
    end

    if gameState.ball.x <= 30 and gameState.ball.y > gameState.player1.y and gameState.ball.y < gameState.player1.y + 80 then
        gameState.ball.speedX = -gameState.ball.speedX
    end
    if gameState.ball.x >= 600 and gameState.ball.y > gameState.player2.y and gameState.ball.y < gameState.player2.y + 80 then
        gameState.ball.speedX = -gameState.ball.speedX
    end
end

local function checkScore()
    if gameState.ball.x <= 0 then
        gameState.player2.score = gameState.player2.score + 1
        resetBall()
    end
    if gameState.ball.x >= 640 then
        gameState.player1.score = gameState.player1.score + 1
        resetBall()
    end
end

local function drawGame()
    draw.clear(COLORS.BACKGROUND)
    draw.rect(20, gameState.player1.y, 10, 80, COLORS.PADDLE)
    draw.rect(610, gameState.player2.y, 10, 80, COLORS.PADDLE)
    draw.rect(gameState.ball.x, gameState.ball.y, 10, 10, COLORS.BALL)
    draw.text(gameState.player1.score .. " - " .. gameState.player2.score, 300, 20, 20, COLORS.TEXT)
end

function load()
    draw.setFps(24)
end

function frame()
    movePaddles()
    moveBall()
    checkScore()
    drawGame()
end
