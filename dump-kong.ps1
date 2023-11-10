docker run -i --rm --network=cu-get-ref-backend_default -v ${PWD}:/deck kong/deck --kong-addr=http://kong:8001/ -o /deck/kong.yaml dump
