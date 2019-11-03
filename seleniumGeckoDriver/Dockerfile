FROM golang:1.13.1


RUN apt-get update  
RUN apt-get install -y wget vim default-jre xvfb

# install firefox
RUN apt-get install -y bzip2 libdbus-glib-1-2
RUN wget -O ~/FirefoxSetup.tar.bz2 "https://download.mozilla.org/?product=firefox-latest&os=linux64"
RUN tar xjf ~/FirefoxSetup.tar.bz2 -C /opt/ && \
    mkdir -p /usr/lib/firefox/firefox && \
    ln -s /opt/firefox/firefox /usr/lib/firefox/firefox