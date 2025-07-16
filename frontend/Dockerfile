# ─── STAGE 1: Builder ───
FROM node:18-alpine AS builder
WORKDIR /app

ARG VITE_APP_API_URL
ARG VITE_FIREBASE_API_KEY
ARG VITE_FIREBASE_AUTH_DOMAIN
ARG VITE_FIREBASE_PROJECT_ID
ARG VITE_FIREBASE_STORAGE_BUCKET

ENV VITE_APP_API_URL=$VITE_APP_API_URL \
    VITE_FIREBASE_API_KEY=$VITE_FIREBASE_API_KEY \
    VITE_FIREBASE_AUTH_DOMAIN=$VITE_FIREBASE_AUTH_DOMAIN \
    VITE_FIREBASE_PROJECT_ID=$VITE_FIREBASE_PROJECT_ID \
    VITE_FIREBASE_STORAGE_BUCKET=$VITE_FIREBASE_STORAGE_BUCKET

COPY package.json package-lock.json ./
RUN npm install

# only now set production so build is optimized
ENV NODE_ENV=production

COPY . .
RUN npm run build

# ─── STAGE 2: Runtime ───
FROM node:18-alpine AS runtime
WORKDIR /app

RUN npm install -g serve
COPY --from=builder /app/dist ./dist

EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
