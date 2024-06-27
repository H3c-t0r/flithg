use tonic_build;
use glob::glob;
use std::path::PathBuf;
use std::fs;

fn main() {
    let mut protos: Vec<PathBuf> = Vec::new();

    // Iterate through directories in protos/flyteidl/
    for entry in fs::read_dir("protos/flyteidl").expect("Failed to read directory") {
        match entry {
            Ok(entry) => {
                if entry.file_type().expect("Failed to get file type").is_dir() {
                    let dir_path = entry.path().join("*.proto");
                    let pattern = dir_path.to_str().expect("Failed to convert path to string");

                    for proto in glob(pattern).expect("Failed to read glob pattern") {
                        match proto {
                            Ok(path) => protos.push(path),
                            Err(e) => eprintln!("Error while reading glob entry: {}", e),
                        }
                    }
                }
            },
            Err(e) => eprintln!("Error while reading directory entry: {}", e),
        }
    }
    
    tonic_build::configure()
        .out_dir("gen/pb_rust")
        .build_server(false)
        .build_client(true)
        .type_attribute(".", "#[::pyo3_macro::with_pyclass]")
        .type_attribute(".", "#[derive(::pyo3_macro::WithNew)]")
        .compile_well_known_types(true)
        .compile(
            &protos.iter().map(|p| p.to_str().unwrap()).collect::<Vec<_>>(),
            &[
                "protos/",
                "protos/google/api/", // Adjust this path to where you have the googleapis protos
                "protos/protoc-gen-openapiv2/options/", // Adjust this path to where you have the grpc-gateway protos
            ],
        )
        .unwrap();
    println!("gRPC Rust client code generation completed successfully.");
}