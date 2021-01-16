docker build . -t stevenimle/wedding-backend:latest -t stevenimle/wedding-backend:0.1.0 || exit 1
docker push stevenimle/wedding-backend:latest || exit 2
docker push stevenimle/wedding-backend:0.1.0 || exit 2
