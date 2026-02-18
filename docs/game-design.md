# Pickle Game Design Document

**Status:** Design Phase (Theory → Implementation)
**Last Updated:** February 18, 2026
**Owner:** Design Team

---

## 🎮 Game Concept Overview

**Pickle** is a 3D RTS game that visualizes AI resource allocation during blockchain record validation. AIs compete to validate business records (supply chain, crypto, ML datasets) while balancing game performance with validation throughput.

The game is **not** just decoration — it's a meta-optimization visualization where AIs solve a real constraint satisfaction problem: how to allocate compute between competitive gameplay and productive validation work.

---

## 🌍 Visual Design

### Arena
- **3D pickle-shaped grid space** filled with "work matter" (500K compressed business records)
- **State visualization:**
  - Start: Grey (unvalidated)
  - End: 2-colored 3D shape (each color = one AI's validated work)
- **Matter density:** Hyperbolic acceleration toward end (more records processed as game accelerates)

### AI Archetypes (Fluid Dynamics Inspired)
Each AI is represented as a fluid-dynamic entity with distinct mechanics:

#### 🔦 **Light**
- **Strength:** Speed, scouting, fast setup
- **RTS role:** Scouts and claims territory rapidly
- **Mechanic:** Spreads quickly, illuminates areas
- **Load time:** Fast (~0.3s for lead miner)
- **Weakness:** Needs density to hold territory (spreads thin)

#### 💧 **Water**
- **Strength:** Pooling, diffusion, gradual spread
- **RTS role:** Relentless spreading, containment
- **Mechanic:** Flows, gathers, absorbs Light over time
- **Load time:** Medium (~0.6s)
- **Interaction:** Diffuses over Light, containment

#### 🟢 **Slime**
- **Strength:** Anchoring, point control, defense
- **RTS role:** Holds key chokepoints
- **Mechanic:** Sticky, adheres, hard to displace
- **Load time:** Slow (~0.8s)
- **Weakness:** Immobile, slow expansion

#### 🔥 **Fire**
- **Strength:** Destructive, aggressive, high variance
- **RTS role:** Offensive capability, breaks entrenched positions
- **Mechanic:** Can't claim/hold territory directly (only damage)
- **Load time:** Fast (~0.4s)
- **Variance:** Unpredictable outcomes (high risk/high reward)

---

## 📊 Match Structure

### Duration
- **3-6 minutes** (recommend 3 min for esports pacing)
- **Hyperbolic acceleration:** Early game slow, late game explosive (as work depletes)

### Flow

#### Phase 1: Pre-Game (1 second setup)
```
MsgAcceptMatch
├── Player1 selects: [Miner1, Miner2]
├── Player2 selects: [Miner1, Miner2]
├── Both see: Block composition (unknown types but will be revealed)
└── Game state: PENDING → IN_PROGRESS
```

**Miner Selection Strategy:**
- Generalist vs Specialist tradeoff (spec 2x faster on type, generalist 1.2x on all)
- Load order choice (first miner loads fastest, gives early game edge)
- Example: Load 1 crypto specialist fast for early advantage, or 2 diverse miners slower but safer

#### Phase 2: Game Active (3 minutes)
**Two parallel processes:**

**A) RTS Territorial Control**
- Build units and structures (vary by class archetype)
- Expand territory, contest matter
- Visual representation of game dominance

**B) Real-Time Validation**
- Simultaneously validate records from the block
- Each validated record = signed + prepared for on-chain storage
- Independent of territorial control

#### Phase 3: Post-Game
```
MsgSubmitGameResult
├── Winner: Territorial control winner
├── Territory split: {Player1: X%, Player2: Y%}
├── Validation results: {Player1: RecordsValidated, Player2: RecordsValidated}
└── Game state: COMPLETE → VALIDATED
```

**Block Finalization:**
- Cluster validates results (post-game)
- Block minted on-chain with all 500K records validated + signed
- Rewards distributed instantly

---

## 💾 The Compute Allocation System

### Core Mechanic
Each AI has **finite compute per match**. They must decide how to split it:

```
Total Compute = Game Compute + Validation Compute

Each second during match:
├── Compute_to_game: Units, structures, pathfinding, territory push
└── Compute_to_validation: Decompression, validation logic, signing records
```

### Strategic Choices
This creates a real optimization problem with dynamic adaptation:

**When Winning:**
- Temptation: Secure the lead (allocate more to game)
- Risk: Opponent pivots to validation, catches up in total rewards
- Strategy: Balance — keep game pressure while gaining validation bonuses

**When Losing:**
- Temptation: Validate like crazy, forget territory
- Risk: Still lose game + validation score too low
- Strategy: Contest game defensively while validating opportunistically

**The Meta:** Best player wins by balancing offense (territory) and defense (validation), adapting to block composition.

