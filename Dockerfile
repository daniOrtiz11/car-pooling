FROM golang:latest 

#RUN apk --no-cache add ca-certificates=20190108-r0 libc6-compat=1.1.19-r10

WORKDIR /table

EXPOSE 9091

COPY table-booking /table

COPY local.env /table

ENTRYPOINT [ "/table/table-booking"]

