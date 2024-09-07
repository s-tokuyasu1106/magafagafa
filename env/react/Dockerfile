FROM node:latest
# WORKDIR /app/react
ENV LANG=C.UTF-8
ENV TZ=Asia/Tokyo

WORKDIR /app/react
RUN cd /app/react
RUN npm init -y
RUN npm install
# RUN npm install --save prop-types
RUN npm install -g create-react-app --template typescript
RUN npm i @chakra-ui/react @emotion/react @emotion/styled framer-motion
RUN npm install react-router-dom
RUN npm install swr
# RUN npm install --save react@^17
# RUN npm install axios
# RUN npm install --save --legacy-peer-deps @chakra-ui/react
# RUN npm i @emotion/react@^11
# RUN npm i @emotion/css@^11
# RUN npm i @emotion/styled@^11
# RUN npm install --save --legacy-peer-deps framer-motion@^4
# RUN npm install --save --legacy-peer-deps  @mui/material
# RUN npm i --save react-ztext
# RUN npm i @material-ui/core
# RUN npm i --save firebase
# RUN npm install --save --legacy-peer-deps react-router-dom@
# RUN npm install npm-link-shared -g
# RUN npm install -D tailwindcss postcss autoprefixer
# RUN npx tailwindcss init
