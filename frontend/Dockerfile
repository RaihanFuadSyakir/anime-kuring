# Use an official Node.js runtime as a base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install
# Copy the entire project to the working directory
COPY . .
RUN npm run build
# Expose the port on which the application will run (adjust as needed)
EXPOSE 5173

# Define the command to run the application when the container starts
CMD ["npm", "run","dev","--","--host"]

