# ─────────── STAGE 1: Builder ───────────
FROM node:18-alpine AS builder
WORKDIR /app/frontend

# 1. Accept build args for your VITE_… vars
ARG VITE_APP_API_URL
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

# 2. Expose those as env vars (Vite picks them up at build time)
ENV VITE_APP_API_URL=$VITE_APP_API_URL \
    VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
    VITE_FIREBASE_AUTH_DOMAIN=$VITE_FIREBASE_AUTH_DOMAIN \
    VITE_FIREBASE_PROJECT_ID=$VITE_FIREBASE_PROJECT_ID \
    VITE_FIREBASE_STORAGE_BUCKET=$VITE_FIREBASE_STORAGE_BUCKET

# 3. Install **all** dependencies (including dev) so `vite` is available
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# 4. Now switch to production mode for an optimized build
ENV NODE_ENV=production

# 5. Copy source & build
COPY frontend ./
RUN npm run build


# ─────────── STAGE 2: Runtime ───────────
FROM node:18-alpine AS runtime
WORKDIR /app

# 6. Static file server
RUN npm install -g serve

# 7. Pull in the built assets
COPY --from=builder /app/frontend/dist ./dist

# 8. Serve on port 5173
EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
