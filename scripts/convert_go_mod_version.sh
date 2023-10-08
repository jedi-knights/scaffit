#!/bin/bash

# Specify the path to your go.mod file
go_mod_file="go.mod"

# Function to convert Go version in go.mod
convert_go_version() {
    local go_mod="$1"
    local converted_go_mod="${go_mod}.tmp"

    # Use awk to modify the Go version line in the go.mod file
    awk '/go [0-9]+\.[0-9]+\.[0-9]+$/ && !converted { sub(/\.[0-9]+$/, "", $2); converted=1 } 1' "$go_mod" > "$converted_go_mod"

    # Rename the temporary file to the original go.mod file
    mv "$converted_go_mod" "$go_mod"
}

# Example usage:
convert_go_version "$go_mod_file"
