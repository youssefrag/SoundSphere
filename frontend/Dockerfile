# ─────────── STAGE 1: Builder ───────────
FROM node:18-alpine AS builder
WORKDIR /app

# 1. Declare build args for each VITE_… var
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

# 2. Export them so Vite sees them during `npm run build`
ENV VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
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
RUN npm install -g serve

# 4. Copy the production build
COPY --from=builder /app/dist ./dist
EXPOSE 5173

# 5. Serve it
CMD ["serve", "-s", "dist", "-l", "5173"]
