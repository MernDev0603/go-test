# Step 1: Build the Next.js app
FROM node:20-alpine AS build

# Set the working directory
WORKDIR /frontend

# Copy package.json and package-lock.json for installing dependencies
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Next.js app for production
RUN npm run build

# Step 2: Production image to serve the app
FROM node:20-alpine

# Set the working directory
WORKDIR /frontend

# Copy the node_modules and built files from the build stage
COPY --from=build /frontend /frontend

# Expose the port Next.js will run on
EXPOSE 3000

# Run the Next.js app in production mode
CMD ["npm", "start"]
