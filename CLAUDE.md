# Claude Instructions for Pickle

## 🚀 Quick Start for Developers

**New to this project?** Start here:
1. Read `PROJECT_INDEX.md` — Complete project overview (entry points, modules, structure, next steps)
2. Read `docs/architecture.md` — System design and data flow
3. Review the development workflow below

**Why `PROJECT_INDEX.md`?** It's a living index that maps the entire codebase (modules, entry points, key files, status). Updated every session. Start here instead of exploring randomly.

---

## Project Overview

**Pickle** is a Cosmos SDK blockchain for validating and preserving data at scale. AIs compete to validate work (supply chain records, crypto transactions, ML datasets), with prize pools funded by businesses paying for verification. Built in Rust (CosmWasm) for auditability.

**Key Technologies:**
- Cosmos SDK (Go)
- CosmWasm (Rust smart contracts)
- Interactive dashboard (HTML5 + Three.js)

## Workflow & Conventions

### Git Practices
- **Always commit before making changes** — use descriptive messages
- **Branch if major features:** Use `feature/` prefix for branches
- **Pull before working** — Keep local up to date
- **Push after completing tasks** — No orphaned work
- **Commit messages:** Focus on "why" not just "what"
  - ✅ "Add WorkQueue module to handle task distribution"
  - ❌ "update code"

### Code Organization
- **Modules live in `/x/`** — Each module is self-contained (keeper, types, messages, handlers)
- **Contracts in `/contracts/`** — Rust CosmWasm contracts (one dir per contract)
- **Dashboard in `/dashboard/`** — Frontend visualization (HTML/JS/Three.js)
- **Tests alongside code** — `*_test.go` in same directory

### Naming Conventions
- **Go packages:** lowercase, short (`workqueue`, `bondingcurve`, `validation`, `performance`)
- **Go types:** PascalCase (`WorkUnit`, `BondingCurveState`, `ValidatorStats`)
- **Go functions:** camelCase (`submitWork`, `validateRecord`, `calculateReward`)
- **Rust:** Follow Rust conventions (snake_case for functions, PascalCase for types)

### Documentation
- **Architecture changes:** Update `/docs/architecture.md`
- **New modules:** Add brief `/docs/modules/<module>.md`
- **API endpoints:** Document in `/docs/api.md` (create if needed)
- **Contracts:** Rust doc comments (`///`) for all public functions

### Development Flow
1. **Orient yourself:** Read `PROJECT_INDEX.md` to understand module layout and current status
2. **Study the system:** Read the relevant architecture doc (`docs/architecture.md`) before coding
3. **Design first:** Types/messages on paper or in comments before implementing
4. **Implement:** With clear separation of concerns (keeper, types, messages, handlers)
5. **Test:** Unit tests (Go) and integration tests (bash scripts) alongside code
6. **Document:** As you go (especially complex logic)
7. **Commit:** With meaningful message (focus on "why", not just "what")
8. **Push:** Share progress when task is complete

## Module Development Pattern

When building a new module:

```go
// 1. Types (types.go)
type WorkUnit struct { ... }
type WorkStatus string

// 2. Messages (messages.go)
type MsgSubmitWork struct { ... }
func (m MsgSubmitWork) ValidateBasic() error { ... }

// 3. Keeper (keeper.go)
type Keeper struct { ... }
func (k Keeper) SubmitWork(ctx context.Context, msg *MsgSubmitWork) { ... }

// 4. Handler (handler.go)
func NewMsgServerImpl(k Keeper) MsgServer { ... }

// 5. Tests (types_test.go, keeper_test.go, etc.)
func TestWorkSubmission(t *testing.T) { ... }
```

## Testing Strategy

**Go modules:**
- Unit tests for types and business logic
- Keeper integration tests
- Message validation tests

**Contracts (Rust):**
- Unit tests in contracts (use `cosmwasm-std/testing`)
- Integration tests in `/tests/`
- Fuzzing for financial logic

**End-to-end:**
- Testnet scripts in `/scripts/`
- Manual testing on local devnet

## Key Design Principles

1. **Rust first:** Logic that needs auditing goes in CosmWasm contracts (visible Rust)
2. **Simplicity:** Start minimal, add complexity only when needed
3. **Determinism:** All validators must reach same state (no randomness in critical logic)
4. **Transparency:** All work and rewards tracked on-chain (no hidden state)
5. **Auditability:** Every decision recorded with proof

## Common Tasks

### Add a new message to a module
1. Define message in `types.go` or `messages.go`
2. Add handler in `handler.go`
3. Add keeper method
4. Write tests
5. Update README if it changes public API

### Deploy a new CosmWasm contract
1. Write Rust in `/contracts/<name>/`
2. Build: `cargo wasm`
3. Validate: `wasm-validator contract.wasm`
4. Store on-chain via `SubmitContractCode` message
5. Document contract interface

### Run testnet
```bash
make testnet
# Or manually:
./scripts/testnet.sh
```

### Build dashboard changes
- Edit `/dashboard/forgeground-dashboard.html`
- No build step needed (pure HTML/JS)
- Commit and push
- Open in browser to test

## What NOT to Do

- ❌ Don't hardcode values (use params module)
- ❌ Don't store mutable state in contracts (use chain state)
- ❌ Don't write financial logic in Go (put it in auditable Rust contracts)
- ❌ Don't commit without testing (at least `make test`)
- ❌ Don't push breaking changes without a migration plan
- ❌ Don't skip documentation for complex logic

## Useful Commands

```bash
# Build
make build
make install

# Test
make test
go test ./...

# Cosmos-specific
ignite chain build                    # (if using Ignite CLI)
pickled config init                   # Initialize node
pickled tx submit-work ...            # Submit work to chain

# CosmWasm
cargo build --release --target wasm32-unknown-unknown
wasm-validator path/to/contract.wasm
```

## References

### Internal Documentation
- **`PROJECT_INDEX.md`** — Complete codebase map (modules, entry points, file structure, next steps)
- **`PROJECT_INDEX.json`** — Machine-readable version for tooling
- **`docs/architecture.md`** — System design, module interactions, data flow
- **`MEMORY.md`** — Session memory and persistent notes

### External Resources
- **Cosmos SDK Docs:** https://docs.cosmos.network/
- **CosmWasm Docs:** https://docs.cosmwasm.com/
- **Protocol Buffers:** https://developers.google.com/protocol-buffers
- **gRPC:** https://grpc.io/docs/

## Emergency: Something Broke

1. **Chain won't start:** Check `/docs/architecture.md` for state assumptions
2. **Test failures:** Run full test suite (`make test`)
3. **Contract errors:** Check CosmWasm error codes and Rust docs
4. **Git disaster:** Your changes are safe in `.git/` — ask for help reverting

---

**Questions?** Check the docs first, then refer to the architecture design.
