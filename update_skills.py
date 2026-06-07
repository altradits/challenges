#!/usr/bin/env python3
"""
Script to update skills.md files based on README.md content in each directory.
Uses 01-only1/skills.md as a template for structure.
"""

import os
import re
from pathlib import Path

# Base directory
BASE_DIR = Path("/home/stathuita/challenges")

# Reference links for continued learning
GENERAL_REFERENCES = [
    "[Go Documentation](https://go.dev/doc/)",
    "[Go Package Documentation](https://pkg.go.dev/)",
    "[Go Tour](https://tour.golang.org/)",
    "[Effective Go](https://go.dev/doc/effective_go)",
]

PACKAGE_REFERENCES = {
    "fmt": "[fmt package](https://pkg.go.dev/fmt)",
    "strings": "[strings package](https://pkg.go.dev/strings)",
    "unicode": "[unicode package](https://pkg.go.dev/unicode)",
    "strconv": "[strconv package](https://pkg.go.dev/strconv)",
    "sort": "[sort package](https://pkg.go.dev/sort)",
    "math": "[math package](https://pkg.go.dev/math)",
    "os": "[os package](https://pkg.go.dev/os)",
    "io": "[io package](https://pkg.go.dev/io)",
    "bufio": "[bufio package](https://pkg.go.dev/bufio)",
    "regexp": "[regexp package](https://pkg.go.dev/regexp)",
    "time": "[time package](https://pkg.go.dev/time)",
    "json": "[encoding/json package](https://pkg.go.dev/encoding/json)",
    "http": "[net/http package](https://pkg.go.dev/net/http)",
}

def get_dirs():
    """Get all directories that have both README.md and skills.md"""
    dirs = []
    for d in sorted(BASE_DIR.iterdir()):
        if d.is_dir() and (d / "README.md").exists() and (d / "skills.md").exists():
            dirs.append(d)
    return dirs

def read_file(path):
    """Read file content"""
    try:
        return path.read_text(encoding="utf-8")
    except Exception as e:
        print(f"Error reading {path}: {e}")
        return ""

def extract_challenge_info(readme_content):
    """Extract challenge name and description from README"""
    lines = readme_content.strip().split("\n")
    challenge_name = ""
    instructions = []
    
    for line in lines:
        line = line.strip()
        if line.startswith("## ") and not challenge_name:
            challenge_name = line.replace("## ", "").strip()
        elif line.startswith("### Instructions"):
            continue
        elif line.startswith("###") or line.startswith("##"):
            continue
        elif line and not line.startswith("```") and not line.startswith("$") and not line.startswith("|"):
            instructions.append(line)
    
    return challenge_name, "\n".join(instructions[:5])

def get_prev_next_links(dir_name, all_dirs):
    """Get previous and next directory links"""
    dirs = [d.name for d in all_dirs]
    try:
        idx = dirs.index(dir_name)
    except ValueError:
        return "", ""
    
    prev_link = ""
    next_link = ""
    
    if idx > 0:
        prev_name = dirs[idx - 1]
        prev_link = f"**Previous:** [{prev_name}](../{prev_name}/skills.md)"
    
    if idx < len(dirs) - 1:
        next_name = dirs[idx + 1]
        next_link = f"**Next:** [{next_name}](../{next_name}/skills.md) - {next_name.replace('-', ' ').title()}"
    
    return prev_link, next_link

def detect_concepts(readme_content, dir_name):
    """Detect what concepts the README teaches based on keywords"""
    concepts = []
    content_lower = readme_content.lower()
    
    # String operations
    if any(word in content_lower for word in ["string", "str", "text", "word", "char"]):
        concepts.append("String manipulation and processing")
    
    # Functions
    if "func " in readme_content or "function" in content_lower:
        concepts.append("Go function definition and usage")
    
    # Loops
    if any(word in content_lower for word in ["loop", "iterate", "range", "for"]):
        concepts.append("Looping constructs (for, range)")
    
    # Conditionals
    if any(word in content_lower for word in ["if", "condition", "return", "bool"]):
        concepts.append("Conditional logic and boolean returns")
    
    # Slices
    if "slice" in content_lower or "[]" in readme_content:
        concepts.append("Slice manipulation and operations")
    
    # Maps
    if "map" in content_lower:
        concepts.append("Map data structures for lookups")
    
    # Numbers/Math
    if any(word in content_lower for word in ["int", "uint", "number", "math", "prime", "digit"]):
        concepts.append("Numeric operations and type conversion")
    
    # Arrays
    if "array" in content_lower or "[10]" in readme_content:
        concepts.append("Array handling and fixed-size collections")
    
    # Pointers (basic)
    if "pointer" in content_lower or "&" in readme_content or "*" in readme_content:
        concepts.append("Pointer basics and memory addresses")
    
    # Packages
    if "import" in readme_content:
        concepts.append("Package imports and standard library usage")
    
    # Command line args
    if "os.Args" in readme_content or "argument" in content_lower:
        concepts.append("Command-line argument handling")
    
    # Printing/output
    if "print" in content_lower or "fmt." in readme_content:
        concepts.append("Formatted output with fmt package")
    
    # Error handling
    if "error" in content_lower or "invalid" in content_lower:
        concepts.append("Error handling and validation")
    
    # If no specific concepts detected, add generic ones
    if not concepts:
        concepts.append("Basic Go programming concepts")
        concepts.append("Problem-solving and algorithm design")
    
    return concepts[:4]  # Limit to 4 main concepts

