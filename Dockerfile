FROM scratch 

ADD src/go-hpa/go-hpa /

CMD ["./go-hpa"]
