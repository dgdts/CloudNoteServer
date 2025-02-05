FROM xsjop-harbor.seasungame.com/xsjweb-k8s/base/golang:1.21.8-alpine3.18 AS builder
LABEL stage=gobuilder

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

ARG ARTIFACT_ID

WORKDIR /build
COPY ./$ARTIFACT_ID/output .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add tzdata




FROM xsjop-harbor.seasungame.com/xsjweb-k8s/base/alpine:3.18

ARG ARTIFACT_ID
USER root 

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /build /application/$ARTIFACT_ID

ENV TZ=Asia/Shanghai
WORKDIR /application/$ARTIFACT_ID

# COPY /usr/share/zoneinfo /usr/share/zoneinfo
# COPY ./$ARTIFACT_ID/output .
EXPOSE 8888

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN chmod +x /application/$ARTIFACT_ID/bin/trade && chmod 4755 /bin/busybox && apk update && apk add curl && apk add busybox-extras

ENTRYPOINT ["./bin/trade"]