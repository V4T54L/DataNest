# frontend/Dockerfile
FROM node:23 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Serve the app with a lightweight HTTP server
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html

# Expose the port the app runs on
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]