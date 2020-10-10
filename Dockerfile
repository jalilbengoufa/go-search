FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/jalilbengoufa/go-search/
WORKDIR /go/src/github.com/jalilbengoufa/go-search/

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
    export GIN_MODE=release \
	main; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080