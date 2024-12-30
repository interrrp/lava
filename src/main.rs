#![deny(clippy::all)]
#![warn(clippy::pedantic)]

use std::path::PathBuf;

use anyhow::Result;
use clap::Parser;

use crate::app::App;

mod app;

#[derive(Parser, Debug)]
struct Args {
    app_dir: PathBuf,
}

fn main() -> Result<()> {
    let args = Args::parse();

    let mut app = App::from_dir(&args.app_dir);
    app.run()?;

    Ok(())
}
