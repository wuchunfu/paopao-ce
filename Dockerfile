# syntax=docker/dockerfile:1

# build frontend
FROM node:19-alpine AS frontend
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
WORKDIR /web
COPY web/ ./
RUN [ $EMBED_UI != yes ] || [ $USE_API_HOST != yes ] || echo "VITE_HOST='$API_HOST'">.env.local
RUN [ $EMBED_UI != yes ] || [ $USE_DIST != no ] || (yarn && yarn build)
RUN [ $EMBED_UI = yes ] || mkdir dist || echo ""

# build go mod cache
FROM bitbus/paopao-ce-backend-builder:latest AS gomodcache
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /paopao-ce
COPY go.mod .
COPY go.sum .
RUN go mod download

# build backend
FROM gomodcache AS backend
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
WORKDIR /paopao-ce
COPY . .
COPY --from=frontend /web/dist ./web/dist
ENV GOPROXY=https://goproxy.cn,direct
RUN [ $EMBED_UI != yes ] || make build TAGS='go_json'
RUN [ $EMBED_UI = yes ] || make build TAGS='slim embed go_json'

FROM bitbus/paopao-ce-backend-runner:latest
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
ENV TZ=Asia/Shanghai
WORKDIR /app/paopao-ce
COPY --from=backend /paopao-ce/release/paopao .
COPY --from=backend /paopao-ce/config.yaml.sample config.yaml
VOLUME ["/app/paopao-ce/custom"]
EXPOSE 8008
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao || exit 1
ENTRYPOINT ["/app/paopao-ce/paopao"]
CMD ["serve"]
