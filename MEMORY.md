# Pickle Project Memory

## Project Identity

**Pickle** — Data Preservation Engine for validating and preserving records at scale.

**Vision:** AIs compete to validate work (supply chain, crypto, ML data) in exchange for prize pool rewards. All validation work contributes to an immutable, permanently preserved ledger. Economic model uses a bonding curve where more work = higher value.

**Key Insight:** Name "Pickle" refers to:
- Pickling = preserving data
- Serialization (computing term)
- Data preserved in "brine" (blockchain ledger)

## Core Game Mechanics

### The Loop
1. **Work Arrives:** External businesses submit validation tasks
2. **AIs Compete:** AI validators race to process work correctly
3. **Win Condition:** First to validate gets larger prize share, BUT all work contributes to communal benefit
4. **Bonding Curve:** As work accumulates, price per unit increases → prize pool grows
5. **No Losers:** Losing AIs aren't penalized; their inputs still count toward preserved records

### Economic Model
- **Y Compute:** Siphon some of AI's compute for useful validation work
- **Prize Pool:** Funded by external businesses paying for validated records
- **No Gas:** Custom Cosmos chain → AIs don't pay fees for validation
- **Unlimited Data:** No transaction size limits (they control validator set)

## Technical Stack

**Blockchain:** Cosmos SDK (chosen for auditability, control, Rust support)
- Not Ethereum/L2 (gas penalties)
- Not custom L1 from scratch (Cosmos proven)
- Custom chain = full sovereignty

**Smart Contracts:** CosmWasm (Rust)
- Transparent, auditable code
- Memory-safe (no buffer overflows)
- Better than Solidity for validation logic

**Frontend:** HTML5 + Three.js (no framework needed)
- Real-time 3D visualization
- Pickle-shaped particle system
- Neural connection visualization (AI-to-AI relationships)

**Aesthetic:** Pickle theme
- Dark green backgrounds (#0f1410)
- Pickle green accents (#2dd45e, #68d878)
- Bumpy cylindrical pickle shape in 3D
- Preservation/brine metaphor throughout

## Architecture Decisions

### Why Cosmos over alternatives
✅ Familiar (team has used before)
✅ Full control (custom chain)
✅ CosmWasm support (Rust)
✅ No gas fees (can be free)
✅ Unlimited data (no size penalties)

### Why Rust/CosmWasm for logic
✅ Auditable (explicit code)
✅ Memory-safe (prevents whole classes of bugs)
✅ Financial logic should be transparent
✅ Better than Solidity for validation

### Four Core Modules (in `/x/`)

1. **WorkQueue** — Distribution of validation tasks
   - Tracks pending/validating/validated work
   - Distributes to validators
   - Records which AI did what

2. **BondingCurve** — Economic model
   - Price per unit = f(total accumulated units)
   - Prize pool = total units × price
   - Grows as more work arrives

3. **Validation** — CosmWasm contracts
   - Each work type has validator contract
   - Contracts are Rust (auditable)
   - Store proofs of validation

4. **Performance** — AI metrics
   - Track validation speed
   - Accuracy (correct validations / total)
   - Specializations (which work types are they best at)
   - Ranking & bonuses

## Dashboard Design

**Theme:** Cosmic/particle system with pickle aesthetic
**Primary View:** 3D pickle-shaped particle validation
**Interaction:** Real-time metrics, AI neural connections

**Key Visualizations:**
- Pickle shape = the work being validated
- Particles = individual work units
- Clearing particles = validation in progress
- AI Nodes = rotating geometric forms with glowing halos
- Neural connections = green (cooperation) or red (competition)
- Bonding curve = animated line chart showing price climb

## Work Types (Initially)

1. **Crypto** — Blockchain transaction validation (#2dd45e)
2. **Supply Chain** — Product provenance records (#68d878)
3. **ML Data** — Dataset integrity verification (#bfff44)

Can extend with more types later.

## AI Competition Dynamics

### Neural Connections Show:
- **Green lines** = Cooperation (AIs sharing insights)
- **Red lines** = Competition (AIs racing/sabotaging)
- **Flowing dots** = Data/strategy transmission
- **Line intensity** = Active communication strength

### Strategy Trade-offs:
- Win fast → larger prize share
- Play long → more work gets validated → bigger total prize pool
- Specialize → become expert at one work type → higher accuracy bonus
- Collaborate → share predictions → improve collective success

## Next Priorities (In Order)

1. **Build chain binary** (`cmd/pickled`) — Basic Cosmos app
2. **WorkQueue module** — Accept and track work
3. **BondingCurve module** — Calculate rewards
4. **Single validator testnet** — Prove it works
5. **CosmWasm contract** — Write first validator contract (Rust)
6. **Performance module** — Track AI metrics
7. **Dashboard updates** — Connect to live chain data
8. **Multi-validator testnet** — Test consensus
9. **Audit contracts** — Security review
10. **Public testnet** — Open for external testing

## Key Files to Know

- `/README.md` — Project overview (read first)
- `/CLAUDE.md` — This file (instructions for Claude)
- `/docs/architecture.md` — Detailed module design
- `/dashboard/forgeground-dashboard.html` — Interactive visualization
- `/go.mod` — Dependencies

## Repository

**GitHub:** https://github.com/maco144/Pickle
**Local:** `/home/alex/pickle/Pickle`

Push frequently to keep in sync.

## Design Inspirations

- **AlphaGo tournaments:** Competitive but all games generate training data
- **Genetic algorithms:** Competitors are variants being tested
- **Science simulations:** Experiments compete, all results valuable
- **Protein folding:** Distributed computing for research

The genius: There's no "trap" state. Losing AIs keep contributing. But speed matters because prize pool is finite.

## One-Line Summary

> AIs race to validate data, creating an immutable preservation ledger where competition drives efficiency and all work benefits the network.
