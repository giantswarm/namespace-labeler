FROM quay.io/giantswarm/alpine:3.9-giantswarm

USER giantswarm
ADD ./namespace-labeler /namespace-labeler

ENTRYPOINT ["/namespace-labeler"]
