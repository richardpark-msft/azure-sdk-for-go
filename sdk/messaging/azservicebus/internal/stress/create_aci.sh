#!/bin/bash

set -e

# .env.sh is created by you and should have these variables
# in it, filled out and with values.
#
# export SERVICEBUS_CONNECTION_STRING=
# export APPINSIGHTS_INSTRUMENTATIONKEY=
# export REGISTRY_USERNAME=
# export REGISTRY_PASSWORD=
# export APP_NAME=
# export RESOURCE_GROUP=
# export DOCKER_REGISTRY=
# export IMAGE_NAME=
. .env.sh

docker build -t ${IMAGE_NAME} -f Dockerfile.aci ../..
az acr login -n ${DOCKER_REGISTRY}
docker push ${IMAGE_NAME}

echo "Creating the container"

az container create \
    --name ${APP_NAME} \
    --image ${IMAGE_NAME} \
    --command-line "./stress tests infiniteSendAndReceive" \
    --resource-group "${RESOURCE_GROUP}" \
    --location eastus \
    --memory 4 \
    --restart-policy Never \
    --registry-login-server ${DOCKER_REGISTRY} \
    --registry-password ${REGISTRY_PASSWORD} \
    --registry-username ${REGISTRY_USERNAME} \
    --secure-environment-variables \
        "SERVICEBUS_CONNECTION_STRING=${SERVICEBUS_CONNECTION_STRING}" \
        "APPINSIGHTS_INSTRUMENTATIONKEY=${APPINSIGHTS_INSTRUMENTATIONKEY}"
