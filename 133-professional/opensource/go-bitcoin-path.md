# Go → Bitcoin: Detailed Contribution Guide

You know Go. These are the steps to go from "I know Go" to "I have merged code into a production Bitcoin project."

## Why Start with Go Bitcoin Projects

- You do not need to learn a new language
- LND and btcd are large, active codebases with real users and real issues
- The review process teaches you how production-grade Go is written
- Merged PRs build your on-chain reputation as a Bitcoin developer

## The Go Bitcoin Ecosystem

### btcd — Bitcoin Full Node in Go

**Repository:** https://github.com/btcsuite/btcd

btcd is a full Bitcoin node implementation written entirely in Go. It implements the Bitcoin wire protocol, handles block validation, maintains the UTXO set, and provides an RPC interface.

It is maintained by the btcsuite organization, which also maintains btcwallet and several supporting libraries.

**Why contribute here first:**
- Smaller and more focused than LND
- Good documentation
- Issues are labelled clearly
- Code is idiomatic Go — reading it makes you a better Go developer

**Getting started:**

```bash
git clone https://github.com/btcsuite/btcd
cd btcd
go build ./...
go test ./...
```

If the tests pass, your environment is set up correctly.

**Where to find issues:**

Go to https://github.com/btcsuite/btcd/issues and filter by:
- `good-first-issue`
- `beginner`
- `documentation`

Documentation improvements and test coverage additions are the best first contributions. They are easier to get merged and they force you to read the code.

### LND — Lightning Network Daemon

**Repository:** https://github.com/lightningnetwork/lnd

LND is the most widely deployed Lightning Network node implementation. It is written in Go and used by wallets, exchanges, and routing nodes worldwide.

**Why contribute here:**
- Large codebase with many areas to contribute
- Active community and responsive maintainers
- `docs/` directory is extensive — good entry point
- Real economic activity runs on this software

**Getting started:**

```bash
git clone https://github.com/lightningnetwork/lnd
cd lnd

# Install build dependencies
make install

# Run tests
make check

# Or run a subset of tests
go test ./lnwallet/...
```

Read `docs/INSTALL.md` and `docs/contributing-guidelines.md` before opening any PR.

**Where to find issues:**

https://github.com/lightningnetwork/lnd/issues — filter by `good-first-issue`

**Focus areas for first contributions:**
- Documentation corrections in `docs/`
- Adding test cases to existing test files
- Fixing linting issues flagged in CI
- Small bug fixes with clear reproduction steps

### Supporting Libraries

These libraries are dependencies of btcd and lnd. They are smaller and often have more accessible issues:

| Library | Repository | Focus |
|---------|-----------|-------|
| btcutil | github.com/btcsuite/btcutil | Bitcoin utility functions |
| btcd/btcec | github.com/btcsuite/btcd/btcec | Elliptic curve cryptography |
| neutrino | github.com/lightninglabs/neutrino | Light client / SPV |
| Loop | github.com/lightninglabs/loop | Submarine swaps |

## The Contribution Workflow

This workflow applies to btcd, lnd, and every other open source project.

### Step 1: Fork the repository

On GitHub, click "Fork" on the repository page. This creates your own copy.

### Step 2: Clone your fork

```bash
git clone https://github.com/YOUR_USERNAME/btcd.git
cd btcd

# Add the upstream repository
git remote add upstream https://github.com/btcsuite/btcd.git
```

### Step 3: Keep your fork updated

```bash
git fetch upstream
git checkout master
git merge upstream/master
```

Do this before starting any new work.

### Step 4: Create a branch for your change

```bash
git checkout -b fix-rpc-typo
# or
git checkout -b add-test-coverage-blockheader
```

Use descriptive branch names.

### Step 5: Make your change

Write the code. Write or update tests. Make sure existing tests still pass:

```bash
go test ./...
```

For lnd:
```bash
make check
```

### Step 6: Commit your change

```bash
git add specific-file.go
git commit -m "rpc: fix typo in error message for GetBlockHeader"
```

Bitcoin project commit messages follow the format: `area: short description`. Look at recent commits in the repo to match the style.

### Step 7: Push and open a PR

```bash
git push origin fix-rpc-typo
```

Go to GitHub and open a pull request from your fork's branch to the upstream `master` branch.

**PR description should include:**
- What the change does
- Why it is needed
- How to test it
- Link to the issue it fixes (if applicable)

### Step 8: Respond to review

Maintainers will comment. Address every comment. Push new commits to the same branch — do not open a new PR.

Be patient. Maintainers are volunteers. Reviews can take days or weeks on large changes.

## Reading the Code

The best way to prepare for contributing is to read the code before you write any of it.

**Start with btcd:**
- `wire/` — Bitcoin network protocol messages
- `blockchain/` — Block validation and chain management
- `txscript/` — Bitcoin Script interpreter
- `rpcclient/` — RPC client for interacting with a node

**Start with LND:**
- `lnwallet/` — On-chain wallet functionality
- `htlcswitch/` — Lightning payment routing
- `routing/` — Path finding algorithms
- `docs/` — Architecture documentation

Reading these files teaches you more than any tutorial. You see real production Go: error handling patterns, interface design, test structure, and code organization.

## What to Avoid Early On

- Do not propose major refactors as a first contribution
- Do not open issues asking "where do I start?" without reading the existing issues first
- Do not submit a PR that breaks existing tests
- Do not ignore review comments

Bitcoin projects attract serious contributors. The bar for code quality is high. Start small, be thorough, and build trust over time.

## After Your First Merged PR

Once you have one merged PR:

1. Apply to the **Qala Fellowship** (qala.dev) — they want to see real contributions
2. Apply to **Btrust Builders** (bitcoindevs.xyz) — same
3. Start reading the Chaincode Labs seminars (see `community.md`)
4. Begin learning Rust (see `rust-path.md`)

One merged PR is the credential. It proves you can navigate a real codebase, write code that passes review, and work with a team.
