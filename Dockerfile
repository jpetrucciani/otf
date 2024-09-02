FROM alpine:3.20

# bubblewrap is for sandboxing, and git permits pulling modules via
# the git protocol
RUN apk add --no-cache bubblewrap git python3 py3-pip aws-cli

COPY otfd /usr/local/bin/otfd

ENTRYPOINT ["/usr/local/bin/otfd"]
