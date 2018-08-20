zip -r app.zip Procfile appdconfig/ glide.lock glide.yaml main.go logging/ sinks/ vendor/
mkdir -p ../resources
cp app.zip ../resources
