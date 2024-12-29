use std::{fs, path::Path};

use anyhow::{Context, Result};
use mlua::{Function, Lua, Table};

#[derive(Debug)]
pub struct App {
    lua: Lua,
}

impl App {
    pub fn from_app_dir(app_dir_path: &Path) -> Result<App> {
        let script_path = &app_dir_path.join("script.lua");
        let script = fs::read_to_string(script_path)
            .context(format!("Failed to read {}", script_path.display()))?;

        let mut app = App { lua: Lua::new() };
        app.load_stdlib()?;
        app.lua.load(&script).exec()?;

        Ok(app)
    }

    fn load_stdlib(&mut self) -> Result<()> {
        let stdlib = self.lua.create_table()?;
        self.lua.globals().set("lava", &stdlib)?;

        Ok(())
    }

    pub fn run(&mut self) -> Result<()> {
        let stdlib: Table = self.lua.globals().get("lava")?;

        let load_fn: Function = stdlib.get("load")?;
        let update_fn: Function = stdlib.get("update")?;
        let draw_fn: Function = stdlib.get("draw")?;

        load_fn.call::<()>(())?;
        loop {
            update_fn.call::<()>(())?;
            draw_fn.call::<()>(())?;
        }
    }
}
