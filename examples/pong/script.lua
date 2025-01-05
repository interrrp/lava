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

local SPEED = {
    PADDLE = 200,
    BALL = 200
}

local gameState = {
    player1 = { y = 140, score = 0 },
    player2 = { y = 140, score = 0 },
    ball = { x = 320, y = 180, speedX = SPEED.BALL, speedY = SPEED.BALL }
}

local function resetBall()
    gameState.ball.x = 320
    gameState.ball.y = 180
end

local function movePaddles()
    local deltaTime = lava.window.deltaTime()
    if lava.input.isKeyDown(KEYS.W) and gameState.player1.y > 0 then
        gameState.player1.y = gameState.player1.y - SPEED.PADDLE * deltaTime
    end
    if lava.input.isKeyDown(KEYS.S) and gameState.player1.y < 280 then
        gameState.player1.y = gameState.player1.y + SPEED.PADDLE * deltaTime
    end
    if lava.input.isKeyDown(KEYS.UP) and gameState.player2.y > 0 then
        gameState.player2.y = gameState.player2.y - SPEED.PADDLE * deltaTime
    end
    if lava.input.isKeyDown(KEYS.DOWN) and gameState.player2.y < 280 then
        gameState.player2.y = gameState.player2.y + SPEED.PADDLE * deltaTime
    end
end

local function moveBall()
    local deltaTime = lava.window.deltaTime()
    gameState.ball.x = gameState.ball.x + gameState.ball.speedX * deltaTime
    gameState.ball.y = gameState.ball.y + gameState.ball.speedY * deltaTime

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
    lava.draw.clear(COLORS.BACKGROUND)
    lava.draw.rect(20, gameState.player1.y, 10, 80, COLORS.PADDLE)
    lava.draw.rect(610, gameState.player2.y, 10, 80, COLORS.PADDLE)
    lava.draw.rect(gameState.ball.x, gameState.ball.y, 10, 10, COLORS.BALL)
    lava.draw.text(gameState.player1.score .. " - " .. gameState.player2.score, 300, 20, 20, COLORS.TEXT)
end

function lava.load()
    lava.window.setTitle("Pong")
end

function lava.frame()
    movePaddles()
    moveBall()
    checkScore()
    drawGame()
end
