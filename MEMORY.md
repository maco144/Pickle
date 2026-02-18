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

## Game Design (Latest Session)

**Core Insight:** The game IS the meta-optimization visualization for mining.

- **Match Structure:** 3D pickle-shaped RTS arena with 500K compressed business records
- **4 AI Archetypes:** Light (speed), Water (containment), Slime (defense), Fire (destruction)
- **Compute Allocation:** Each AI divides compute between RTS gameplay and real-time validation work
- **Dual Victory:** Territory control + validation count both matter for rewards
- **Validation:** Real business records (crypto, supply chain, ML) validated during game, signed, prepared for on-chain storage
- **Comebacks:** Losing game but validating more = can win overall (validates game theory balance)

**Key Files:**
- `/docs/game-design.md` — Full game design document (theory & mechanics)
- See "Open Design Questions" section for implementation challenges

## Next Priorities (In Order)

**Backend (Cosmos):**
1. **Define validation logic** per data type (crypto signatures, supply chain checks, etc.)
2. **Implement `x/game` module** — Match acceptance, game results
3. **Implement `x/bondingcurve`** — Reward calculation (territory + validation)
4. **Implement `x/performance`** — AI stats, specialization tracking
5. **Create synthetic block generator** — For testing without real data

**Frontend (Game):**
6. **3D visualization** — Pickle-shaped RTS arena with fluid dynamics
7. **Unit/structure system** — Game mechanics that consume compute
8. **Compute allocation UI** — Show game vs validation split
9. **Real-time validation integration** — Display records being validated during match

**Integration:**
10. **End-to-end test** — Single match from game → validation → rewards
11. **Testnet with real data** — When business record sources available

## Key Files to Know

- `/README.md` — Project overview (read first)
- `/CLAUDE.md` — Instructions for Claude (conventions & guidelines)
- `/PROJECT_INDEX.md` — Complete codebase map (read this for orientation)
- `/docs/architecture.md` — Module architecture & system design
- `/docs/game-design.md` — **LATEST:** Game mechanics, compute allocation, validation strategy
- `/dashboard/forgeground-dashboard.html` — Interactive 3D visualization
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
