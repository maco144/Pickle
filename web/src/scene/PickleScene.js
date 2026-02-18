// ============================================
// THREE.JS SCENE — imperative, lives outside React
// ============================================
import * as THREE from 'three';
import { VALIDATOR_DEFS } from '../engine/gameEngine';

export class PickleScene {
  constructor(container) {
    this.container = container;
    this._nodes = [];
    this._connections = [];
    this._flowPositions = [];
    this._particlePositions = null;
    this._particleColors = null;
    this._particleCount = 0;
    this._maxParticles = 2000;
    this._time = 0;
    this._raf = null;
    this._init();
  }

  _init() {
    const w = this.container.clientWidth;
    const h = this.container.clientHeight;

    // Scene
    this.scene = new THREE.Scene();
    this.scene.fog = new THREE.Fog(0x0f1410, 100, 500);

    // Camera
    this.camera = new THREE.PerspectiveCamera(75, w / h, 0.1, 1000);
    this.camera.position.set(0, 40, 40);
    this.camera.lookAt(0, 0, 0);

    // Renderer
    this.renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
    this.renderer.setSize(w, h);
    this.renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    this.renderer.setClearColor(0x0f1410, 0.1);
    this.renderer.shadowMap.enabled = true;
    this.container.appendChild(this.renderer.domElement);

    // Lights
    this.scene.add(new THREE.AmbientLight(0xffffff, 0.4));
    const pt = new THREE.PointLight(0x2dd45e, 1, 200);
    pt.position.set(50, 50, 50);
    pt.castShadow = true;
    this.scene.add(pt);

    // Grid
    const grid = new THREE.GridHelper(100, 20, 0x2dd45e, 0x1a3d2a);
    grid.position.y = -0.1;
    this.scene.add(grid);

    this._buildParticles();
    this._buildNodes();
    this._buildConnections();
    this._buildFlowParticles();

    // Resize
    this._onResize = () => {
      const w2 = this.container.clientWidth;
      const h2 = this.container.clientHeight;
      this.camera.aspect = w2 / h2;
      this.camera.updateProjectionMatrix();
      this.renderer.setSize(w2, h2);
    };
    window.addEventListener('resize', this._onResize);

    this._animate();
  }

  _buildParticles() {
    const geo = new THREE.BufferGeometry();
    this._particlePositions = new Float32Array(this._maxParticles * 3);
    this._particleColors = new Float32Array(this._maxParticles * 3);
    geo.setAttribute('position', new THREE.BufferAttribute(this._particlePositions, 3));
    geo.setAttribute('color', new THREE.BufferAttribute(this._particleColors, 3));
    geo.setDrawRange(0, 0);
    const mat = new THREE.PointsMaterial({ size: 0.6, vertexColors: true, transparent: true, opacity: 0.8 });
    this._particles = new THREE.Points(geo, mat);
    this.scene.add(this._particles);
    this._particlesGeo = geo;
  }

  _buildNodes() {
    this._nodes = VALIDATOR_DEFS.map((v, i) => {
      const angle = (i / VALIDATOR_DEFS.length) * Math.PI * 2;
      const x = Math.cos(angle) * 30;
      const z = Math.sin(angle) * 30;

      const geo = new THREE.IcosahedronGeometry(2.5, 4);
      const mat = new THREE.MeshStandardMaterial({
        color: v.color, emissive: v.color, emissiveIntensity: 0.5,
        metalness: 0.8, roughness: 0.2,
      });
      const mesh = new THREE.Mesh(geo, mat);
      mesh.position.set(x, 15, z);
      this.scene.add(mesh);

      const glowGeo = new THREE.IcosahedronGeometry(4, 4);
      const glowMat = new THREE.MeshBasicMaterial({ color: v.color, transparent: true, opacity: 0.1 });
      const glow = new THREE.Mesh(glowGeo, glowMat);
      glow.position.set(x, 15, z);
      this.scene.add(glow);

      return { mesh, glow, position: { x, y: 15, z }, validatorId: v.id };
    });
  }

  _buildConnections() {
    for (let i = 0; i < this._nodes.length; i++) {
      for (let j = i + 1; j < this._nodes.length; j++) {
        const pts = [
          new THREE.Vector3(this._nodes[i].position.x, this._nodes[i].position.y, this._nodes[i].position.z),
          new THREE.Vector3(this._nodes[j].position.x, this._nodes[j].position.y, this._nodes[j].position.z),
        ];
        const geo = new THREE.BufferGeometry().setFromPoints(pts);
        const mat = new THREE.LineBasicMaterial({ color: 0x2dd45e, transparent: true, opacity: 0.15 });
        const line = new THREE.Line(geo, mat);
        this.scene.add(line);
        this._connections.push({ line, from: i, to: j, intensity: 0 });
      }
    }
  }

