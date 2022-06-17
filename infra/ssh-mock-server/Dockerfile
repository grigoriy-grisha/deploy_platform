FROM ubuntu:latest

RUN apt update && apt install  openssh-server sudo -y

RUN groupadd sshgroup && useradd -ms /bin/bash -g sshgroup sshuser

RUN mkdir -p /home/sshuser/.ssh

RUN service ssh start

EXPOSE 22

CMD ["/usr/sbin/sshd","-D"]