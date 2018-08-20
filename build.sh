zip -r nozzle.zip Procfile appdconfig/ glide.lock glide.yaml main.go logging/ sinks/ vendor/
mkdir -p ../resources
cp nozzle.zip ../resources
