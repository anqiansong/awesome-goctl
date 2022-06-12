FROM golang
RUN mkdir -p /app/goctl/test
WORKDIR /app/goctl/test

COPY goctl.sh .
RUN chmod +x goctl.sh

CMD ["/bin/bash","-c","/app/goctl/test/goctl.sh"]