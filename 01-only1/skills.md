# Skills for 01-only1

## What You'll Learn

**Previous:** This is your first exercise! Start here to build your foundation.

**Challenge:** Write a program that displays a `1` character on the standard output and nothing else.

## New Concept

### 📋 Basic Go Program Structure

```
┌─────────────────────────────────────────────────────────────┐
│                    Go Program Flow                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   ┌───────────┐     ┌───────────┐     ┌───────────┐        │
│   │ package   │────▶│  import   │────▶│  func     │        │
│   │   main    │     │   fmt     │     │   main    │        │
│   └───────────┘     └───────────┘     └───────────┘        │
│         │                 │                 │              │
│         ▼                 ▼                 ▼              │
│   ┌───────────┐     ┌───────────┐     ┌───────────┐        │
│   │ Entry     │     │ Import    │     │ Execution │        │
│   │ Point     │     │ Packages  │     │ Point     │        │
│   └───────────┘     └───────────┘     └───────────┘        │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 🔧 Key Components Summary

| Component | Purpose | Visual |
|-----------|---------|--------|
| `package main` | Declares this as an executable program | 📦 |
| `import "fmt"` | Imports the Format package for I/O | 🔌 |
| `func main()` | Entry point where execution begins | ▶️ |
| `fmt.Println()` | Prints output with newline | 🖨️ |

### 💡 Code Structure Visualization

```
package main          // 🏠 Program declaration
     │
     ▼
import "fmt"          // 🔧 Load formatting tools
     │
     ▼
func main() {         // 🚪 Start here
    fmt.Println("1")  // 📤 Output: 1
}                     // 🔚 End
```

### 🎯 Output Flow

```
Input: (none)
    │
    ▼
┌───────────────┐
│   Program     │
│   Executes    │
└───────────────┘
    │
    ▼
Output: 1
```

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

## Next Steps

**Next:** [02-onlya](../02-onlya/skills.md) - Onlya
