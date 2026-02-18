# 🥒 Pickle Game Mechanics Guide

## Overview

Pickle is a **live, interactive game** where AI validators compete to validate incoming work (data preservation tasks) in exchange for rewards. The game demonstrates Pickle's core economic model in real-time.

## Starting the Game

### Option 1: Open Dashboard Directly
```bash
# From the Pickle directory
open dashboard/forgeground-dashboard.html
# Or in any web browser, navigate to the file
```

### Option 2: Serve via Local Server
```bash
# Run a simple HTTP server in the dashboard folder
cd dashboard
python3 -m http.server 8000
# Then open: http://localhost:8000/forgeground-dashboard.html
```

## Game Screen Overview

The dashboard is split into 4 main areas:

### 1. **Header (Top)**
Real-time stats about the current game state:
- **Work Queue**: Number of pending work units waiting for validation
- **Prize Pool**: Total $USD accumulated from completed validations
- **Validators Active**: Number of AI validators running (currently 4)

### 2. **Main Viewport (Center-Left)**
3D visualization showing:
- **Pickle Grid**: The foundational state (visual metaphor)
- **4 Validator Nodes**: Spinning, glowing spheres representing each AI
  - Validator Prime (bright green)
  - DataFlow (medium green)
  - PyroMind (dark green)
  - NeuralSwarm (yellow)
- **Particle System**: Shows active work flowing through the system
  - Green particles = Crypto validation work
  - Light green particles = Supply chain validation work
  - Yellow particles = ML data validation work
- **Neural Connections**: Lines between validators showing cooperation/competition

### 3. **Leaderboard (Top-Right)**
Live rankings of validators:
- **#1, #2, #3**: Highlighted with medals (gold, silver, bronze)
- **Per validator shown**:
  - Name
  - Total validated work units
  - Prediction accuracy percentage
  - Total earned reward ($USD)
- **Automatically sorted** by earned rewards

### 4. **Work Distribution & Stats (Bottom-Right)**
Shows queue breakdown:
- **Crypto Validation**: # of pending crypto work units
- **Supply Chain**: # of pending supply chain work units
- **ML Data**: # of pending ML dataset work units

### 5. **Bottom Panel**
Three sections:

#### Bonding Curve Chart
- X-axis: Total work units accumulated
- Y-axis: Price per unit (USD)
- Shows the **dynamic pricing model**: as more work is validated, price increases
- Economy grows organically as network activity increases

#### Work Statistics
- **Total Units**: Cumulative validated work since game start
- **Price/Unit**: Current market price for validating one unit of work

#### Controls & Progress
- **+ Add Work**: Manually submit 5 random work units
- **Reset**: Clear all state and restart the game
- **Validation Rate**: Live progress bar showing validation throughput

## How the Game Works

### 1. Work Submission
Work units are **automatically submitted** every 2 seconds (plus manual submissions with "+ Add Work"):
- Work type is randomly chosen (crypto, supply, ML)
- Work enters the queue as "pending"

### 2. Validator Assignment
Each validator is specialized in certain work types:
- **Validator Prime**: Crypto specialist (60% match rate, 30% general)
- **DataFlow**: Supply chain specialist
- **PyroMind**: ML data specialist
- **NeuralSwarm**: Crypto backup specialist

When work is submitted, a random validator is assigned based on:
- **Specialization match**: 60% chance if they specialize in that type
- **General capability**: 30% chance they'll take any work type
- **Accuracy factor**: Some validations may fail (based on validator accuracy)

### 3. Validation Execution
Once assigned:
- Validator processes the work (takes 100-500ms depending on speed)
- Validator's glow sphere intensifies (visual "thinking" indicator)
- Work particles flow toward the validator node
- Validation succeeds based on validator's accuracy rating

### 4. Reward Distribution
On successful validation:
- **Reward per unit** = Base Price + (Total Units × Multiplier)
- **Base Price**: $1.20
- **Multiplier**: $0.00015 per accumulated unit
- **Validator earns**: 25% of the price per validated unit (simplified model)
- **Prize pool grows**: Each validation adds price amount to total pool

### 5. Leaderboard Updates
After each validation:
- Validator's "validated" count increments
- Validator's "earned" total updates
- Leaderboard re-sorts automatically
- Neural connections pulse to show activity

## Economic Model (Simplified)

### Bonding Curve Formula
```
Price Per Unit = Base + (Total_Units_Validated × Multiplier) + Volatility

Base = $1.20
Multiplier = 0.00015
Volatility = Random noise (currently not implemented but can be)
```

### Example Progression
| Total Units | Price/Unit | Cumulative Pool | Per Validator |
|-------------|-----------|-----------------|---------------|
| 0           | $1.20     | $0.00           | $0.30         |
| 100         | $1.215    | $121.50         | $0.30         |
| 500         | $1.275    | $637.50         | $0.32         |
| 1000        | $1.35     | $1,350.00       | $0.34         |

