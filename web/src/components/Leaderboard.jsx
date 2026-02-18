const MEDAL = ['#ffd700', '#c0c0c0', '#cd7f32'];

export function Leaderboard({ validators }) {
  const sorted = [...validators].sort((a, b) => b.earned - a.earned);

  return (
    <div style={panelStyle}>
      <div style={titleStyle}>🏆 Leaderboard</div>
      <div style={{ display: 'flex', flexDirection: 'column', gap: 6, overflowY: 'auto' }}>
        {sorted.map((v, idx) => {
          const medalColor = MEDAL[idx] ?? '#2dd45e';
          return (
            <div
              key={v.id}
              style={{
                padding: '10px 10px',
                background: idx < 3 ? `${medalColor}11` : 'rgba(15,30,20,0.5)',
                borderLeft: `3px solid ${medalColor}`,
                borderRadius: 4,
                display: 'flex',
                alignItems: 'center',
                gap: 10,
                fontSize: 11,
              }}
            >
              <span style={{ fontWeight: 700, color: medalColor, minWidth: 22 }}>#{idx + 1}</span>
              <div style={{ flex: 1 }}>
                <div style={{ fontWeight: 600, color: '#fff' }}>{v.name}</div>
                <div style={{ color: '#888', fontSize: 10 }}>
                  {v.validated} validated · {(v.accuracy * 100).toFixed(0)}% acc · <span style={{ color: v.hexColor }}>{v.specialization}</span>
                </div>
              </div>
              <div style={{ fontWeight: 700, color: '#2dd45e', textAlign: 'right' }}>
                ${v.earned.toFixed(2)}
              </div>
            </div>
          );
        })}
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
  flex: 1,
  minHeight: 0,
};

const titleStyle = {
  fontSize: 12,
  fontWeight: 600,
  textTransform: 'uppercase',
  letterSpacing: '0.5px',
  color: '#2dd45e',
  marginBottom: 10,
  paddingBottom: 8,
  borderBottom: '1px solid rgba(45,212,94,0.1)',
};
