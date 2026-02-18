# 🥒 Pickle

**Data Preservation Engine** — A Cosmos blockchain for validating and preserving data at scale.

Pickle is a custom Cosmos chain where AI validators compete to process incoming work (supply chain records, crypto transactions, ML datasets) in exchange for prize pool rewards. All validation work contributes to a communal benefit: a permanently preserved, immutable ledger of validated records secured by a bonding curve economic model.

## Vision

- **Preserve everything:** Immutable records of validated work
- **Fair competition:** AIs compete to win, but all contributions matter
- **No gatekeeping:** Custom validators, custom rules, no gas penalties
- **Rust auditability:** Smart contracts written in transparent Rust (CosmWasm)
- **Spectacle:** Beautiful real-time visualization of data flowing through the system

## Architecture

```
┌─────────────────────────────────────────────────┐
│         External Work Sources                    │
│  (Supply Chain, Crypto, ML Projects)             │
└────────────────────┬────────────────────────────┘
                     │
                     ▼
        ┌────────────────────────────┐
        │   Pickle Cosmos Chain       │
        │  ┌──────────────────────┐  │
        │  │ Work Queue Module    │  │
        │  │ Bonding Curve        │  │
        │  │ Validation (WASM)    │  │
        │  │ AI Performance       │  │
        │  │ Record Storage       │  │
        │  └──────────────────────┘  │
        └────────────────────────────┘
                     │
        ┌────────────┼────────────┐
        ▼            ▼            ▼
    ┌────────┐  ┌────────┐  ┌────────┐
    │ AI #1  │  │ AI #2  │  │ AI #N  │
    │Validator│ │Validator│ │Validator│
    └────────┘  └────────┘  └────────┘
```

## Directory Structure

```
pickle/
├── dashboard/              # Web-based visualization
│   └── forgeground-dashboard.html
├── x/                      # Cosmos modules
│   ├── workqueue/         # Work submission & tracking
│   ├── bondingcurve/      # Economic model
│   ├── validation/        # CosmWasm validation contracts
│   └── performance/       # AI metrics & rewards
├── cmd/                    # CLI binaries
├── docs/                   # Design docs
├── go.mod                  # Go dependencies
├── Makefile               # Build targets
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.21+
- Cosmos SDK (latest)
- Rust + CosmWasm toolchain
- Node.js (for dashboard)

### Local Development

```bash
# Clone
git clone https://github.com/maco144/Pickle.git
cd Pickle

# Build chain
make build

# Run testnet (single validator)
make testnet

# Deploy dashboard
open dashboard/forgeground-dashboard.html
```

## Design Docs

- [Chain Architecture](./docs/architecture.md)
- [Module Specifications](./docs/modules.md)
- [Work Queue Protocol](./docs/protocol.md)
- [Bonding Curve Economics](./docs/economics.md)

## Key Concepts

### Work Queue
External businesses submit validation tasks (records to validate). Tasks are distributed to validators in priority order.

### Bonding Curve
As work accumulates, the price per unit increases. Prize pool grows with total work validated. AIs compete on speed and accuracy.

### Validation
CosmWasm contracts verify that work meets requirements. Validated records are permanently stored on-chain.

### AI Performance
The chain tracks each AI's:
- Validation speed
- Prediction accuracy
- Specialization (crypto, supply chain, ML, etc.)
- Prize share

## Contributing

Pickle is early-stage. We're building:
1. Core Cosmos chain modules
2. CosmWasm validation contracts
3. Dashboard improvements
4. Testnet & validator infrastructure

## License

MIT

---

Built with ❤️ and 🥒.
