const WORK_COLORS = { crypto: '#2dd45e', supply: '#68d878', ml: '#bfff44' };
const WORK_LABELS = { crypto: 'Crypto Validation', supply: 'Supply Chain', ml: 'ML Data' };

export function WorkDistribution({ workQueue }) {
  const counts = { crypto: 0, supply: 0, ml: 0 };
  workQueue.forEach(w => { if (counts[w.type] !== undefined) counts[w.type]++; });

  return (
    <div style={panelStyle}>
      <div style={titleStyle}>⚡ Work Queue</div>
      {Object.entries(counts).map(([type, count]) => (
        <div key={type} style={{ display: 'flex', alignItems: 'center', gap: 10, fontSize: 12, margin: '4px 0' }}>
          <div style={{
            width: 8, height: 8, borderRadius: '50%',
            background: WORK_COLORS[type],
            boxShadow: `0 0 6px ${WORK_COLORS[type]}`,
            animation: 'pulse 2s infinite',
          }} />
          <span style={{ flex: 1, color: '#ccc' }}>{WORK_LABELS[type]}</span>
          <span style={{ fontWeight: 700, color: WORK_COLORS[type] }}>{count}</span>
        </div>
      ))}
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
  gap: 2,
};

const titleStyle = {
  fontSize: 12,
  fontWeight: 600,
  textTransform: 'uppercase',
  letterSpacing: '0.5px',
  color: '#2dd45e',
  marginBottom: 8,
  paddingBottom: 8,
  borderBottom: '1px solid rgba(45,212,94,0.1)',
};
