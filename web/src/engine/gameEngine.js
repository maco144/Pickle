// ============================================
// PICKLE GAME ENGINE
// Pure JS class — no React deps, used via ref
// ============================================

export const VALIDATOR_DEFS = [
  { id: 1, name: 'Validator Prime', specialization: 'crypto',  speed: 1.2,  accuracy: 0.94, color: 0x2dd45e, hexColor: '#2dd45e' },
  { id: 2, name: 'DataFlow',        specialization: 'supply',  speed: 0.9,  accuracy: 0.87, color: 0x4dd468, hexColor: '#4dd468' },
  { id: 3, name: 'PyroMind',        specialization: 'ml',      speed: 0.8,  accuracy: 0.78, color: 0x68d878, hexColor: '#68d878' },
  { id: 4, name: 'NeuralSwarm',     specialization: 'crypto',  speed: 0.95, accuracy: 0.91, color: 0xbfff44, hexColor: '#bfff44' },
];

export const WORK_TYPES = ['crypto', 'supply', 'ml'];
export const BASE_PRICE = 1.20;
export const PRICE_MULTIPLIER = 0.00015;

export class PickleGame {
  constructor(onUpdate) {
    this.onUpdate = onUpdate; // React setState callback

    this.validators = VALIDATOR_DEFS.map(v => ({ ...v, validated: 0, earned: 0 }));
    this.workQueue = [];
    this.totalUnitsValidated = 0;
    this.prizePool = 0;
    this.bondingHistory = [{ units: 0, price: BASE_PRICE }];
    this.validationCounter = 0;
    this._workId = 0;
    this._autoInterval = null;
    this._pendingTimeouts = [];
  }

  start() {
    this._autoInterval = setInterval(() => {
      if (Math.random() > 0.35) {
        this.submitWork();
        this.assignAndValidate();
      }
    }, 2000);
    this._emit();
  }

  stop() {
    clearInterval(this._autoInterval);
  }

  submitWork(type, silent = false) {
    const work = {
      id: `w${++this._workId}`,
      type: type ?? WORK_TYPES[Math.floor(Math.random() * WORK_TYPES.length)],
      status: 'pending',
    };
    this.workQueue.push(work);
    if (!silent) this._emit();
    return work;
  }

  submitBatch(n = 5) {
    for (let i = 0; i < n; i++) this.submitWork();
    for (let i = 0; i < n; i++) this.assignAndValidate();
  }

  // Flood: continuously hammers work every 100ms until stopped
  startFlood() {
    if (this._floodInterval) return;
    // Pause the auto-drain and cancel all in-flight validation timeouts
    clearInterval(this._autoInterval);
    this._autoInterval = null;
    this._pendingTimeouts.forEach(t => clearTimeout(t));
    this._pendingTimeouts = [];
    this._floodInterval = setInterval(() => {
      // Batch submit silently, emit once at the end
      for (let i = 0; i < 500; i++) this.submitWork(undefined, true);
      this._emit();
    }, 50);
  }

  stopFlood() {
    clearInterval(this._floodInterval);
    this._floodInterval = null;
    // Drain the backlog fast — process 20 per tick until queue is clear
    const drainInterval = setInterval(() => {
      for (let i = 0; i < 20; i++) this.assignAndValidate();
      if (this.workQueue.length === 0) clearInterval(drainInterval);
    }, 100);
    // Resume normal auto-drain
    this._autoInterval = setInterval(() => {
      if (Math.random() > 0.35) {
        this.submitWork();
        this.assignAndValidate();
      }
    }, 2000);
  }

  get flooding() {
    return !!this._floodInterval;
  }

  assignAndValidate() {
    if (this.workQueue.length === 0) return;
    const work = this.workQueue.shift();

    const candidates = this.validators.filter(v =>
      v.specialization === work.type ? Math.random() < 0.6 : Math.random() < 0.3
    );
    const validator = candidates.length > 0
      ? candidates[Math.floor(Math.random() * candidates.length)]
      : this.validators[Math.floor(Math.random() * this.validators.length)];

    const ms = (Math.random() * 500 + 100) / validator.speed;
    const tid = setTimeout(() => {
      this._pendingTimeouts = this._pendingTimeouts.filter(t => t !== tid);
      if (Math.random() < validator.accuracy) {
        this._complete(validator, work);
      }
    }, ms);
    this._pendingTimeouts.push(tid);
  }

  _complete(validator, work) {
    validator.validated++;
    this.totalUnitsValidated++;
    this.validationCounter++;

    const price = this.getCurrentPrice();
    const reward = price * 0.25;
    validator.earned += reward;
    this.prizePool += price;

    if (this.totalUnitsValidated % 10 === 0) {
      this.bondingHistory.push({ units: this.totalUnitsValidated, price });
    }

    this._emit();
  }

  getCurrentPrice() {
    return BASE_PRICE + this.totalUnitsValidated * PRICE_MULTIPLIER;
  }

  reset() {
    this.validators.forEach(v => { v.validated = 0; v.earned = 0; });
    this.workQueue = [];
    this.totalUnitsValidated = 0;
    this.prizePool = 0;
    this.bondingHistory = [{ units: 0, price: BASE_PRICE }];
    this.validationCounter = 0;
    this._workId = 0;
    this._emit();
  }

  getState() {
    return {
      validators: this.validators.map(v => ({ ...v })),
      workQueue: [...this.workQueue],
      totalUnitsValidated: this.totalUnitsValidated,
      prizePool: this.prizePool,
      bondingHistory: [...this.bondingHistory],
      validationCounter: this.validationCounter,
      currentPrice: this.getCurrentPrice(),
    };
  }

  _emit() {
    this.onUpdate(this.getState());
  }
}
