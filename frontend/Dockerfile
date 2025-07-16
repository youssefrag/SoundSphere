# ───────────── STAGE 1: Build ─────────────
FROM node:18-alpine AS builder

# 1. Set working directory
WORKDIR /app

# 2. Copy manifest & install deps
COPY package.json package-lock.json ./
RUN npm install

# 3. Copy source & build
COPY . .
RUN npm run build



# ───────────── STAGE 2: Runtime ─────────────
FROM node:18-alpine AS runtime

# 4. Install `serve` to deliver static files
RUN npm install -g serve

# 5. Create app directory & pull in built assets
WORKDIR /app
COPY --from=builder /app/dist ./dist

# 6. Expose the port we’ll serve on
EXPOSE 5173

# 7. Start the static server
CMD ["serve", "-s", "dist", "-l", "5173"]

