# Rust Learning Path for Bitcoin Developers

Rust is the dominant language for new Bitcoin infrastructure. If Go is where you start contributing, Rust is where you go deep.

## Why Rust for Bitcoin

- **Memory safety without a garbage collector** — Bitcoin software handles real money. Rust eliminates entire classes of bugs (buffer overflows, use-after-free, data races) at compile time.
- **Performance** — Rust code is as fast as C but far safer. For cryptographic operations and transaction processing, this matters.
- **The ecosystem is moving here** — rust-bitcoin, BDK, and LDK are the foundations of the next generation of Bitcoin tooling. New wallets, new Lightning implementations, and new layer-2 protocols are being built in Rust.
- **Fedimint and Nostr** — Two of the most exciting Bitcoin ecosystem projects (federated e-cash and decentralized social) are written in Rust.

## The Rust Learning Path

### Step 1: The Rust Book (2–4 weeks)

**URL:** https://doc.rust-lang.org/book/

The official Rust book is free, online, and excellent. It is the standard starting point.

Key chapters to not skip:
- Chapter 4: Ownership — this is what makes Rust different from every other language
- Chapter 10: Generics, Traits, Lifetimes — essential for reading any library code
- Chapter 15: Smart Pointers — you will see `Arc`, `Rc`, `Box` constantly in Bitcoin code
- Chapter 16: Concurrency — Bitcoin nodes are concurrent systems

Do the exercises in each chapter. Do not just read.

### Step 2: Rustlings (1 week)

**URL:** https://github.com/rust-lang/rustlings

Small exercises that force you to fix Rust code. Run them after the first half of the Rust Book. They give you the muscle memory that reading alone does not.

```bash
cargo install rustlings
rustlings init
rustlings watch
```

### Step 3: Programming Bitcoin (2–3 weeks)

**Book:** "Programming Bitcoin" by Jimmy Song

The book uses Python to implement Bitcoin from scratch: elliptic curve cryptography, transaction parsing, Script, blocks, and Merkle trees. The Python is simple — the Bitcoin is the point.

After reading each chapter, implement the same concept in Rust. This gives you:
- Deep understanding of Bitcoin primitives
- Practice writing real Rust code
- A portfolio of Bitcoin implementations

**Buy it:** programmingbitcoin.com  
**Free on GitHub (older edition):** github.com/jimmysong/programmingbitcoin

### Step 4: rust-bitcoin — Read First, Then Contribute

**Repository:** https://github.com/rust-bitcoin/rust-bitcoin

rust-bitcoin is the foundational Rust library for Bitcoin. It provides:
- Transaction and block parsing and serialization
- Script types and Script builder
- Address encoding and decoding (Legacy, P2SH, Bech32, Taproot)
- Network types and constants
- BIP32 key derivation

**Start by reading:**
- `bitcoin/src/blockdata/transaction.rs` — transaction structure
- `bitcoin/src/blockdata/script/` — Script types
- `bitcoin/src/address/` — address parsing and generation
- `CONTRIBUTING.md` — the onboarding guide

**Issues to look for:**
- `good-first-issue` label
- `documentation` label
- Test coverage gaps

### Step 5: Bitcoin Dev Kit (BDK)

**Repository:** https://github.com/bitcoindevkit/bdk

BDK is a wallet development toolkit built on top of rust-bitcoin. It handles:
- Descriptor-based wallets (using output descriptors)
- Coin selection
- Transaction building and signing
- Blockchain data sources (Electrum, Esplora, Bitcoin Core RPC)

BDK is what developers use when they want to build a Bitcoin wallet in Rust. It is actively maintained with regular releases and a welcoming contributor community.

**Contributing:** https://github.com/bitcoindevkit/bdk/blob/master/CONTRIBUTING.md

### Step 6: Lightning Dev Kit (LDK)

**Repository:** https://github.com/lightningdevkit/rust-ldk

LDK is a Lightning Network library in Rust. Unlike LND (a complete daemon), LDK is designed to be embedded into wallets and other applications. It is used by Cash App and several other major Bitcoin products.

LDK is more complex than BDK. Aim for it after you are comfortable in rust-bitcoin and BDK.

## Key Rust Concepts for Bitcoin Code

### Ownership and the borrow checker

```rust
let tx = Transaction::new();
let txid = tx.txid();  // borrows tx
println!("{}", txid);  // still valid — txid is a reference
```

The borrow checker prevents you from accidentally modifying transaction data while holding a reference to it. In financial software, this matters.

### Error handling with Result

```rust
let tx: Result<Transaction, Error> = Transaction::consensus_decode(&mut cursor);
match tx {
    Ok(transaction) => println!("txid: {}", transaction.txid()),
    Err(e) => eprintln!("Failed to decode: {}", e),
}
```

Rust has no exceptions. Every function that can fail returns `Result<T, E>`. This forces you to handle every error case explicitly.

### Traits

```rust
use bitcoin::consensus::Encodable;

// Any type implementing Encodable can be serialized to the wire format
tx.consensus_encode(&mut writer)?;
```

Bitcoin types implement standard traits for serialization, hashing, and display. Understanding traits is essential for reading and writing Bitcoin library code.

## Rust Bitcoin Projects Worth Watching

| Project | Repository | Description |
|---------|-----------|-------------|
| rust-bitcoin | github.com/rust-bitcoin/rust-bitcoin | Core primitives |
| BDK | github.com/bitcoindevkit/bdk | Wallet toolkit |
| LDK | github.com/lightningdevkit/rust-ldk | Lightning library |
| Fedimint | github.com/fedimint/fedimint | Federated e-cash |
| Nostr | github.com/rust-nostr/nostr | Decentralized social |
| Payjoin | github.com/payjoin/rust-payjoin | Privacy-preserving transactions |
| BIP324 | github.com/bitcoin/bips/blob/master/bip-0324.mediawiki | Encrypted P2P transport |

## After Rust: Bitcoin Core (C++)

Bitcoin Core is the reference implementation. It is written in C++ and maintained by the most rigorous review process in Bitcoin development.

Contributing to Bitcoin Core requires:
- Deep C++ knowledge (C++17 and later)
- Understanding of Bitcoin's consensus rules
- Understanding of the peer-to-peer protocol
- Patience — review cycles can take months

Do not rush to Bitcoin Core. Contribute meaningfully to Rust and Go projects first. Build your reputation. Then, when you engage with Bitcoin Core, maintainers will already know who you are.

## Timeline Estimate

| Phase | Duration | Milestone |
|-------|---------|---------|
| The Rust Book + Rustlings | 3–5 weeks | Can write and read Rust |
| Programming Bitcoin in Rust | 3–4 weeks | Understand Bitcoin primitives |
| rust-bitcoin reading + first issue | 2–4 weeks | First rust-bitcoin PR open |
| BDK contribution | ongoing | Merged PR in BDK |
| LDK contribution | ongoing | Deep Lightning knowledge |

These are not deadlines. They are estimates. The most important thing is to start.
