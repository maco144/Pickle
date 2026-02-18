import { useEffect, useRef } from 'react';
import { Chart } from 'chart.js/auto';

export function BondingCurve({ bondingHistory, currentPrice, prizePool, totalUnitsValidated }) {
  const canvasRef = useRef(null);
  const chartRef = useRef(null);

  useEffect(() => {
    if (!canvasRef.current) return;
    chartRef.current = new Chart(canvasRef.current, {
      type: 'line',
      data: { labels: ['Start'], datasets: [chartDataset([1.2])] },
      options: chartOptions(),
    });
    return () => chartRef.current?.destroy();
  }, []);

  useEffect(() => {
    if (!chartRef.current || bondingHistory.length === 0) return;
    chartRef.current.data.labels = bondingHistory.map(p => `${p.units}u`);
    chartRef.current.data.datasets[0].data = bondingHistory.map(p => p.price);
    chartRef.current.update('none');
  }, [bondingHistory]);

  return (
    <div style={panelStyle}>
      <div style={titleStyle}>📈 Bonding Curve</div>
      <canvas ref={canvasRef} style={{ flex: 1, minHeight: 0 }} />
    </div>
  );
}

function chartDataset(data) {
  return {
    label: 'Price Per Unit',
    data,
    borderColor: '#2dd45e',
    backgroundColor: 'rgba(45,212,94,0.08)',
    borderWidth: 2,
    fill: true,
    tension: 0.4,
    pointBackgroundColor: '#2dd45e',
    pointBorderColor: '#fff',
    pointRadius: 3,
    pointHoverRadius: 5,
  };
}

function chartOptions() {
  return {
    responsive: true,
    maintainAspectRatio: false,
    animation: false,
    plugins: { legend: { display: false } },
    scales: {
      y: {
        beginAtZero: true,
        min: 1,
        ticks: { color: '#888', font: { size: 10 } },
        grid: { color: 'rgba(45,212,94,0.08)' },
      },
      x: {
        ticks: { color: '#888', font: { size: 10 } },
        grid: { display: false },
      },
    },
  };
}

const panelStyle = {
  background: 'rgba(20,30,22,0.6)',
  backdropFilter: 'blur(10px)',
  border: '1px solid rgba(45,212,94,0.15)',
  borderRadius: 8,
  padding: 16,
  display: 'flex',
  flexDirection: 'column',
  overflow: 'hidden',
};

const titleStyle = {
  fontSize: 12,
  fontWeight: 600,
  textTransform: 'uppercase',
  letterSpacing: '0.5px',
  color: '#2dd45e',
  marginBottom: 8,
};
