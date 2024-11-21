## Example

<img width="815" alt="readme-screenshot" src="https://github.com/user-attachments/assets/f503e942-bb1a-4ce2-97d5-aa9235e287fe">

## Generate Folder Structure

A Golang-based utility to automatically generate a `README.md` file that documents the structure of your project.

The generated file lists all directories and files in a clean, readable format, making it easier for contributors and users to navigate your project.

### Features

- 📄 **Generates `README.md` with project structure**: Creates a Markdown file with a visual representation of your project's directory tree.
- 🛑 **Ignores hidden files by default**: Automatically skips files and folders starting with a dot (`.`), such as `.gitignore` and `.env`.
- 📝 **Customizable ignore list**: Use the `-ignore` flag to specify additional directories to exclude.

### Installation

```bash
go install github.com/mkafonso/Generate-Folder-Structure@latest
```

Or clone this repository:

```bash
git clone git@github.com:mkafonso/Generate-Folder-Structure.git generate_readme
cd generate_readme
go build -o generate-readme generate_readme.go

# mover para um diretório no PATH
mv generate-readme /usr/local/bin/
```

### Usage

Navigate to the root directory of your project and run:

```bash
generate-readme
```

By default, the tool will generate a README.md file with your project's directory structure, ignoring hidden files.

### Ignoring Specific Directories

To exclude specific directories (e.g., node_modules and vendor), use the -ignore flag:

```bash
generate-readme -ignore="node_modules,vendor"
```

### License

NONE

### Author

mkafonso.dev@gmail.com