### Example Allocation Scenario
```
Block: 60% Crypto, 40% Supply Chain (500K total)
Timeframe: 3 minutes

Alice (Light + CryptoSpecialist):
├── Game allocation: 70% compute → Claims 60% territory
├── Validation allocation: 30% compute → Validates 300 records
│   ├── Crypto records: 5ms/record (2x specialist bonus)
│   └── Supply chain: 15ms/record (1x generalist)
└── Total: 60% territory + 300 records validated

Bob (Water + SupplyChainSpecialist):
├── Game allocation: 50% compute → Claims 40% territory
├── Validation allocation: 50% compute → Validates 350 records
│   ├── Crypto records: 10ms/record (1x generalist)
│   └── Supply chain: 7.5ms/record (2x specialist bonus)
└── Total: 40% territory + 350 records validated
```

---

## 🔐 Validation Mechanics

### What is "Validating a Record"?

For each business record in the block:

1. **Decompress** (emergent-language reversal — <1ms)
2. **Identify type** (Crypto, Supply Chain, ML, etc.)
3. **Run validation logic** (type-dependent):
   - **Crypto TX:** Verify signature, check balance, verify nonce
   - **Supply Chain:** Verify chain of custody, validate timestamps, check provenance
   - **ML Dataset:** Schema validation, data quality checks, integrity checksums
4. **Sign with validator key** (proof AI validated it)
5. **Prepare for on-chain storage** (add to validated pool)

### Throughput Model

**Time per record varies by:**
- Record type (crypto = different complexity than supply chain)
- AI's miner specialization (specialist = 2x faster on their type)
- Total compute allocated to validation

**Formula:**
```
Records_per_second = Compute_allocated / Time_per_record

Time_per_record = BaseTime / Specialization_bonus

Examples:
├── Crypto generalist: 10ms/record
├── Crypto specialist (2x): 5ms/record
├── Supply chain generalist: 15ms/record
└── Supply chain specialist (2x): 7.5ms/record
```

---

## 💰 Reward Structure

### Components
Rewards come from two sources:

#### 1. Territory Bonus (Game Outcome)
```
Territory_reward = Bonding_curve_payout × Territory_percentage

Example: 1000 coin block
├── Winner (60% territory): 600 coins
└── Loser (40% territory): 400 coins
```

#### 2. Validation Bonus (Work Done)
```
Validation_reward = Records_validated × Per_record_payment

Example: Same 1000 coin block
├── AI1 validates 300 records: 300 × 1.33 = 400 coins
├── AI2 validates 200 records: 200 × 1.33 = 267 coins
└── (Assuming 500 per-record limit to keep bonding curve balanced)
```

#### Total Rewards
```
Total_reward = Territory_reward + Validation_reward

Example outcome:
├── Alice: 600 (territory) + 400 (validation) = 1000 coins ✓ Winner overall
├── Bob: 400 (territory) + 267 (validation) = 667 coins
└── (Alice dominated both game and validation)

Counter-example (validation comeback):
├── Charlie: 200 (territory) + 600 (validation) = 800 coins ✓ Wins overall despite losing game
└── Diana: 800 (territory) + 100 (validation) = 900 coins (lost on validation count)
```

### Dynamics
- **Pure game-focused:** High territory, low validation → lose if opponent validates more
- **Balanced:** Moderate both → most consistent
- **Validation-focused:** Low territory, high validation → can win overall despite losing game
- **Rock-paper-scissors:** Different archetypes have natural advantages in different scenarios

---

## 🏗️ Class Dynamics & Matchups

### Strategic Archetypes

| Matchup | Dynamic | Meta |
|---------|---------|------|
| **Light vs Water** | Speed vs Containment | Light must expand fast before being boxed in |
| **Water vs Slime** | Diffusion vs Anchoring | Slime holds chokepoints, Water flows around |
| **Slime vs Fire** | Defense vs Aggression | Slime structures resist Fire, but Fire breaks defenses |
| **Fire vs Light** | Chaos vs Speed | Light's scattered claims vulnerable to Fire AoE |
| **Light vs Slime** | Speed vs Defense | Light scouts fast but Slime fortifies |
| **Water vs Fire** | Spread vs Destruction | Fire disrupts but Water keeps coming |

### Compute Allocation by Archetype

**Light:** Tends toward game-focused (speed advantage in RTS)
- Fast miner load means early game advantage
- Can leverage speed to claim territory before validating

**Water:** Tends toward balanced
- Slower load but steady—good for sustained validation
- Natural at multi-tasking (spreading + validating)

**Slime:** Tends toward validation-focused
- Defensive gameplay uses less compute (hold position)
- Extra compute for validation while defending

**Fire:** High variance
- Aggressive phases (all-in game push) vs defensive phases (validate while holding)
- Explosive moments when committing to territory or validation

---

## 🗺️ Map & Block Design

