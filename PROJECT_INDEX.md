# Pickle - Project Index

**Generated:** February 18, 2026
**Project Type:** Cosmos SDK Blockchain
**Language:** Go (backend), Rust (smart contracts), HTML/JS (frontend)

---

## 📋 Quick Summary

**Pickle** is a Cosmos SDK blockchain for validating and preserving data at scale. AI validators compete to process validation work (supply chain records, crypto transactions, ML datasets) in exchange for rewards from dynamically-priced bonding curves. All validation contributes to a permanent, immutable ledger secured on-chain.

**Tech Stack:**
- **Backend:** Cosmos SDK v0.50 (Go 1.24)
- **Smart Contracts:** CosmWasm (Rust) for auditable validation logic
- **Frontend:** HTML5 + Three.js interactive dashboard
- **Protocol:** Protocol Buffers (protobuf) for type definitions

---

## 📁 Project Structure

```
pickle/
├── app.go                          # Chain app initialization
├── cmd/
│   └── pickled/                   # CLI binary entry point
├── x/                             # Cosmos SDK modules
│   ├── workqueue/                 # Work submission & tracking (IMPLEMENTED)
│   ├── bondingcurve/              # Economic model (scaffolding)
│   ├── validation/                # CosmWasm validation contracts (scaffolding)
│   └── performance/               # AI metrics & rewards (scaffolding)
├── contracts/                     # CosmWasm contracts (Rust) - TBD
├── dashboard/                     # Web-based visualization
│   └── forgeground-dashboard.html # Real-time data flow visualization
├── proto/                         # Protobuf definitions
├── docs/                          # Architecture & design docs
├── Makefile                       # Build targets
├── CLAUDE.md                      # Project guidelines & conventions
├── MEMORY.md                      # Session memory (auto-generated)
├── go.mod / go.sum               # Go dependencies
├── buf.yaml / buf.gen.yaml       # Protobuf code generation config
└── README.md                      # User-facing documentation
```

---

## 🚀 Entry Points

### CLI Binary
- **Path:** `cmd/pickled/main.go`
- **Purpose:** Cosmos chain binary for running Pickle nodes
- **Commands:** Standard Cosmos SDK commands (chain init, transactions, queries)
- **Usage:** `pickled` with Cosmos SDK subcommands

### Chain App
- **Path:** `app.go`
- **Purpose:** Application initialization and module registration
- **Key Exports:** `NewApp()`, `DefaultNodeHome`
- **Modules:** Registers workqueue, bondingcurve, validation, performance modules

### Dashboard
- **Path:** `dashboard/forgeground-dashboard.html`
- **Purpose:** Real-time visualization of data flowing through the system
- **Tech:** HTML5 + Three.js (no build step needed)
- **Usage:** Open directly in browser

---

## 📦 Core Modules

### Module: WorkQueue
- **Path:** `x/workqueue/`
- **Status:** ✅ IMPLEMENTED (2,495 LOC)
- **Purpose:** Core module for work submission and tracking
- **Key Components:**
  - `keeper/keeper.go` — Work state management
  - `types/workqueue.pb.go` — Work unit definitions (protobuf)
  - `types/events.go` — Work lifecycle events
  - `client/cli/` — CLI commands for work submission
- **Key Types:**
  - `WorkUnit` — Individual work task
  - `MsgSubmitWork` — Submit work transaction
  - `WorkStatus` — Work state tracking
- **Query Server:** `keeper/query_server.go` (gRPC queries)
- **Msg Server:** `keeper/msg_server.go` (transaction handlers)

### Module: Bonding Curve
- **Path:** `x/bondingcurve/`
- **Status:** 📋 Scaffolding
- **Purpose:** Dynamic pricing and prize pool economics
- **Expected Components:**
  - State management for curve parameters
  - Price calculation functions
  - Reward distribution logic

### Module: Validation
- **Path:** `x/validation/`
- **Status:** 📋 Scaffolding
- **Purpose:** CosmWasm contract integration for work validation
- **Expected Components:**
  - Contract storage and execution
  - Validation result tracking

### Module: Performance
- **Path:** `x/performance/`
- **Status:** 📋 Scaffolding
- **Purpose:** AI validator metrics and reward calculations
- **Expected Components:**
  - Speed/accuracy tracking
  - Specialization scoring
  - Leaderboard calculations

---

## 🔧 Configuration & Build

### Makefile Targets
```bash
make build          # Build pickle chain binary → ./bin/pickled
make install        # Install to $GOPATH/bin
make test           # Run unit tests
make proto          # Generate protobuf code (buf generate)
make testnet        # Start single-validator testnet
make clean          # Remove build artifacts
make docker-build   # Build Docker image
make help           # Show all targets
```

### Protocol Buffer Configuration
- **Config:** `buf.yaml` — Buf lint/format settings
- **Generation:** `buf.gen.yaml` — Code generation rules
- **Proto Files:** `proto/` directory (modules define their own .proto files)
- **Generated Code:** `x/*/types/*.pb.go` (auto-generated, do not edit)

### Dependencies (Go)
- `cosmossdk.io/` (store, log, collections)
- `github.com/cosmos/cosmos-sdk` v0.50.8
- `github.com/cometbft/cometbft` v0.38.12 (consensus engine)
- `google.golang.org/grpc` (gRPC for queries)
- `github.com/spf13/cobra` (CLI framework)

---

## 📚 Documentation

### Architecture
- **Path:** `docs/architecture.md`
- **Topics:**
  - System design and module interactions
  - Data flow from work submission to validation
  - Economic model overview
  - State management strategy

