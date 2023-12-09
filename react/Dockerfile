#build stage
FROM node:16-alpine as build
WORKDIR /root
COPY ./package*.json ./
RUN npm ci --cache .npm
COPY . . 
RUN npm run build

#final stage
FROM nginx:alpine as prod
EXPOSE 3000
COPY --from=build /root/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
