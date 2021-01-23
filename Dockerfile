FROM debian:bullseye-slim
LABEL maintainer=""

ARG LOAD_LANG=eng

RUN apt update \
    && apt install -y \
      ca-certificates \
      libtesseract-dev=4.1.1-2+b1 \
      tesseract-ocr=4.1.1-2+b1 \
      golang=2:1.15~1

ENV GO111MODULE=on
ENV GOPATH=${HOME}/go
ENV PATH=${PATH}:${GOPATH}/bin

ADD . $GOPATH/src/ocr-api
WORKDIR $GOPATH/src/ocr-api
RUN go get -v ./... && go install .

# Load languages
RUN if [ -n "${LOAD_LANG}" ]; then apt-get install -y tesseract-ocr-${LOAD_LANG}; fi

ENV PORT=8080
CMD ["ocr-api"]
