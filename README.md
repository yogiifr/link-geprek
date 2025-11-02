# LINK GEPREK

🚧 On work

> A fullstack URL shortener built with **Go (Fiber)** + **Next.js 15 (TypeScript)** + **Postgres + Redis**.  
> Real-time analytics, type-safe, scalable, production-ready.

---

## Core Requirements

- Shorten long URL → short code
- Redirect short → original
- Count clicks
- Analytics dashboard

## Data Flow

1. User → POST /api/shorten
2. Validate → Generate hashid → Save DB
3. Return short URL
4. User click → GET /:code
5. Redis cache? → Redirect + INCR
6. No? → DB → Cache → Redirect

## Security

- URL validation (govalidator)
- Rate limiting (100/min)
- Hashids salt
- CORS

---

## Tech Stack

| Layer      | Tech |
|-----------|------|
| Backend   | Go + Fiber + GORM |
| Frontend  | Next.js 15 + TypeScript + Drizzle + shadcn/ui + Recharts |
| Database  | PostgreSQL + Redis |
| Deploy    | Docker + Vercel + Railway |

---

## Phase Progress Log

| Phase | Status | Date | Notes |
|-------|--------|------|-------|
| **Phase 1: Setup Monorepo + Docker + DB** | ✅ Done | 2025-11-01 | Docker up, GORM & Drizzle schema sync |
| **Phase 2: Backend MVP (Shorten + Redirect)** | ⏳ | - | - |
| **Phase 3: Frontend MVP (Form + List)** | ⏳ | - | - |
| **Phase 4: Analytics Dashboard** | ⏳ | - | - |
| **Phase 5: Production Polish** | ⏳ | - | - |