  _buildFlowParticles() {
    const n = this._connections.length;
    const geo = new THREE.BufferGeometry();
    const pos = new Float32Array(n * 3);
    const col = new Float32Array(n * 3);
    for (let i = 0; i < n; i++) { col[i * 3] = 0.2; col[i * 3 + 1] = 0.8; col[i * 3 + 2] = 0.3; }
    geo.setAttribute('position', new THREE.BufferAttribute(pos, 3));
    geo.setAttribute('color', new THREE.BufferAttribute(col, 3));
    const mat = new THREE.PointsMaterial({ size: 1.0, vertexColors: true, transparent: true, opacity: 0.9 });
    this._flowPoints = new THREE.Points(geo, mat);
    this.scene.add(this._flowPoints);
    this._flowGeo = geo;
    this._flowProgress = new Float32Array(n).fill(0);
  }

  _spawnParticle(node) {
    if (this._particleCount >= this._maxParticles) return;
    const idx = this._particleCount++;
    const angle = Math.random() * Math.PI * 2;
    const r = Math.random() * 18 + 4;
    this._particlePositions[idx * 3]     = node.position.x + Math.cos(angle) * r;
    this._particlePositions[idx * 3 + 1] = 15 + (Math.random() - 0.5) * 10;
    this._particlePositions[idx * 3 + 2] = node.position.z + Math.sin(angle) * r;
    const colors = [[0.2, 0.8, 0.3], [0.3, 0.9, 0.2], [0.8, 0.8, 0.1]];
    const c = colors[Math.floor(Math.random() * 3)];
    this._particleColors[idx * 3] = c[0];
    this._particleColors[idx * 3 + 1] = c[1];
    this._particleColors[idx * 3 + 2] = c[2];
  }

  _animate() {
    this._raf = requestAnimationFrame(() => this._animate());
    this._time += 0.016;
    const t = this._time;

    // Spawn particles
    if (this._particleCount < this._maxParticles * 0.8) {
      for (const node of this._nodes) {
        if (Math.random() < 0.25) this._spawnParticle(node);
      }
    }

    // Move particles
    for (let i = 0; i < this._particleCount; i++) {
      this._particlePositions[i * 3 + 1] += Math.sin(t + i * 0.1) * 0.07;
      if (this._particlePositions[i * 3 + 1] > 55 || this._particlePositions[i * 3 + 1] < -10) {
        this._particlePositions[i * 3 + 1] = (Math.random() - 0.5) * 20 + 15;
      }
    }
    this._particlesGeo.attributes.position.needsUpdate = true;
    this._particlesGeo.setDrawRange(0, this._particleCount);

    // Animate validator nodes
    this._nodes.forEach((node, i) => {
      node.mesh.rotation.x += 0.003;
      node.mesh.rotation.y += 0.005;
      node.mesh.position.y = 15 + Math.sin(t * 0.5 + i) * 2;
      const pulse = Math.sin(t + i * 1.5) * 0.5 + 0.5;
      node.glow.material.opacity = 0.06 + pulse * 0.14;
    });

    // Animate connection flows
    const flowPos = this._flowGeo.attributes.position.array;
    this._connections.forEach((conn, idx) => {
      this._flowProgress[idx] = (this._flowProgress[idx] + 0.04) % 1;
      const p = this._flowProgress[idx];
      const from = this._nodes[conn.from].position;
      const to   = this._nodes[conn.to].position;
      flowPos[idx * 3]     = from.x + (to.x - from.x) * p;
      flowPos[idx * 3 + 1] = from.y + (to.y - from.y) * p;
      flowPos[idx * 3 + 2] = from.z + (to.z - from.z) * p;

      conn.intensity = Math.min(0.45, conn.intensity + Math.random() * 0.1);
      conn.intensity = Math.max(0.05, conn.intensity - 0.015);
      conn.line.material.opacity = conn.intensity;
    });
    this._flowGeo.attributes.position.needsUpdate = true;

    this.renderer.render(this.scene, this.camera);
  }

  dispose() {
    cancelAnimationFrame(this._raf);
    window.removeEventListener('resize', this._onResize);
    this.renderer.dispose();
    if (this.renderer.domElement.parentNode) {
      this.renderer.domElement.parentNode.removeChild(this.renderer.domElement);
    }
  }
}
