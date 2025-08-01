use std::{io::Result, path::Path};

use prost_build::Config;

fn main() -> Result<()> {
    Config::new()
        .out_dir(Path::new(&std::env::var_os("CARGO_MANIFEST_DIR").unwrap()).join("src"))
        .compile_protos(&["../../../proto/raft.proto"], &["../../../proto"])
}
