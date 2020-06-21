FROM node:10
WORKDIR /opt
RUN npm install -g json-server@0.13.0
COPY db.json .
COPY routes.json .
CMD json-server db.json --routes routes.json
