FROM node:14


RUN mkdir /app

WORKDIR /app

COPY . .

RUN npm install

EXPOSE 3002

CMD ["node", "app.js"]