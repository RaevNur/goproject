echo "Building image"
docker build -t ascii-art-web .
echo "Running image"
docker run -d -p 8080:8080 --name web ascii-art-web
echo "Images list"
docker images
echo "Containers list"
docker container ls