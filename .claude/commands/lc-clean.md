# Clean Build Artifacts

Clean any build artifacts and temporary files from the project.

## Usage
```
/lc:clean
```

## Description
This command removes any compiled binaries or temporary files that may have been created during testing or running.

## Implementation
```bash
find src/ -name "*.exe" -delete && find src/ -name "main" -type f -delete && go clean -cache
```

This ensures a clean workspace by removing executables and clearing Go's build cache.