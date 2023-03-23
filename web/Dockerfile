# syntax=docker/dockerfile:1.4

# 1. For build React app
FROM node:lts AS development

# Set working directory
WORKDIR /app

# 
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

# Same as npm install
RUN npm ci

COPY . /app

ENV CI=true

ARG REACT_APP_VIDEOS_SERVICE_API_URL
ENV REACT_APP_VIDEOS_SERVICE_API_URL $REACT_APP_VIDEOS_SERVICE_API_URL

ARG REACT_APP_TOPIC
ENV REACT_APP_TOPIC $REACT_APP_TOPIC

RUN npm run build

# 2. For Nginx setup
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

# Remove default nginx static assets
RUN rm -rf ./*

# Copy static assets from builder stage
COPY --from=development /app/build .

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]