### Project Guidelines
- **Path:** `CLAUDE.md`
- **Contents:**
  - Git workflow conventions
  - Code organization rules
  - Module development pattern
  - Testing strategy
  - Naming conventions (Go: camelCase functions, PascalCase types)

### Session Memory
- **Path:** `MEMORY.md`
- **Purpose:** Persistent notes across coding sessions
- **Updated by:** Claude Code agent (auto-generated)

### User Documentation
- **Path:** `README.md`
- **Contents:**
  - Vision and overview
  - Getting started guide
  - Architecture diagram
  - Directory structure
  - Contributing guidelines

---

## 🧪 Testing

### Current Test Coverage
- **Total Test Files:** 0 (tests not yet written)
- **Status:** Tests are required per CLAUDE.md guidelines
- **Strategy:** Unit tests in `*_test.go`, integration tests in `tests/`

### Test Commands
```bash
make test                    # Run all tests
go test -v ./x/workqueue/... # Test specific module
```

---

## 🔗 Key Design Patterns

### Module Structure (per CLAUDE.md)
Each Cosmos SDK module follows this structure:
```
module/
├── types/
│   ├── types.go              # Core type definitions
│   ├── messages.go           # Transaction message types
│   ├── *.pb.go               # Protobuf generated code
│   ├── keys.go               # State store keys
│   ├── events.go             # Event definitions
│   └── codec.go              # Register types for encoding
├── keeper/
│   ├── keeper.go             # State management
│   ├── msg_server.go         # Transaction handlers
│   ├── query_server.go       # Query handlers
│   └── genesis.go            # Genesis state initialization
├── client/cli/
│   ├── tx.go                 # Transaction command builders
│   └── query.go              # Query command builders
└── module.go                 # Module registration
```

### State Management
- **Keys:** Defined in `types/keys.go` (no hardcoding)
- **Storage:** Uses Cosmos SDK `sdk.KVStore`
- **Collections:** Cosmos SDK collections pattern for complex state

### Transaction Flow
1. User submits transaction via CLI or API
2. `MsgServer` in keeper processes it
3. State updated via keeper methods
4. Events emitted (defined in `types/events.go`)
5. Result committed to blockchain

---

## 💡 Development Workflow

### Adding a New Feature
1. **Design:** Define types in `.proto` or `types.go`
2. **Messages:** Create `MsgYourFeature` struct with `ValidateBasic()`
3. **Keeper:** Add state management method to `keeper.go`
4. **Handler:** Register message handler in `msg_server.go`
5. **CLI:** Add command in `client/cli/tx.go`
6. **Test:** Write unit tests alongside code
7. **Commit:** Use descriptive commit message (focus on "why")

### Running the Chain Locally
```bash
make build                    # Build binary
./scripts/testnet.sh         # Start single-validator testnet
# Or manually:
pickled config init home=~/.pickle
pickled start
```

### Querying the Chain
```bash
pickled q workqueue list-work    # Query work units
pickled tx workqueue submit-work ...
```

---

## 🔐 Security & Auditability

### Design Principles (from CLAUDE.md)
- **Rust-first:** Critical financial logic in CosmWasm (auditable Rust)
- **Determinism:** No randomness in critical logic (all validators reach same state)
- **Transparency:** All work and rewards tracked on-chain
- **Auditability:** Every decision recorded with cryptographic proof

### Best Practices
- ✅ Validate all inputs in `ValidateBasic()`
- ✅ Use params module for configurable values (no hardcoding)
- ✅ Store all state on-chain (no hidden state)
- ✅ Document complex logic thoroughly

---

## 📊 Module Status Summary

| Module | Status | LOC | Purpose |
|--------|--------|-----|---------|
| workqueue | ✅ Implemented | 2,495 | Work submission & tracking |
| bondingcurve | 📋 Scaffolding | ~0 | Dynamic pricing & rewards |
| validation | 📋 Scaffolding | ~0 | Contract validation logic |
| performance | 📋 Scaffolding | ~0 | AI metrics & scoring |

---

## 🎯 Next Steps (High Priority)

From git history and CLAUDE.md guidelines:
1. **Implement bonding curve module** — Economic model for work pricing
2. **Deploy validation contracts** — CosmWasm contracts for work verification
3. **Implement performance tracking** — AI validator metrics and leaderboard
4. **Write integration tests** — Test full work submission → validation flow
5. **Dashboard refinement** — Real-time updates from chain events
6. **Testnet deployment** — Run multi-validator testnet

---

## 📖 Quick Reference

### Important Paths
| Path | Purpose |
|------|---------|
| `app.go` | App initialization & module registration |
| `cmd/pickled/main.go` | CLI entry point |
| `x/workqueue/` | Main module (implemented) |
| `dashboard/forgeground-dashboard.html` | Frontend visualization |
| `CLAUDE.md` | **Read this first for conventions** |
| `docs/architecture.md` | **Read this for system design** |

### Key Commands
```bash
make build && make testnet        # Build and run testnet
make test                         # Run tests
pickled tx workqueue submit-work  # Submit work
pickled q workqueue list-work     # Query work units
```

### Module Development
See `CLAUDE.md` section "Module Development Pattern" for the exact template to follow.

---

## 🔗 External Resources

- [Cosmos SDK Docs](https://docs.cosmos.network/)
- [CosmWasm Docs](https://docs.cosmwasm.com/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [gRPC](https://grpc.io/docs/)

---

**Last Updated:** February 18, 2026
**Index Version:** 1.0
**Codebase Size:** ~2.5K lines (workqueue), ~100K dependencies (go.sum)
