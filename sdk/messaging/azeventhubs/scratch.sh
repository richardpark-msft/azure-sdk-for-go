#!/bin/bash

helm template goeh \
    ./internal/eh/stress \
    -n ripark \
    --set stress-test-addons.env=pg \
    --values ./internal/eh/stress/generatedValues.yaml \
    --debug > deployablething.yaml

helm upgrade goeh \
    ./internal/eh/stress \
    -n ripark \
    --install \
    --set stress-test-addons.env=pg \
    --values ./internal/eh/stress/generatedValues.yaml