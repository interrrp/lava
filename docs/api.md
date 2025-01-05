# Lava API Documentation

## Window Functions

### `lava.window.setFps(fps)`

Sets the target frames per second.

- **Parameters:**
  - `fps` (int): The target frames per second.

### `lava.window.setTitle(title)`

Sets the window title.

- **Parameters:**
  - `title` (string): The new title for the window.

### `lava.window.deltaTime()`

Gets the time elapsed between the last two frames.

- **Returns:**
  - (float): The time elapsed between the last two frames.

## Draw Functions

### `lava.draw.clear(color)`

Clears the screen with the specified color.

- **Parameters:**
  - `color` (table): A table with `r`, `g`, `b`, and `a` components representing the color.

### `lava.draw.text(text, x, y, fontSize, color)`

Draws text on the screen.

- **Parameters:**
  - `text` (string): The text to draw.
  - `x` (int): The x-coordinate of the text.
  - `y` (int): The y-coordinate of the text.
  - `fontSize` (int): The font size of the text.
  - `color` (table): A table with `r`, `g`, `b`, and `a` components representing the color.

### `lava.draw.rect(x, y, width, height, color)`

Draws a rectangle on the screen.

- **Parameters:**
  - `x` (int): The x-coordinate of the rectangle.
  - `y` (int): The y-coordinate of the rectangle.
  - `width` (int): The width of the rectangle.
  - `height` (int): The height of the rectangle.
  - `color` (table): A table with `r`, `g`, `b`, and `a` components representing the color.

## Input Functions

### `lava.input.isKeyPressed(key)`

Checks if a key is pressed.

- **Parameters:**

  - `key` (int): The key code to check.

- **Returns:**
  - (bool): `true` if the key is pressed, `false` otherwise.

### `lava.input.isKeyDown(key)`

Checks if a key is currently being held down.

- **Parameters:**

  - `key` (int): The key code to check.

- **Returns:**
  - (bool): `true` if the key is down, `false` otherwise.

### `lava.input.isKeyReleased(key)`

Checks if a key was released.

- **Parameters:**

  - `key` (int): The key code to check.

- **Returns:**
  - (bool): `true` if the key was released, `false` otherwise.

### `lava.input.isKeyUp(key)`

Checks if a key is currently up.

- **Parameters:**

  - `key` (int): The key code to check.

- **Returns:**
  - (bool): `true` if the key is up, `false` otherwise.
