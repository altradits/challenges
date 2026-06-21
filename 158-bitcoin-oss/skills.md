# Skills for 158-bitcoin-oss

## What You'll Learn

**Previous:** [157-capstone](../157-capstone/skills.md) | **Next:** End of course

How to navigate a large production Go codebase, find a good first issue, write a patch that passes review, and get it merged into a real Bitcoin project.

## Reading a Large Go Codebase Without Getting Lost

When you open btcd or LND for the first time, start at `cmd/` — that is the entry point.

```bash
git clone https://github.com/btcsuite/btcd
cd btcd
# Build to confirm your environment works
go build ./...
# Run tests
go test ./...
# Find the main entry point
ls cmd/
```

Then trace the main function call stack: `main()` → what it calls → what those call.
You do not need to understand everything. You need to understand one small piece at a time.

## Finding the Right Issue

```
Search on GitHub Issues:
  label:good-first-issue
  label:help-wanted
  label:documentation
```

The best first issue is one where you already understand the concept from this course.
For example: a string parsing issue in btcd is something you can solve now.

Look for issues where a maintainer has already described the expected fix in a comment —
that means less ambiguity and faster review.

## Understanding the Existing Code Before Changing It

Before writing a single line, answer these questions:
- What does this function do?
- What calls it?
- What does it call?
- What tests cover it?

```bash
# Find all callers of a function
grep -r "FunctionName" --include="*.go" .

# Find the test file for a package
ls blockchain/  # look for *_test.go files
```

## Writing a Contribution-Quality Patch

Production Go is not the same as exercise Go. The bar is higher:

**Error handling:** every error is checked, not ignored.
```go
// Wrong
result, _ := someFunc()

// Right
result, err := someFunc()
if err != nil {
    return fmt.Errorf("context: %w", err)
}
```

**Tests:** your change must have a test. If you fix a bug, add a regression test.
```go
func TestMyFix(t *testing.T) {
    t.Parallel()  // btcd uses t.Parallel()
    got := myFunction(input)
    if got != want {
        t.Errorf("myFunction(%v) = %v, want %v", input, got, want)
    }
}
```

**Commit message:** one line summary (50 chars max), blank line, then detail.
```
blockchain: fix off-by-one in block height calculation

The block height was calculated as len(chain) when it should be
len(chain)-1 because the genesis block has height 0, not 1.

Fixes #1234
```

## The Review Process

After you submit a PR:
1. A maintainer will review it — this takes days to weeks, not hours
2. They will ask for changes — this is normal, not a rejection
3. Address every comment in a new commit (do not amend during review)
4. When all comments are resolved, the maintainer will approve and merge

Patience and thoroughness in review response is how you build trust.

## Next Step

Read [opensource/go-bitcoin-path.md](opensource/go-bitcoin-path.md) for the full step-by-step guide, and [opensource/community.md](opensource/community.md) for fellowships, newsletters, and IRC channels.

**Next:** End of course
