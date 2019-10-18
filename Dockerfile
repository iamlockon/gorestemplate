FROM scratch

WORKDIR $GOPATH/src/github.com/iamlockon/gorestemplate
COPY . $GOPATH/src/github.com/iamlockon/gorestemplate

EXPOSE 8001
ENTRYPOINT [ "./gin-template" ]