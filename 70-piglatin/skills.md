# Skills for piglatin

## What You'll Learn

**Previous:** [69-options](../69-options/skills.md) | **Next:** [71-romannumbers](../71-romannumbers/README.md)

**Challenge:** Transform a single word into Pig Latin: vowel-start words get `"ay"` appended; consonant-start words move leading consonants to the end and append `"ay"`. No vowels → print `"No vowels"`.

## Core Concept: String Prefix Transformation

### What Is It?

Pig Latin is a simple string transformation rule applied to a single word:
1. If the word starts with a vowel (`a e i o u`): add `"ay"` to the end.
2. If the word starts with consonant(s): move all leading consonants to the end, then add `"ay"`.
3. If the word has no vowels: print `"No vowels"`.

### How It Works

**Step 1 — Check vowels exist:**

```go
func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
           c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func pigLatin(word string) string {
    // Check if any vowel exists
    hasVowel := false
    for i := 0; i < len(word); i++ {
        if isVowel(word[i]) {
            hasVowel = true
            break
        }
    }
    if !hasVowel {
        return "No vowels"
    }
```

**Step 2 — Find the first vowel position:**

```go
    firstVowel := 0
    for firstVowel < len(word) && !isVowel(word[firstVowel]) {
        firstVowel++
    }
```

**Step 3 — Apply the transformation:**

```go
    if firstVowel == 0 {
        // Starts with vowel: just add "ay"
        return word + "ay"
    }
    // Starts with consonant(s): move prefix to end, add "ay"
    return word[firstVowel:] + word[:firstVowel] + "ay"
}
```

**Main:**

```go
func main() {
    if len(os.Args) != 2 {
        return
    }
    fmt.Println(pigLatin(os.Args[1]))
}
```

**Trace `"pig"`:**
- No leading vowel, has vowel `i` at index 1.
- `firstVowel = 1`
- `word[1:] + word[:1] + "ay"` = `"ig"` + `"p"` + `"ay"` = `"igpay"` ✓

**Trace `"Is"` (capital vowel):**
- `isVowel('I')` = true (if you handle uppercase) → `firstVowel = 0`
- `word + "ay"` = `"Isay"` ✓

**Trace `"crunch"`:**
- Consonants: `c`, `r` → `firstVowel = 2`
- `"unch"` + `"cr"` + `"ay"` = `"unchcray"` ✓

**Trace `"crnch"` (no vowels):**
- `hasVowel = false` → return `"No vowels"` ✓

**Key insight — string slicing does the heavy lifting:**

`word[firstVowel:]` — everything from the first vowel to end

`word[:firstVowel]` — the leading consonants

Together: `word[firstVowel:] + word[:firstVowel] + "ay"` is the full transformation.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not checking for no-vowel case first | Words like `"crnch"` run forever or return wrong result | Check `hasVowel` before finding `firstVowel` |
| Only checking lowercase vowels | `"Is"` → `"Isay"` fails if `'I'` not in vowel set | Include uppercase `A E I O U` in `isVowel` |
| Off-by-one in consonant slice | `word[:firstVowel+1]` includes the first vowel | `word[:firstVowel]` stops before the vowel |

## Solving This Challenge

### Algorithm
1. Validate exactly 1 argument.
2. Check if any vowel (a/e/i/o/u, case-insensitive) exists → if not, print `"No vowels"`.
3. Find `firstVowel` = index of the first vowel.
4. If `firstVowel == 0`: print `word + "ay"`.
5. Else: print `word[firstVowel:] + word[:firstVowel] + "ay"`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [71-romannumbers](../71-romannumbers/README.md)