## Game Mechanics Deep Dive

### Validation Specialization
Each AI has two metrics that define their profile:

**Validation Speed** (multiplier on processing time):
- Validator Prime: 1.2x (fastest - 83-416ms)
- NeuralSwarm: 0.95x (fast - 105-526ms)
- DataFlow: 0.9x (normal - 111-555ms)
- PyroMind: 0.8x (slower - 125-625ms)

**Prediction Accuracy** (confidence in correct validation):
- Validator Prime: 94% (most accurate)
- NeuralSwarm: 91%
- DataFlow: 87%
- PyroMind: 78% (less accurate)

### Competition Strategy
Validators naturally compete through:
1. **Speed race**: Faster validators process more work per unit time
2. **Accuracy advantage**: Higher accuracy = more validated work = more rewards
3. **Specialization bonus**: 60% match rate for specialized types
4. **Work queue dynamics**: When queue builds up, specialists shine

### Cooperation Mechanics (Visual)
Neural connections between validators show relationship types:
- **Green lines**: Cooperation (validators helping each other)
- **Red lines**: Competition (direct rivalry)
- **Purple lines**: Neutral relationships

(Currently cosmetic but can implement actual bonus/penalty mechanics)

## Interactive Controls

### Manual Work Submission
```
Click "+ Add Work" button
→ Submits 5 random work units to the queue
→ All automatically assigned and processed
```

### Reset Game
```
Click "Reset" button
→ Clear all validators' earned rewards
→ Clear all work from queue
→ Reset leaderboard
→ Reset bonding curve chart
→ Start fresh from $0.00 prize pool
```

## Real-World Applications

This game simulates Pickle's actual use cases:

### 1. Supply Chain Validation
- Companies submit supply chain records to validate
- AIs verify provenance, authenticity, timestamps
- Validated records preserved immutably on-chain

### 2. Cryptocurrency Validation
- Validate transaction signatures, amounts, and sequences
- Detect fraud or anomalies
- Preserve trusted transaction history

### 3. ML Dataset Validation
- Validate ML training datasets for quality/integrity
- Check for mislabeled data, duplicates, corruption
- Preserve clean versions for future model training

## Advanced Features (Future)

### Planned Enhancements
1. **Multi-signature validation**: Require 2-3 validators to agree
2. **Slashing penalties**: Validators lose earned rewards for incorrect validations
3. **Prediction layer**: AIs predict work arrival patterns, get bonus rewards
4. **Specialization tracking**: Track emerging specializations over time
5. **Dispute resolution**: Community votes on conflicting validations
6. **Cross-chain support**: Validate work from other blockchains via IBC

### Difficulty Scaling
- **Hard mode**: Validators only earn for accuracy
- **Survival mode**: Limited work available (scarcity competition)
- **Chaos mode**: Random validator crashes/slowdowns

## Metrics & Analysis

### Key Performance Indicators (KPIs)
- **Total throughput**: Work units/second
- **Average accuracy**: Overall network accuracy
- **Price trend**: Is bonding curve growing consistently?
- **Validator participation**: Is any validator dominating?
- **Queue depth**: Is work backing up?

### Optimization Strategies
Watch for:
- **Specialist advantage**: Do specialists earn more?
- **Speed vs accuracy tradeoff**: Do fast validators sacrifice accuracy?
- **Market efficiency**: Does price follow fundamental value?

## Tips for Maximum Engagement

1. **Watch the leaderboard**: Notice who's earning the most
2. **Spam work submissions**: Click "+ Add Work" repeatedly for chaos
3. **Look for patterns**: Notice how validator behavior changes with work type
4. **Check the bonding curve**: See exponential growth as system scales
5. **Monitor neural connections**: Watch the validation network light up

## Troubleshooting

### Dashboard Not Loading
- Check browser console (F12) for JavaScript errors
- Ensure you're serving via HTTP/file, not opening raw file in some browsers
- Clear browser cache

### Chart Not Updating
- Chart updates every 10 validated units
- If no chart: Submit work manually with "+ Add Work"
- Zoom in on bonding curve to see small price changes

### Validators Not Validating
- Work gets auto-assigned every 2 seconds
- Check if queue has work (should see numbers in bottom-right)
- Accuracy: Some validations intentionally "fail" (< accuracy %)

## Game Philosophy

Pickle's game is designed to:
- **Educate**: Show how bonding curves create sustainable incentives
- **Visualize**: Make abstract blockchain concepts tangible
- **Engage**: Create competitive gameplay that's inherently interesting
- **Scale**: Demonstrate exponential growth through economic models
- **Preserve**: Show the core mission: immutable data preservation

The goal is **not** to win/lose, but to understand how AI validators, economic incentives, and decentralized networks can work together to preserve data at scale.

---

**Start the game now**: Open `dashboard/forgeground-dashboard.html` in your browser! 🥒

