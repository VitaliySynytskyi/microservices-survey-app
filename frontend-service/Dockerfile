# Use the official Node.js image as the base image
FROM node:lts-alpine

# Install a simple HTTP server to serve static content
RUN npm install -g http-server

# Set the working directory inside the container to /app
WORKDIR /app

# Copy package.json and package-lock.json (if available) to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the project files and folders to the working directory
COPY . .

# Build the application for production with minification
RUN NODE_OPTIONS=--openssl-legacy-provider npm run build

# Expose the default port for http-server
EXPOSE 8080

# Serve the built application using http-server
CMD [ "http-server", "dist" ]
