FROM node:lts-alpine as build-stage

WORKDIR /app

COPY package.json .
COPY yarn.lock .
COPY . .

RUN yarn && yarn build

FROM nginx:1.21-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
