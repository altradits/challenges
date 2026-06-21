# 157-bitcoin-oss — Contributing to Bitcoin Open Source in Go

## Objective

Make your first real open-source contribution to a production Go Bitcoin project.

The Go Bitcoin ecosystem is built in Go. You already know the language.
The only remaining step is learning the codebase, finding an issue, and shipping a PR.

## The Target Projects

### btcd — Bitcoin Full Node in Go

Repository: https://github.com/btcsuite/btcd

btcd is a complete Bitcoin node implemented in Go. It handles block validation, the peer-to-peer network, and the UTXO set. Reading btcd code is one of the best ways to see production-grade Go in practice.

Start here. The issues are well-labelled and the codebase is approachable.

### LND — Lightning Network Daemon

Repository: https://github.com/lightningnetwork/lnd

LND is the most widely-deployed Lightning Network implementation in production. It is written in Go. Running a node, reading the code, and fixing a bug in LND is how you become a Lightning developer.

Start here after btcd.

## How to Get Your First Merged PR

1. Clone the repo and run the tests: `go build ./...` && `go test ./...`
2. Read the contribution guide (`CONTRIBUTING.md`)
3. Search issues labelled `good-first-issue`
4. Pick one that matches something you have already learned in this course
5. Ask a clarifying question on the issue before writing code
6. Write the fix, write the test, submit the PR
7. Respond to every review comment promptly

The full step-by-step is in [opensource/go-bitcoin-path.md](opensource/go-bitcoin-path.md).
The community map (where to find developers, conferences, fellowships) is in [opensource/community.md](opensource/community.md).

## Why This Matters

A merged PR in btcd or LND is more credible than any course certificate.
It proves you can read unfamiliar Go code, understand a real system, and ship production-quality changes.

## Fellowships for African Developers

Once you have one merged PR:
- **Qala Fellowship** — https://qala.dev — paid 3-month program for African developers
- **Btrust Builders** — https://bitcoindevs.xyz — mentorship program funded by Jack Dorsey

Details in [opensource/community.md](opensource/community.md).
