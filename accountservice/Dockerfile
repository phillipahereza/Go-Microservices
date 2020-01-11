FROM iron/base

EXPOSE 6767
ADD accountservice-linux-amd64 /
#HEALTHCHECK --interval=5s --timeout=5s CMD ["./healthchecker-linux-amd64", "-addr=accountservice","-port=6767"] || exit 1
ENTRYPOINT ["./accountservice-linux-amd64"]