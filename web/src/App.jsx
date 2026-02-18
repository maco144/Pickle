import { useEffect, useRef, useState } from 'react';
import { PickleGame } from './engine/gameEngine';
import { Viewport3D } from './components/Viewport3D';
import { Leaderboard } from './components/Leaderboard';
import { WorkDistribution } from './components/WorkDistribution';
import { BondingCurve } from './components/BondingCurve';
import { BottomControls } from './components/BottomControls';

const INITIAL_STATE = {
  validators: [],
  workQueue: [],
  totalUnitsValidated: 0,
  prizePool: 0,
  bondingHistory: [{ units: 0, price: 1.2 }],
  validationCounter: 0,
  currentPrice: 1.2,
};

export default function App() {
  const [gameState, setGameState] = useState(INITIAL_STATE);
  const gameRef = useRef(null);

  useEffect(() => {
    const game = new PickleGame(setGameState);
    gameRef.current = game;
    game.start();
    return () => game.stop();
  }, []);

  const [flooding, setFlooding] = useState(false);

  const handleAddWork = () => gameRef.current?.submitBatch(5);
  const handleReset   = () => {
    gameRef.current?.stopFlood();
    setFlooding(false);
    gameRef.current?.reset();
  };
  const handleFlood = () => {
    const game = gameRef.current;
    if (!game) return;
    if (flooding) {
      game.stopFlood();
      setFlooding(false);
    } else {
      game.startFlood();
      setFlooding(true);
    }
  };

  const {
    validators, workQueue, totalUnitsValidated,
    prizePool, bondingHistory, validationCounter, currentPrice,
  } = gameState;

  return (
    <div style={dashboardStyle}>
      {/* ── Header ── */}
      <header style={headerStyle}>
        <div style={logoStyle}>🥒 PICKLE · LIVE GAME</div>
        <div style={{ display: 'flex', gap: 32 }}>
          <Stat label="Work Queue"        value={`${workQueue.length} pending`} />
          <Stat label="Prize Pool"        value={`$${prizePool.toFixed(2)}`} />
          <Stat label="Validators Active" value={`${validators.length} AIs`} />
          <Stat label="Price / Unit"      value={`$${currentPrice.toFixed(3)}`} />
        </div>
      </header>

      {/* ── 3D Viewport ── */}
      <Viewport3D />

      {/* ── Right Sidebar ── */}
      <div style={sidebarStyle}>
        <Leaderboard validators={validators} />
        <WorkDistribution workQueue={workQueue} />
      </div>

      {/* ── Bottom row ── */}
      <div style={bottomStyle}>
        <BondingCurve bondingHistory={bondingHistory} currentPrice={currentPrice} />
        <BottomControls
          totalUnitsValidated={totalUnitsValidated}
          currentPrice={currentPrice}
          validationCounter={validationCounter}
          workQueueSize={workQueue.length}
          onAddWork={handleAddWork}
          onReset={handleReset}
          onFlood={handleFlood}
          flooding={flooding}
        />
      </div>
    </div>
  );
}

function Stat({ label, value }) {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
      <span style={{ color: '#888', textTransform: 'uppercase', letterSpacing: '0.5px', fontSize: 10 }}>
        {label}
      </span>
      <span style={{ fontWeight: 600, color: '#2dd45e', fontSize: 13 }}>{value}</span>
    </div>
  );
}

// ─── Layout styles ────────────────────────────────────────────────────────────

const dashboardStyle = {
  display: 'grid',
  gridTemplateColumns: '1fr 360px',
  gridTemplateRows: '60px 1fr 210px',
  height: '100vh',
  gap: 8,
  padding: 8,
  background: 'linear-gradient(135deg, #0f1410 0%, #1a2218 100%)',
  boxSizing: 'border-box',
};

const headerStyle = {
  gridColumn: '1 / -1',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  padding: '0 24px',
  background: 'rgba(20,30,22,0.8)',
  backdropFilter: 'blur(10px)',
  borderBottom: '1px solid rgba(100,200,100,0.2)',
  borderRadius: 8,
};

const logoStyle = {
  fontSize: 18,
  fontWeight: 700,
  letterSpacing: '-0.5px',
  background: 'linear-gradient(135deg, #2dd45e, #68d878)',
  WebkitBackgroundClip: 'text',
  WebkitTextFillColor: 'transparent',
  backgroundClip: 'text',
};

const sidebarStyle = {
  gridColumn: 2,
  gridRow: 2,
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
  minHeight: 0,
};

const bottomStyle = {
  gridColumn: '1 / -1',
  display: 'grid',
  gridTemplateColumns: '2fr 1fr',
  gap: 8,
};
