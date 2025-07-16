# ───────────── STAGE 1: Build ─────────────
FROM node:18-alpine AS builder

# 1. Set working directory
WORKDIR /app

# 2. Copy dependency manifests
COPY package.json package-lock.json ./

# 3. Copy env‑files so Vite can read VITE_… variables
#    (If you have .env.production, copy that instead or in addition.)
COPY .env ./

# 4. Install deps
RUN npm install

# 5. Copy the rest of your source
COPY . .

# 6. Build for production
RUN npm run build



# ───────────── STAGE 2: Serve ─────────────
FROM node:18-alpine AS runtime

# 7. Install a static file server
RUN npm install -g serve

# 8. Copy only the built assets
WORKDIR /app
COPY --from=builder /app/dist ./dist

# 9. Expose the port
EXPOSE 5173

# 10. Start the server
CMD ["serve", "-s", "dist", "-l", "5173"]
