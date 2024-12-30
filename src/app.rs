use std::{
    borrow::BorrowMut,
    fmt::{self, Debug},
    fs,
    path::{Path, PathBuf},
    rc::Rc,
    sync::Arc,
};

use anyhow::{Context, Result};
use mlua::{Lua, ObjectLike, UserData};
use raylib::{
    color::Color,
    prelude::{RaylibDraw, RaylibDrawHandle},
    RaylibHandle, RaylibThread,
};

#[derive(Debug)]
pub struct App {
    dir: PathBuf,
    lua: Lua,
    sdl: Sdl,
}

impl App {
    pub fn from_dir(dir: &Path) -> App {
        let sdl_context = sdl3::init()?;
        sdl_context.event_pump()?;

        App {
            dir: dir.to_owned(),
            lua: Lua::new(),
            raylib_handle,
            raylib_thread,
        }
    }

    pub fn run(&mut self) -> Result<()> {
        self.run_script()?;

        let globals = self.lua.globals();
        globals.call_function::<()>("load", ())?;

        while !self.raylib_handle.window_should_close() {
            let draw_handle = self.raylib_handle.begin_drawing(&self.raylib_thread);
            let lua_draw_handle = LuaDrawHandle(draw_handle);

            globals.call_function::<()>("frame", lua_draw_handle)?;
        }

        Ok(())
    }

    fn run_script(&mut self) -> Result<()> {
        let script_path = &self.dir.join("script.lua");

        let script = fs::read_to_string(script_path)
            .context(format!("Reading script file at {}", script_path.display()))?;

        self.lua.load(script).exec().context("Running script file")
    }
}

struct LuaDrawHandle<'a>(RaylibDrawHandle<'a>);

impl<'a> UserData for LuaDrawHandle<'a> {
    fn add_methods<M: mlua::UserDataMethods<Self>>(methods: &mut M) {
        methods.add_method_mut(
            "text",
            |_, this, (text, x, y, font_size): (String, i32, i32, i32)| {
                this.0.draw_text(&text, x, y, font_size, Color::WHITE);
                Ok(())
            },
        );
    }
}
