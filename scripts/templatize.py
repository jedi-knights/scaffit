#!/usr/bin/env python3

import os
import sys

def walk_directory(directory):
    for root, dirs, files in os.walk(directory):
        for file in files:
            if not file.endswith('.tmpl'):
                base, ext = os.path.splitext(file)
                new_name = f"{base}.tmpl"
                old_path = os.path.join(root, file)
                new_path = os.path.join(root, new_name)
                os.rename(old_path, new_path)
                print(f"Renamed {old_path} to {new_path}")

if __name__ == "__main__":
    if len(sys.argv) > 1:
        target_directory = sys.argv[1]
        walk_directory(target_directory)
    else:
        print("Error: Please provide a target directory as a parameter.")