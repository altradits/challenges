# Bitcoin Developer Community Resources

The Bitcoin developer community is distributed across mailing lists, IRC channels, GitHub, and in-person events. This is the map.

## Programs for African Developers

These are the highest-priority resources for an African developer entering Bitcoin open source.

### Qala Fellowship

**Website:** https://qala.dev

A 3-month fellowship for African developers wanting to contribute to Bitcoin and Lightning open source. You work on a real project, receive mentorship from experienced Bitcoin developers, and get paid a stipend.

Applications open periodically. Check the website. When you have one merged PR in a Go or Rust Bitcoin project, apply.

### Btrust Builders

**Website:** https://bitcoindevs.xyz

Btrust is a nonprofit funded by Twitter founder Jack Dorsey to support Bitcoin development in Africa. The Builders program pairs African developers with mentors in the Bitcoin open source community.

Free. Remote. Open to developers at all levels.

Apply as soon as you start working on your first issue in btcd or LND.

### Chaincode Labs FOSS Program

**Website:** https://bitcoindev.org

Online seminars covering Bitcoin protocol development. Topics include:
- Bitcoin Script and transaction structure
- The peer-to-peer network and mempool
- Lightning Network protocol design
- Cryptographic primitives used in Bitcoin

Free and open to everyone. The reading lists are available on GitHub at github.com/chaincodelabs/seminars even without enrolling.

## Newsletters and Publications

### Bitcoin Optech

**Website:** https://bitcoinops.org

Weekly newsletter covering:
- Changes to Bitcoin Core
- New BIPs (Bitcoin Improvement Proposals)
- Lightning Network protocol developments
- Wallet software updates
- Research papers

Reading Optech weekly is the single best way to stay current with Bitcoin development without reading every mailing list thread. The technical topics section is excellent for developers.

### Delving Bitcoin

**Website:** https://delvingbitcoin.org

A forum for deep technical discussion of Bitcoin protocol changes, research, and implementation. Less noisy than mailing lists. High signal-to-noise ratio.

## Mailing Lists

### Bitcoin-Dev

**URL:** https://lists.linuxfoundation.org/mailman/listinfo/bitcoin-dev

The primary mailing list for Bitcoin protocol development. New BIPs are first posted here. Protocol changes are discussed here. If you want to understand what is happening at the protocol level, read this list.

You do not need to post. Reading the archives alone is valuable.

### Lightning-Dev

**URL:** https://lists.linuxfoundation.org/mailman/listinfo/lightning-dev

The primary mailing list for Lightning Network development. BOLT (Basis of Lightning Technology) changes are discussed here.

## IRC and Chat

### Libera.Chat IRC

Many Bitcoin developers still use IRC. Channels:

- `#bitcoin-core-dev` — Bitcoin Core development discussion
- `#lnd` — LND development
- `#rust-bitcoin` — rust-bitcoin and BDK

Connect with any IRC client (HexChat, WeeChat, irssi) to irc.libera.chat.

### Bitcoin Core Slack

Invitation at https://bitcoincore.org/en/contact/

### LND Slack

Invitation at https://lightning.engineering/slack/

## Books

### Mastering Bitcoin (3rd Edition) — Free Online

**GitHub:** https://github.com/bitcoinbook/bitcoinbook

Andreas Antonopoulos's definitive guide to Bitcoin for developers. Covers:
- Transaction structure and Script
- Wallets and key management
- The peer-to-peer network
- Blockchain and consensus

Read this before contributing to btcd or Bitcoin Core.

### Mastering the Lightning Network — Free Online

**GitHub:** https://github.com/lnbook/lnbook

The equivalent guide for Lightning. Covers:
- Payment channels
- HTLCs and routing
- BOLT specifications
- Node management

Read this before contributing to LND.

### Programming Bitcoin — Jimmy Song

**URL:** https://programmingbitcoin.com

Implements Bitcoin from scratch in Python. The best book for understanding how the cryptographic primitives actually work. After reading, implement the same code in Go or Rust.

## BIPs — Bitcoin Improvement Proposals

**Repository:** https://github.com/bitcoin/bips

BIPs are the formal specification documents for Bitcoin protocol changes. Reading BIPs is how you understand what Bitcoin actually does at the protocol level.

Key BIPs to read early:
- BIP32 — Hierarchical Deterministic Wallets (key derivation)
- BIP39 — Mnemonic code (seed phrases)
- BIP141 — Segregated Witness
- BIP340 — Schnorr Signatures
- BIP341 — Taproot
- BIP343 — Tapscript
- BOLT1–BOLT11 — Lightning Network specifications (in the lightning/bolts repo)

## Events

### Bitcoin++ (btcplusplus.dev)

A developer-focused Bitcoin conference. Usually held in Austin, Texas and sometimes other cities. Workshops and technical talks. Attend if possible — you meet the people whose PRs you have been reviewing.

### Bitcoin Africa Conference

An annual conference focused on Bitcoin in Africa. More accessible geographically than Austin or Amsterdam events.

### TabConf

A smaller, more technical Bitcoin developer conference. Very high signal-to-noise ratio.

## GitHub Organizations to Follow

- **btcsuite** — github.com/btcsuite — btcd, btcwallet, and libraries
- **lightningnetwork** — github.com/lightningnetwork — LND
- **lightninglabs** — github.com/lightninglabs — Loop, Pool, Taproot Assets
- **rust-bitcoin** — github.com/rust-bitcoin — rust-bitcoin and related
- **bitcoindevkit** — github.com/bitcoindevkit — BDK
- **lightningdevkit** — github.com/lightningdevkit — LDK
- **fedimint** — github.com/fedimint — Fedimint
- **bitcoin** — github.com/bitcoin — Bitcoin Core

Watch these organizations on GitHub to see new issues and PRs as they appear.

## A Note on Community Culture

Bitcoin development culture is direct and technically rigorous. Code reviews are not personal. When a maintainer says "this needs to be rewritten," they are talking about the code, not you.

The culture rewards:
- Technical precision
- Patience in review
- Respect for the existing codebase
- Clear communication about what a change does and why

It does not reward:
- Large PRs with no discussion
- Bypassing review process
- Complaining about slow review times publicly

Earn trust by writing good code, responding thoroughly to review, and being consistently present over months — not weeks.
