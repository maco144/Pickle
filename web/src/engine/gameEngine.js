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

  submitWork(type) {
    const work = {
      id: `w${++this._workId}`,
      type: type ?? WORK_TYPES[Math.floor(Math.random() * WORK_TYPES.length)],
      status: 'pending',
    };
    this.workQueue.push(work);
    this._emit();
    return work;
  }

  submitBatch(n = 5) {
    for (let i = 0; i < n; i++) this.submitWork();
    for (let i = 0; i < n; i++) this.assignAndValidate();
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
    setTimeout(() => {
      if (Math.random() < validator.accuracy) {
        this._complete(validator, work);
      }
    }, ms);
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
