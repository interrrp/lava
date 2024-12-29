#![deny(clippy::all)]
#![warn(clippy::pedantic)]

use std::path::PathBuf;

use anyhow::Result;
use clap::Parser;

use crate::app::App;

mod app;

#[derive(Parser, Debug)]
struct Args {
    game_path: PathBuf,
}

fn main() -> Result<()> {
    let args = Args::parse();

    let mut game = App::from_app_dir(&args.game_path)?;
    game.run()?;

    Ok(())
}
