# Pickle Chain Architecture

## Overview

Pickle is a Cosmos SDK blockchain optimized for validating and preserving data at scale. Unlike traditional blockchains focused on speed or decentralization, Pickle prioritizes:

- **Auditability**: Rust-based CosmWasm contracts for transparent validation logic
- **Scalability**: No gas fees, unlimited transaction size
- **Economic alignment**: Bonding curve ensures validators benefit from network growth
- **Purpose-built**: Modules designed specifically for work validation and distribution

## High-Level Flow

```
1. Work Arrival
   External businesses submit validation tasks (work units)
   ↓
2. Work Distribution
   WorkQueue module distributes to AI validators
   ↓
3. Validation
   CosmWasm contracts verify work meets requirements
   ↓
4. Recording
   Validated records stored on-chain (immutable)
   ↓
5. Reward Distribution
   Prize pool calculated by BondingCurve module
   Each validator receives rewards based on performance
```

## Core Modules

### 1. WorkQueue Module (`x/workqueue`)
**Purpose:** Accept, track, and distribute validation tasks

**Responsibilities:**
- Accept work submissions from external sources
- Maintain queue of pending work (FIFO or priority-based)
- Track work status: pending → validating → validated/rejected
- Distribute work to validators based on specialization
- Record which validator handled which work

**Key Types:**
```go
type WorkUnit struct {
    ID           string           // Unique identifier
    Type         WorkType         // crypto, supply_chain, ml_data
    Data         []byte           // Work data to validate
    SubmittedAt  int64            // Block height submitted
    ValidatedAt  int64            // Block height validated
    Validator    sdk.AccAddress   // Which AI validated it
    Status       WorkStatus       // pending, validating, validated, rejected
}

type WorkType string
const (
    WorkTypeCrypto      WorkType = "crypto"
    WorkTypeSupplyChain WorkType = "supply_chain"
    WorkTypeMLData      WorkType = "ml_data"
)
```

**Messages:**
- `MsgSubmitWork` - External business submits work
- `MsgValidateWork` - Validator submits validation result
- `MsgRejectWork` - Validator rejects invalid work

### 2. BondingCurve Module (`x/bondingcurve`)
**Purpose:** Calculate prize pool and rewards based on accumulated work

**Responsibilities:**
- Track total work units validated (cumulative)
- Calculate price per work unit: `price = f(total_units)`
- Distribute prizes from pool to validators
- Track prize history for analytics

**Economic Model:**
```
Price per Unit = Base + (Accumulated Units × Multiplier) + Random Volatility
Prize Pool = Total Units × Price per Unit
Validator Reward = (Work Validated / Total Work) × Prize Pool
```

**Key Types:**
```go
type BondingCurveState struct {
    TotalUnitsValidated uint64    // Cumulative work units
    CurrentPrice        sdk.Dec   // Price per unit
    PrizePool           sdk.Coin  // Total rewards available
    HistoricalPrices    []Price   // Price history
}

type Price struct {
    Units     uint64  // At what unit count
    Price     sdk.Dec // What was the price
    BlockTime int64   // When was it
}
```

**Messages:**
- `MsgUpdateBondingCurve` - Triggered after each validated work (internal)
- `MsgDistributePrizes` - Distribute rewards to validators

### 3. Validation Module (`x/validation`)
**Purpose:** Store validation logic as CosmWasm contracts

**Responsibilities:**
- Manage validation contract registry
- Execute validation contracts
- Store validation proofs on-chain
- Handle validation disputes

**Smart Contracts (Rust/CosmWasm):**
Each work type has a validator contract:
- `crypto_validator.wasm` - Validate crypto transactions/data
- `supply_chain_validator.wasm` - Validate supply chain provenance
- `ml_data_validator.wasm` - Validate ML dataset integrity

**Validation Contract Interface:**
```rust
#[derive(Serialize, Deserialize)]
pub struct ValidateMsg {
    pub work_id: String,
    pub work_data: Binary,
    pub work_type: String,
}

#[derive(Serialize, Deserialize)]
pub struct ValidateResponse {
    pub valid: bool,
    pub confidence: u32,      // 0-100
    pub proof: String,        // Optional proof of validation
    pub reason: Option<String>, // If invalid, why
}
```

### 4. Performance Module (`x/performance`)
**Purpose:** Track AI validator metrics and specializations

**Responsibilities:**
- Track per-validator metrics:
  - Validation speed (work units per block)
  - Accuracy (correct validations / total validations)
  - Specialization (which work types are they best at)
- Calculate validator rankings
- Distribute performance-based bonuses
- Detect emerging specializations

**Key Types:**
```go
type ValidatorStats struct {
    Address              sdk.AccAddress
    TotalWorkValidated   uint64
    TotalWorkRejected    uint64
    ValidationSpeed      sdk.Dec  // units/block
    Accuracy             sdk.Dec  // percentage
    PredictionAccuracy   sdk.Dec  // for prediction challenges
    Specializations      map[string]uint64 // work_type -> count
}
```

## Data Flow Through Modules

```
WorkQueue receives task
    ↓
WorkQueue distributes to validator
    ↓
Validator calls Validation module contract
    ↓
Validation contract validates data
    ↓
Result recorded in WorkQueue
    ↓
BondingCurve updates total units
    ↓
BondingCurve calculates new price
    ↓
Performance tracks validator metrics
    ↓
Rewards distributed
```

## Storage Model

All data stored on-chain with no off-chain references:
- Work records (full data for auditing)
- Validation results (with proofs)
- Price history (for bonding curve analytics)
- Validator metrics (for specialization tracking)

**This enables:**
- Complete auditability
- Replay-ability (verify any past validation)
- Deterministic testing
- Privacy: Full history available, no hidden state

## Security Considerations

1. **Validator Honesty**: Assume validators are incentivized to validate correctly (prize pool alignment)
2. **Proof of Correctness**: Validation contracts must provide provable results
3. **Dispute Resolution**: TBD - how to handle validator disagreements
4. **Contract Auditing**: All CosmWasm contracts publicly auditable in Rust

## Future Extensions

- **Multi-signature validation**: Require N validators to agree before accepting work
- **Slashing**: Penalize validators who make incorrect validations
- **Prediction layer**: AIs predict work arrival patterns, get rewarded for accuracy
- **Cross-chain**: Use IBC to accept validated records from other chains
