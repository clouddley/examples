# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json first for dependency installation
COPY package*.json ./

# Install dependencies
RUN npm install --omit=dev  # Use --omit=dev for production builds

# Copy the rest of the application code
COPY . .

# Expose the application port (ensure it matches the port your app listens on)
EXPOSE 3011

# Define the command to run the application
CMD ["node", "./bin/www"]