def get_relevant_packages(readme_content):
    """Extract relevant Go packages from README"""
    packages = []
    content = readme_content.lower()
    
    pkg_map = {
        "fmt": ["fmt.", "print", "println", "printf", "sprintf"],
        "strings": ["strings.", "split", "join", "replace", "trim", "hasprefix", "hassuffix", "index", "contains"],
        "unicode": ["unicode.", "toupper", "tolower", "isspace", "isdigit", "isalpha"],
        "strconv": ["strconv.", "itoa", "atoi", "parseint", "formatint"],
        "sort": ["sort.", "sort.int", "sort.strings"],
        "math": ["math.", "sqrt", "pow", "abs", "max", "min"],
        "os": ["os.", "args", "stdout", "stdin"],
        "io": ["io.", "writer", "reader"],
        "bufio": ["bufio.", "scanner", "reader", "writer"],
        "regexp": ["regexp.", "match", "find", "replace"],
        "time": ["time.", "now", "format", "parse"],
        "encoding/json": ["json.", "marshal", "unmarshal", "encode", "decode"],
        "net/http": ["http.", "get", "post", "handlefunc", "listenandserve"],
    }
    
    for pkg, keywords in pkg_map.items():
        if any(kw in content for kw in keywords):
            packages.append(pkg)
    
    # Always include fmt if printing is mentioned
    if "print" in content and "fmt" not in packages:
        packages.append("fmt")
    
    return packages[:3]  # Limit to 3 packages

