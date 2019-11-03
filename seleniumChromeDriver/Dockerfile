FROM golang:1.13.1


RUN apt-get update
RUN apt-get install -y libnss3-dev libx11-dev vim

RUN curl -sS -o - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add &&\
    echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list &&\
    apt-get -y update &&\
    apt-get -y install google-chrome-stable

RUN go get -v github.com/tebeka/selenium \
    github.com/fatih/color \
    github.com/happierall/l \
    github.com/kr/pretty \
    github.com/mattn/go-colorable \
    github.com/mattn/go-isatty