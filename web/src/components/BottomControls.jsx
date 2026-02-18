export function BottomControls({ totalUnitsValidated, currentPrice, validationCounter, workQueueSize, onAddWork, onReset, onFlood, flooding }) {
  const rate = (validationCounter % 100);

  return (
    <div style={panelStyle}>
      <div style={{ display: 'flex', gap: 8, marginBottom: 12, flexWrap: 'wrap' }}>
        <button style={btnStyle} onClick={onAddWork}>+ Add Work</button>
        <button
          onClick={onFlood}
          style={{
            ...btnStyle,
            background: flooding ? 'rgba(255,60,60,0.25)' : 'rgba(45,212,94,0.2)',
            border: flooding ? '1px solid rgba(255,60,60,0.8)' : '1px solid rgba(45,212,94,0.5)',
            color: flooding ? '#ff4444' : '#2dd45e',
            animation: flooding ? 'pulse 1s infinite' : 'none',
          }}
        >
          {flooding ? '⚡ FLOODING…' : '🌊 Flood'}
        </button>
        <button style={{ ...btnStyle, opacity: 0.7 }} onClick={onReset}>Reset</button>
      </div>
      <div style={{ display: 'flex', gap: 20, marginBottom: 10 }}>
        <div>
          <div style={labelStyle}>Total Validated</div>
          <div style={valueStyle}>{totalUnitsValidated.toLocaleString()}</div>
        </div>
        <div>
          <div style={labelStyle}>Price / Unit</div>
          <div style={valueStyle}>${currentPrice.toFixed(3)}</div>
        </div>
        <div>
          <div style={labelStyle}>Queue Depth</div>
          <div style={{
            ...valueStyle,
            color: workQueueSize > 50 ? '#ff4444' : workQueueSize > 20 ? '#ffaa00' : '#2dd45e',
          }}>
            {workQueueSize.toLocaleString()}
          </div>
        </div>
      </div>
      <div>
        <div style={labelStyle}>Validation Rate</div>
        <div style={{ height: 4, background: 'rgba(45,212,94,0.1)', borderRadius: 2, overflow: 'hidden', marginTop: 4 }}>
          <div style={{
            height: '100%',
            width: `${rate}%`,
            background: 'linear-gradient(90deg, #2dd45e, #68d878)',
            borderRadius: 2,
            transition: 'width 0.3s ease',
          }} />
        </div>
      </div>
    </div>
  );
}

const panelStyle = {
  background: 'rgba(20,30,22,0.6)',
  backdropFilter: 'blur(10px)',
  border: '1px solid rgba(45,212,94,0.15)',
  borderRadius: 8,
  padding: 16,
  display: 'flex',
  flexDirection: 'column',
};

const btnStyle = {
  background: 'rgba(45,212,94,0.2)',
  border: '1px solid rgba(45,212,94,0.5)',
  color: '#2dd45e',
  padding: '8px 14px',
  borderRadius: 4,
  cursor: 'pointer',
  fontSize: 12,
  fontWeight: 600,
  transition: 'background 0.2s',
};

const labelStyle = {
  color: '#888',
  textTransform: 'uppercase',
  letterSpacing: '0.5px',
  fontSize: 10,
};

const valueStyle = {
  fontWeight: 700,
  color: '#2dd45e',
  fontSize: 18,
};