### Block Composition
**Data sources:** Real business records (TBD sourcing strategy)
- Unknown mix of data types until game reveals
- Example: 60% supply chain, 40% crypto, 0% ML

### Synthetic Blocks (MVP)
For testing without real data:
- Procedurally generated or fixture-based
- Equal distribution: 33% each type (balanced for testing)
- Can vary difficulty/types to test specialization

### Future: Real Data Integration
When business data sources available:
- Actual distribution varies per block
- Miners that specialized for wrong type disadvantaged
- Creates need for adaptation + meta-learning

---

## 🔗 Blockchain Integration (Cosmos Backend)

### Architecture Decisions

**Records:** On-chain (stored in `x/workqueue`)
- 500K compressed records per block
- Private Cosmos chain (no gas/size limits)

**Game State:** Minimal on-chain
- Store only: `{winner, territory_split, records_validated}`
- Not: Full event log (would be too much data)

**Validation:** Secondary (off-chain orchestrated, results submitted)
- Cluster validates post-game
- Results submitted via `MsgSubmitValidationResults`

**Rewards:** Instant
- Calculated immediately upon validation submission
- Distributed to winner and loser proportionally

### Message Flow

```
1. MsgAcceptMatch
   ├── Players select miners
   └── Block assigned

2. Game runs off-chain (3 min)
   ├── RTS simulation
   ├── In-game validation
   └── No on-chain activity

3. MsgSubmitGameResult
   ├── Winner, territory split
   ├── Records validated by each AI
   └── Game state stored

4. MsgSubmitValidationResults (off-chain triggered)
   ├── Block hash, validation success
   ├── Final record count
   └── Validation rewards calculated

5. Rewards distributed
   ├── Block minted
   └── Next match begins
```

### Modules Involved

- **`x/game`** (NEW): Match acceptance, game results
- **`x/workqueue`** (ENHANCE): Block storage, assignment
- **`x/validation`** (ENHANCE): Validation orchestration, results
- **`x/performance`** (IMPLEMENT): AI stats, specialization tracking
- **`x/bondingcurve`** (IMPLEMENT): Reward calculation, pool management

---

## 🎯 Open Design Questions

### Validation & Compute
1. What's the exact validation logic per data type? (Signatures? Rules? Checksums?)
2. How is compute-to-throughput mapped? (Records per second?)
3. Can AIs see opponent's validation progress in real-time or only post-game?

### Game Mechanics
1. What are actual units/structures in RTS? (Builders? Harvesters? Cannons?)
2. How much does validating a record slow down game performance?
3. Can miners be swapped/adjusted mid-game or locked at start?

### Economics
1. How are territory bonus and validation bonus weighted? (60/40? Dynamic?)
2. What's the bonding curve formula? (Linear? Exponential?)
3. Are there slashing conditions for failed validation?

### Data
1. Where do real business records come from? (API? Direct submission? Batches?)
2. Can blocks be specialized (all crypto) or always mixed?
3. How is record difficulty determined?

---

## 📈 Next Implementation Phases

### Phase 1: Backend Architecture (In Progress)
- [ ] Design message types (MsgAcceptMatch, MsgSubmitGameResult, etc.)
- [ ] Implement `x/game` module
- [ ] Implement `x/bondingcurve` module (reward calculation)
- [ ] Implement `x/performance` module (AI stats tracking)
- [ ] Create synthetic block generator for testing

### Phase 2: Game Engine
- [ ] 3D visualization (pickle grid, fluid dynamics)
- [ ] RTS unit/structure mechanics
- [ ] Territorial control calculation
- [ ] Miner loading system (compute allocation)

### Phase 3: Validation Integration
- [ ] Define validation logic per data type
- [ ] Implement validation executor (runs in-game)
- [ ] Integrate with blockchain (submit results)
- [ ] Track and reward validation work

### Phase 4: Polish & Testing
- [ ] Playtesting (game balance, reward fairness)
- [ ] Performance optimization (hyperbolic acceleration)
- [ ] Leaderboard/stats display
- [ ] Real data source integration

---

## 📝 Design Principles

1. **Game visualizes optimization problem:** Territory and validation are the solution to "how do I allocate compute?"
2. **Real work, real rewards:** All validated records are actual business data that gets stored on-chain
3. **Transparent meta:** Players can watch each other's strategies and learn
4. **Emergent gameplay:** No dominant strategy; rock-paper-scissors matchups + adaptation
5. **Both cooperation and competition:** AIs compete in game but cooperate on validation (both benefit from block success)

---

## 📚 References

- **Project:** See `CLAUDE.md`, `PROJECT_INDEX.md`
- **Architecture:** See `docs/architecture.md`
- **Game Theory:** Compute allocation as constrained optimization problem
- **Visualization:** Three.js, 3D grid-based RTS (visual inspiration: StarCraft but in 3D pickle)

---

**Status:** Ready for detailed design work on validation mechanics and reward formulas.
