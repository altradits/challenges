# 158-bitcoin-oss — Contributing to Bitcoin Open Source in Go

## Task

Make your first real open-source contribution to a production Go Bitcoin
project: **btcd** first, then **LND**. This is a manual challenge — there are
no hidden tests, only a merged PR.

- btcd (Bitcoin full node in Go): https://github.com/btcsuite/btcd
- LND (Lightning Network daemon in Go): https://github.com/lightningnetwork/lnd

The full step-by-step is in
[opensource/go-bitcoin-path.md](opensource/go-bitcoin-path.md); the community
map (developers, conferences, fellowships) is in
[opensource/community.md](opensource/community.md).

## Input / Output

- Input: a `good-first-issue` in btcd or LND that matches something you
  learned in this course.
- Output: a merged upstream PR — more credible than any course certificate.

## Examples

How to get the first merged PR:

1. Clone the repo and run the tests: `go build ./...` && `go test ./...`
2. Read the contribution guide (`CONTRIBUTING.md`)
3. Search issues labelled `good-first-issue`
4. Pick one that matches something you have already learned here
5. Ask a clarifying question on the issue before writing code
6. Write the fix, write the test, submit the PR
7. Respond to every review comment promptly

Once you have one merged PR, look at African-developer fellowships: **Qala**
(https://qala.dev) and **Btrust Builders** (https://bitcoindevs.xyz).

## Constraints

- The PR must be yours: your own code, your own tests, your own words in
  review replies.
- Follow the upstream project’s contribution guide exactly (commit style,
  DCO sign-off if they require it, review etiquette).
- One small, clean PR beats one big, noisy PR.

## Time limit

No clock — upstream review sets the pace. Time-box your side: have a PR open
within **14 days** of starting this challenge.

## Blind rule

This challenge is about working in public, so the blind rule bends: reading
upstream code, docs, and review feedback is the job. The code you submit must
still be written by you, from understanding — no pasted snippets you cannot
explain line by line.

## How to check

```bash
go run ./cmd/drill check 158
```

`drill check` walks a self-attestation checklist (PR link, merged status) and
records an honest pass. Your real proof is the merged PR badge on your
GitHub profile.

## Repeat schedule

No 1/3/7/21-day repeats — instead, keep contributing. Every further merged PR
is a rep. `go run ./cmd/drill due` never lists 158 as overdue.
