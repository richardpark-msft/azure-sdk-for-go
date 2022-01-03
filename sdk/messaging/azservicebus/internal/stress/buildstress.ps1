set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0


$IMAGE_NAME="stresstesttbiruti6oi24k.azurecr.io/ripark/gosbtest:latest"

Write-Host $IMAGE_NAME

# call az acr login -n stresstestregistry
# if errorlevel 1 goto :EOF


# pushd ..\..

# call docker build . -t %IMAGE_NAME%
# if errorlevel 1 goto :EOF

# call docker push %IMAGE_NAME%
# if errorlevel 1 goto :EOF

kubectl replace -f job.yaml --force