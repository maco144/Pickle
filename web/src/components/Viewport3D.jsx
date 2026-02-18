import { useEffect, useRef } from 'react';
import { PickleScene } from '../scene/PickleScene';

export function Viewport3D() {
  const containerRef = useRef(null);
  const sceneRef = useRef(null);

  useEffect(() => {
    if (!containerRef.current) return;
    sceneRef.current = new PickleScene(containerRef.current);
    return () => sceneRef.current?.dispose();
  }, []);

  return (
    <div
      ref={containerRef}
      style={{
        gridColumn: 1,
        gridRow: 2,
        borderRadius: 8,
        overflow: 'hidden',
        position: 'relative',
        border: '1px solid rgba(45,212,94,0.15)',
        background: 'radial-gradient(ellipse at 50% 50%, rgba(45,212,94,0.08) 0%, transparent 70%), linear-gradient(to bottom, #0f1410, #1a2218)',
      }}
    />
  );
}
