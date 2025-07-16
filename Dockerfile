# ─────────── STAGE 1: Builder ───────────
FROM node:18-alpine AS builder
WORKDIR /app/frontend

ARG VITE_APP_API_URL
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

# Expose env‑vars for Vite
ENV NODE_ENV=production \
    VITE_APP_API_URL=$VITE_APP_API_URL \
    VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
    VITE_FIREBASE_AUTH_DOMAIN=$VITE_FIREBASE_AUTH_DOMAIN \
    VITE_FIREBASE_PROJECT_ID=$VITE_FIREBASE_PROJECT_ID \
    VITE_FIREBASE_STORAGE_BUCKET=$VITE_FIREBASE_STORAGE_BUCKET

# Install deps from frontend/package.json
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# Copy the rest of your frontend code and build
COPY frontend ./
RUN npm run build

# ─────────── STAGE 2: Runtime ───────────
FROM node:18-alpine AS runtime
WORKDIR /app

RUN npm install -g serve

# Copy built assets out of the builder
COPY --from=builder /app/frontend/dist ./dist

EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
