# ─────────── STAGE 1: Builder ───────────
FROM node:18-alpine AS builder
# Set working directory inside container to the frontend folder
WORKDIR /app/frontend

# 1. Declare build args for your VITE_… vars
ARG VITE_APP_API_URL
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

# 2. Export them so Vite picks them up during build
ENV VITE_APP_API_URL=$VITE_APP_API_URL \
    VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
    VITE_FIREBASE_AUTH_DOMAIN=$VITE_FIREBASE_AUTH_DOMAIN \
    VITE_FIREBASE_PROJECT_ID=$VITE_FIREBASE_PROJECT_ID \
    VITE_FIREBASE_STORAGE_BUCKET=$VITE_FIREBASE_STORAGE_BUCKET

# 3. Copy dependency manifests and install ALL deps (including dev)
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# 4. Switch to production mode for an optimized build
ENV NODE_ENV=production

# 5. Copy the rest of your frontend source and build
COPY frontend ./
RUN npm run build



# ─────────── STAGE 2: Runtime ───────────
FROM node:18-alpine AS runtime
# Serve files from /app
WORKDIR /app

# 6. Install a static server
RUN npm install -g serve

# 7. Copy built assets from builder
COPY --from=builder /app/frontend/dist ./dist

# 8. Expose port and start serving
EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
