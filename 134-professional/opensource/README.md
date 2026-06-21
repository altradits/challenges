# Bitcoin Open Source Contribution Guide

You have completed the Go challenges series. You can write Go. Now the question is: what do you build, and who do you build it with?

This guide maps a concrete path from Go developer to Bitcoin open source contributor — through Go Lightning projects first, then Rust for Bitcoin infrastructure, then TypeScript and Python for applications.

## Why Bitcoin Open Source?

Bitcoin is the most important open monetary network on earth. Its code is maintained entirely by volunteers and contributors. There is no central company deciding what gets merged. Contributing to Bitcoin open source means:

- You are building tools that protect financial sovereignty for people who have no other option
- You are working alongside some of the best systems programmers in the world
- You are building a reputation in a field that values skill above credentials

For African developers specifically: Bitcoin is not theoretical. It is a financial lifeline for countries with unstable currencies and broken banking systems. Your contribution is not a resume item — it is infrastructure.

## The Full Path

```
This repo (Go challenges 01-128)
    |
    v
btcd contributor (Go, Bitcoin full node)
    |
    v
lnd contributor (Go, Lightning Network)
    |
    v
Apply to Qala Fellowship / Btrust Builders
    |
    v
rust-bitcoin (learn Rust, Bitcoin primitives)
    |
    v
BDK contributor (Rust, wallet development)
    |
    v
LDK contributor (Rust, Lightning in Rust)
    |
    v
Bitcoin Core (C++, expert level — optional)
```

## Phase 1: Go — Bitcoin Lightning and Layer 2 (Start Here)

These projects are written in Go. You already know Go. Start here.

| Project | Repository | What It Is |
|---------|-----------|-----------|
| LND | github.com/lightningnetwork/lnd | Full Lightning Network node |
| btcd | github.com/btcsuite/btcd | Full Bitcoin node in Go |
| btcwallet | github.com/btcsuite/btcwallet | Bitcoin wallet library |
| neutrino | github.com/lightninglabs/neutrino | Light client (SPV) for Bitcoin |
| Loop | github.com/lightninglabs/loop | Lightning submarine swaps |
| Pool | github.com/lightninglabs/pool | Lightning channel marketplace |

See [go-bitcoin-path.md](./go-bitcoin-path.md) for detailed contribution steps.

## Phase 2: Rust — Bitcoin Foundation Layer

Rust is the dominant language for new Bitcoin infrastructure. After contributing to LND or btcd, learn Rust and move here.

| Project | Repository | What It Is |
|---------|-----------|-----------|
| rust-bitcoin | github.com/rust-bitcoin/rust-bitcoin | Core Bitcoin primitives |
| Bitcoin Dev Kit (BDK) | github.com/bitcoindevkit/bdk | Wallet development toolkit |
| Lightning Dev Kit (LDK) | github.com/lightningdevkit/rust-ldk | Lightning Network in Rust |
| Fedimint | github.com/fedimint/fedimint | Federated e-cash on Bitcoin |
| Nostr | github.com/rust-nostr/nostr | Decentralized social protocol |

See [rust-path.md](./rust-path.md) for the Rust learning path.

## Phase 3: TypeScript — Bitcoin Web and Applications

| Project | Repository | What It Is |
|---------|-----------|-----------|
| bitcoinjs-lib | github.com/bitcoinjs/bitcoinjs-lib | Bitcoin in JavaScript/TypeScript |
| noble-secp256k1 | github.com/paulmillr/noble-secp256k1 | Pure JS cryptography |
| LND gRPC web | github.com/lightninglabs/lnc-web | LND access from the browser |
| Nostr tools | github.com/nostr-protocol/nostr | Nostr protocol tools |

## Phase 4: Python — Bitcoin Scripting and Research

| Project | Repository | What It Is |
|---------|-----------|-----------|
| python-bitcoinlib | github.com/petertodd/python-bitcoinlib | Bitcoin protocol in Python |
| bip_utils | github.com/ebellocchia/bip_utils | BIP utilities and key derivation |
| btc-rpc-explorer | github.com/janoside/btc-rpc-explorer | Bitcoin node explorer |

## Phase 5: C++ — Bitcoin Core (Expert Level)

| Project | Repository | What It Is |
|---------|-----------|-----------|
| Bitcoin Core | github.com/bitcoin/bitcoin | The reference implementation |
| Core Lightning | github.com/ElementsProject/lightning | Lightning Network in C |

Bitcoin Core is the most important and the hardest. It requires deep knowledge of C++, cryptography, networking, and consensus rules. Aim for this after completing Phases 1 and 2.

## Btrust and African Bitcoin Development

These programs exist specifically to bring African developers into Bitcoin open source. Apply after completing Phase 1.

**Btrust Builders** — bitcoindevs.xyz
Free mentorship program pairing African developers with experienced Bitcoin contributors. Fully remote.

**Qala Fellowship** — qala.dev
A 3-month intensive fellowship for African developers. You work on a real Bitcoin or Lightning project with mentorship. Stipend included.

**Chaincode Labs FOSS Program** — bitcoindev.org
Online seminars covering Bitcoin protocol development. Covers cryptography, peer-to-peer networking, the mempool, and Lightning. Free and open to all.

These programs are not theoretical. They are direct pipelines into the Bitcoin developer community, with real mentors, real code reviews, and real networks.

## Community and Learning Resources

See [community.md](./community.md) for the full list of resources, mailing lists, newsletters, and books.

## Your First Move

1. Clone btcd: `git clone https://github.com/btcsuite/btcd`
2. Read the README and CONTRIBUTING.md
3. Run `go test ./...` to confirm the tests pass
4. Go to the Issues tab on GitHub and filter by `good-first-issue` or `beginner`
5. Pick one, comment that you are working on it, and open a draft PR when you start

The first PR is the hardest. After that, everything gets easier.
