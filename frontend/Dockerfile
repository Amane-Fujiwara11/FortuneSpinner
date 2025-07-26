FROM node:18-alpine AS builder

WORKDIR /app

# Copy package files first for better caching
COPY package*.json ./

# Use npm ci for faster, more reliable installs
RUN npm ci --only=production && \
    npm cache clean --force

# Copy source code
COPY . .

# Install dev dependencies and build
RUN npm ci && \
    npm run build

# Production stage
FROM nginx:alpine

# Copy built files to nginx
COPY --from=builder /app/build /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]