def generate_skills_content(dir_name, readme_content, prev_link, next_link):
    """Generate skills.md content based on README"""
    
    challenge_name, instructions = extract_challenge_info(readme_content)
    concepts = detect_concepts(readme_content, dir_name)
    packages = get_relevant_packages(readme_content)
    
    # Get package references
    pkg_refs = [PACKAGE_REFERENCES.get(p, f"[{p} package](https://pkg.go.dev/{p})") for p in packages]
    all_refs = GENERAL_REFERENCES + pkg_refs
    
    # Build the skills.md content
    content = f"""# Skills for {dir_name}

## What You'll Learn

{prev_link}

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** {challenge_name if challenge_name else dir_name.replace('-', ' ').title()}

## Prerequisite Knowledge

Before starting, you should understand:
- Basic Go syntax and program structure
- How to define and call functions
- String manipulation basics in Go
- Control flow (if/else, loops)

## New Concepts Explained

"""
    
    # Add concept explanations based on detected concepts
    for i, concept in enumerate(concepts, 1):
        content += f"""### {i}. {concept}

"""
        if "String" in concept:
            content += """In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

"""
        elif "function" in concept.lower():
            content += """Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

"""
        elif "Loop" in concept:
            content += """Go has only one looping construct: the `for` loop. It can be used in several ways:

```go
// Traditional for loop
for i := 0; i < 10; i++ { }

// While-style loop
for condition { }

// Range loop (for collections)
for index, value := range collection { }
```

For strings, `for...range` iterates over runes, making it safe for UTF-8.

"""
        elif "Conditional" in concept:
            content += """Go uses `if/else` for conditional branching. The condition doesn't need parentheses:

```go
if condition {
    // do something
} else if otherCondition {
    // do something else
} else {
    // default case
}
```

Boolean operators: `&&` (AND), `||` (OR), `!` (NOT).

"""
        elif "Slice" in concept:
            content += """Slices are dynamic, flexible views into arrays. They're the most common data structure in Go:

```go
// Create a slice
numbers := []int{1, 2, 3, 4, 5}

// Slice an existing slice
subset := numbers[1:4]  // [2, 3, 4]

// Append to a slice
numbers = append(numbers, 6)
```

Slices have length (current elements) and capacity (max elements without reallocation).

"""
        elif "Map" in concept:
            content += """Maps are Go's built-in associative data type (hash tables):

```go
// Create a map
seen := make(map[rune]bool)
seen['a'] = true

// Check if key exists
if seen['b'] {
    // key exists
}

// Delete a key
delete(seen, 'a')
```

Maps are reference types - when you pass them to functions, modifications affect the original.

"""
        elif "Numeric" in concept:
            content += """Go supports various numeric types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`.

Common operations:
- `%` (modulo) for remainders
- `/` for division (integer division truncates)
- Type conversion: `int(x)`, `float64(x)`

"""
        elif "Array" in concept:
            content += """Arrays in Go have fixed size. The size is part of the type:

```go
var arr [10]byte  // Array of 10 bytes
arr[0] = 'h'
arr[1] = 'e'
```

For dynamic sizes, use slices instead: `[]byte`

"""
        elif "Pointer" in concept:
            content += """Pointers hold memory addresses. Use `&` to get address, `*` to dereference:

```go
x := 42
ptr := &x    // ptr points to x
fmt.Println(*ptr)  // Prints 42 (dereferenced)
```

In Go, pointers are rarely needed for basic tasks due to pass-by-value semantics for most types.

"""
        elif "Package" in concept:
            content += """Go programs are organized into packages. The `main` package creates an executable:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

Use `go mod init` to start a new module and manage dependencies.

"""
        elif "Command" in concept:
            content += """Access command-line arguments via `os.Args`:

```go
import "os"

func main() {
    args := os.Args[1:]  // Skip program name
    for _, arg := range args {
        fmt.Println(arg)
    }
}
```

Or use the `flag` package for more complex argument parsing.

"""
        elif "Error" in concept or "validation" in concept.lower():
            content += """Go handles errors explicitly. Functions often return `(value, error)`:

```go
result, err := someFunction()
if err != nil {
    // handle error
    return
}
// use result
```

Always check errors - Go doesn't have exceptions!

"""
        else:
            content += """This concept is fundamental to solving the challenge. Review the README.md for specific requirements and examples.

"""
    
    # Add complete solution example if we can extract function signature
    func_match = re.search(r'func\s+(\w+)\s*\(([^)]*)\)\s*([^{\n]*)\s*\{', readme_content)
    if func_match:
        func_name = func_match.group(1)
        params = func_match.group(2)
        return_type = func_match.group(3).strip()
        
        content += f"""## Complete Solution Example

```go
package main

import "fmt"

{func_name} implementation here

func main() {{
    // Test your function
    fmt.Println({func_name}())
}}
```

### Line-by-Line Explanation

1. `package main` - Declares this as an executable program
2. `import "fmt"` - Imports the formatting package
3. Function implementation follows the expected signature
4. `main()` - Tests the function with sample inputs

"""
    
    content += f"""## How to Run Your Program

1. Save the file as `main.go`
2. Open a terminal in the same directory
3. Run: `go run main.go`
4. Verify the output matches expected results

## Skills You'll Learn

After completing this exercise, you'll be able to:
1. Apply {concepts[0].lower() if concepts else 'new programming concepts'} to solve challenges
2. Build on previous skills without repeating them
3. Progress to the next challenge with confidence

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Text processing and validation
- **Savings Calculator** - Input handling and formatting
- **Investment Tracker** - Data parsing and analysis
- **Net Worth Tracker** - String manipulation for reports

## Step-by-Step Approach

1. **Review previous skills** if needed - click the link above
2. **Understand the challenge** - Read the README carefully
3. **Plan your solution** - Sketch the logic before coding
4. **Implement incrementally** - Build and test piece by piece
5. **Verify edge cases** - Test with empty, single, and boundary inputs

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Skipping prerequisites | You'll miss foundational concepts | Review previous skills.md first |
| Overcomplicating | Simple problems have simple solutions | Start with the simplest approach |
| Off-by-one errors | Index math is tricky | Draw it out on paper |
| Wrong return type | Type mismatch errors | Check function signature |

## Practice Tips

- **Draw it out**: Visualize the problem with diagrams
- **Trace manually**: Follow your code step by step with sample input
- **Test boundaries**: Empty string, single char, max values
- **Review previous**: If stuck, go back to the previous skills.md

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

## Knowledge Check

Before coding, make sure you can answer:
1. What new concept does this exercise introduce?
2. Which previous skill does this build on? (Click the link above)
3. What edge cases should I test?

## Reference Links

For continued learning, explore these resources:

"""
    
    for ref in all_refs:
        content += f"- {ref}\n"
    
    content += f"""
{next_link}

---

## Need Help?

If you're stuck:
1. **Review previous skills** - Click the "Previous" link above
2. **Re-read the README** - It contains all the theory you need
3. **Draw the logic** on paper - trace through examples manually
4. **Test with simple input** first, then edge cases

**Remember:** Every expert was once a beginner. Take your time, understand each concept, and don't skip ahead!
"""
    
    return content

def main():
    dirs = get_dirs()
    print(f"Found {len(dirs)} directories with README.md and skills.md")
    
    for d in dirs:
        dir_name = d.name
        readme_path = d / "README.md"
        skills_path = d / "skills.md"
        
        readme_content = read_file(readme_path)
        if not readme_content:
            print(f"Skipping {dir_name}: empty README")
            continue
        
        prev_link, next_link = get_prev_next_links(dir_name, dirs)
        
        new_content = generate_skills_content(dir_name, readme_content, prev_link, next_link)
        
        # Write the updated skills.md
        try:
            skills_path.write_text(new_content, encoding="utf-8")
            print(f"Updated: {dir_name}/skills.md")
        except Exception as e:
            print(f"Error writing {dir_name}/skills.md: {e}")

if __name__ == "__main__":
    main()
