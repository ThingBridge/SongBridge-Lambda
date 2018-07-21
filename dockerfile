FROM debian:latest

RUN apt-get -y update
RUN apt-get -y install python
RUN apt-get -y install python-pip
RUN pip install awscli --upgrade --user

CMD export PATH=~/.local/bin/ && /bin/bash