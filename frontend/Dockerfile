# ─────────── STAGE 1: Builder ───────────
FROM node:18-alpine AS builder
WORKDIR /app/frontend

# Build args for your VITE_… vars
ARG VITE_APP_API_URL
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

# Export them so Vite can pick them up
ENV NODE_ENV=production \
    VITE_APP_API_URL=$VITE_APP_API_URL \
    VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
    VITE_FIREBASE_AUTH_DOMAIN=$VITE_FIREBASE_AUTH_DOMAIN \
    VITE_FIREBASE_PROJECT_ID=$VITE_FIREBASE_PROJECT_ID \
    VITE_FIREBASE_STORAGE_BUCKET=$VITE_FIREBASE_STORAGE_BUCKET

# Copy only package manifests, install all deps (including dev)
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# Copy all frontend source & build
COPY frontend ./
RUN npm run build



# ─────────── STAGE 2: Runtime ───────────
FROM node:18-alpine AS runtime
WORKDIR /app

# Install a static server
RUN npm install -g serve

# Pull in just the built assets
COPY --from=builder /app/frontend/dist ./dist

EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
