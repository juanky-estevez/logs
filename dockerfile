FROM node:22-alpine

RUN apk update && apk upgrade
RUN npm install -g npm@latest