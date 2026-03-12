# Phase 8: Deployment & Environment

This plan ensures the app is ready for production-like environments.

## Objective
Containerize both frontend and backend and use a single `.env` file.

## Key Files & Context
- `Dockerfile.backend`: Multi-stage build for Go.
- `Dockerfile.frontend`: Build with Vite, serve with Nginx.
- `docker-compose.yml`: Orchestra.
- `.env`: Shared environment file.

## Implementation Steps
1. **Dockerize Backend**:
   - `FROM golang:alpine AS builder` (build).
   - `FROM alpine` (run).
2. **Dockerize Frontend**:
   - `FROM node:alpine AS builder` (build).
   - `FROM nginx:alpine` (serve assets).
3. **Unified Config**:
   - Configure Docker to pass relevant `.env` variables to both services.
   - Setup Nginx to proxy API calls from `/api` to the backend container.
4. **Final Polish**:
   - Readme with setup instructions.
   - Basic health check for Docker.

## Verification & Testing
- `docker-compose up --build`: Ensure both containers start and talk to each other.
- Access the app via browser.
