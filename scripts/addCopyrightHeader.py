#!/usr/bin/env python3
# coding: UTF-8
import os
import re

# Get the current year dynamically
CURRENT_YEAR = str(os.popen("date +%Y").read().strip())

# Regular expression to match copyright lines
COPYRIGHT_REGEX = re.compile(r"(// Copyright © )(\d{4})( Ping Identity Corporation)")

def process_go_file(file_path):
    """Reads a Go file, updates or inserts the copyright header if necessary."""
    with open(file_path, "r") as f:
        lines = f.readlines()

    updated = False
    anyMatch = False

    # Check if a copyright line exists and update the year if needed
    for i, line in enumerate(lines):
        match = COPYRIGHT_REGEX.match(line)
        if match:
            anyMatch = True
            old_year = match.group(2)
            if old_year != CURRENT_YEAR:
                lines[i] = f"{match.group(1)}{CURRENT_YEAR}{match.group(3)}\n"
                updated = True
            break  # Stop checking after the first copyright line

    # If no copyright header was found, add it on the first line
    if not anyMatch:
        lines.insert(0, f"// Copyright © {CURRENT_YEAR} Ping Identity Corporation\n\n")
        updated = True

    # Write back only if updates were made
    if updated:
        with open(file_path, "w") as f:
            f.writelines(lines)
        print(f"✅ Fixed copyright in: {file_path}")

def find_and_fix_go_files():
    """Finds all Go files in the repository and processes them."""
    for root, _, files in os.walk("."):
        for file in files:
            if file.endswith(".go"):
                process_go_file(os.path.join(root, file))

    print("✅ All Go files are up to date.")

if __name__ == "__main__":
    find_and_fix_go_files()