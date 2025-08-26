# Roguelike Chess Game - Project Context

## Project Overview
Building a roguelike chess game where players start with a standard chess board, earn money by capturing pieces, and can purchase upgrades/special pieces between matches. The AI opponents increase in difficulty as players progress through stages.

## Core Game Loop
1. Player starts with standard chess board vs AI with standard pieces
2. Player earns "money" for captures and other actions during gameplay
3. Between matches, player visits a shop to buy special pieces or upgrades
4. Player faces increasingly difficult AI opponents
5. Game continues until player loses (V1 - endless mode)

## Technical Stack

### Frontend
- **Framework**: Next.js with React + TypeScript
- **Animation**: Framer Motion (handles piece dragging, placement, and UI animations)
- **Styling**: CSS/Tailwind for UI, Canvas/SVG for chess board rendering
- **Deployment**: Vercel (automatic deployment from GitHub)

### Backend
- **Language**: Go 
- **Framework**: Gin or Echo
- **Chess Engine**: Stockfish via UCI protocol with goroutines for parallel processing
- **Database**: PostgreSQL (likely via Supabase for auth + free tier)
- **Deployment**: Railway or Render (auto-deploy from GitHub)

### Repository Structure
- **Monorepo**: Managed with Turbo
- **Apps**: `frontend/` (Next.js) and `backend/` (Go API)
- **Packages**: `shared/` for TypeScript types used by both frontend and backend

## Chess Engine Integration

### Standard Pieces
- Use Stockfish directly via UCI protocol for all standard chess pieces
- Engine handles move generation and position evaluation for normal gameplay

### Special Pieces
- Custom move generation required for pieces with special abilities
- Hybrid approach: Generate custom moves manually, use Stockfish to evaluate resulting positions
- Example special piece: Knight that can move twice per turn

## Key Technical Decisions

### Why Go for Backend?
- Superior handling of concurrent chess calculations via goroutines
- Can run multiple Stockfish instances simultaneously for different games
- Good portfolio showcase for backend development skills
- Better than Node.js for CPU-intensive chess engine workloads

### Why Stockfish Locally?
- Stockfish runs locally (compiled to WebAssembly for browser, or as process on server)
- No API calls or external dependencies for chess calculations
- Zero latency and offline capability
- No ongoing costs for AI computations

### Data Storage Considerations
- Chess games are very lightweight (1-5KB per complete game)
- FEN notation for positions (~100 bytes)
- Supabase free tier should handle thousands of games before hitting storage limits
- Player progression and unlocked pieces stored in database

## Deployment Strategy
- **Frontend**: Automatic Vercel deployment on git push
- **Backend**: Automatic Railway/Render deployment on git push  
- **Database**: Supabase (PostgreSQL with built-in auth)
- **CI/CD**: Turbo handles builds across both apps

## Future Mobile Considerations
- Current web stack can become PWA with minimal changes
- React Native migration possible (share game logic, rewrite UI)
- Capacitor option for native app wrapping
- Framer Motion drag system works with touch events

## Development Priorities (V1)
1. Basic chess gameplay with standard pieces
2. Simple shop system with piece upgrades
3. Progressive AI difficulty
4. Player progression/money system
5. Special piece mechanics (start with simple examples)