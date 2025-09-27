[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/a1mora2788/overlay-consensus-sim)
# EACPO — Energy-Aware Blockchain Website

<img width="758" height="860" alt="image" src="https://github.com/user-attachments/assets/11022d01-fb42-485c-84bc-285663dbe626" />


This is the landing page for **EACPO overlay consensus**, showcasing research results and promoting energy-aware blockchain design.

## Features
- ⚡ Overlay consensus: ~66% less energy than Proof-of-Work while preserving security.
- 📊 Live simulation graphs (overlay vs PoW energy ratio and cumulative energy).
- 🌱 Built with React, Vite, Tailwind CSS, and Framer Motion.
- 🔒 Icons with graceful fallbacks for resilience.

## Project Structure
```
├── index.html
├── package.json
├── postcss.config.js
├── tailwind.config.js
├── vite.config.js
├── public/
│   └── media/
│       ├── overlay_vs_pow.png
│       ├── cumulative_energy.png
│       ├── live_graph.gif
│       └── live_graph.mp4
├── screenshot.png
└── src/
    ├── index.css
    ├── main.jsx
    └── Page.jsx
```

## Getting Started

1. **Install dependencies**
   ```bash
   npm install
   ```

2. **Run the dev server**
   ```bash
   npm run dev
   ```
   Open [http://localhost:5173](http://localhost:5173) in your browser.

3. **Build for production**
   ```bash
   npm run build
   npm run preview
   ```

## Adding Media
- Place your generated graphs and animations in `public/media/`.
- Default file names expected:
  - `overlay_vs_pow.png`
  - `cumulative_energy.png`
  - `live_graph.gif`
  - `live_graph.mp4` (optional, generated with ffmpeg)

If files are missing, the site gracefully falls back to text placeholders or the GIF instead of MP4.

## Deployment
- Push this repo to GitHub.
- Deploy directly on [Vercel](https://vercel.com) or [Netlify](https://www.netlify.com). Both auto-detect Vite projects.
- Your public site will be live within minutes.

## License
Open-source, MIT/Apache friendly. Contributions welcome!

---


Built with ❤️ to make blockchain greener and smarter.

