# ─────────── STAGE 1: Builder ───────────
FROM node:18-alpine AS builder
WORKDIR /app

# 1. Accept build args for your VITE_… vars
ARG VITE_APP_API_URL
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

# 2. Export them so Vite can see them during npm run build
ENV NODE_ENV=production \
    VITE_APP_API_URL=$VITE_APP_API_URL \
    VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
    VITE_FIREBASE_AUTH_DOMAIN=$VITE_FIREBASE_AUTH_DOMAIN \
    VITE_FIREBASE_PROJECT_ID=$VITE_FIREBASE_PROJECT_ID \
    VITE_FIREBASE_STORAGE_BUCKET=$VITE_FIREBASE_STORAGE_BUCKET

# 3. Install deps & build
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

# ─────────── STAGE 2: Runtime ───────────
FROM node:18-alpine AS runtime
WORKDIR /app

# 4. Install a static asset server
RUN npm install -g serve

# 5. Copy the production build
COPY --from=builder /app/dist ./dist

# 6. Expose the port and serve
EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
