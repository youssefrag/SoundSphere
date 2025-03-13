# Step 1: Use an official Node.js image as a parent image
FROM node:18-alpine

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy package.json and package-lock.json to the working directory
COPY package.json package-lock.json ./

# Step 4: Install the dependencies
RUN npm install

# Step 5: Copy the rest of the application code
COPY . .

# Step 6: Expose the port your app will run on
EXPOSE 3000

# Step 7: Run your frontend application
CMD ["npm", "start"]
