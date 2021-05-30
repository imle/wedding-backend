docker build -t stevenimle/wedding-backend:latest -t stevenimle/wedding-backend:0.1.0 . || exit 1
docker push stevenimle/wedding-backend:latest || exit 2
docker push stevenimle/wedding-backend:0.1.0 || exit 2

docker build -t stevenimle/wedding-migrations:latest -t stevenimle/wedding-migrations:0.1.0 -f migrations.dockerfile . || exit 1
docker push stevenimle/wedding-migrations:latest || exit 2
docker push stevenimle/wedding-migrations:0.1.0 || exit 2

docker build -t harbor.imle.io/library/stevenimle/wedding-backend:latest -t harbor.imle.io/library/stevenimle/wedding-backend:0.1.0 . || exit 1
docker push harbor.imle.io/library/stevenimle/wedding-backend:latest || exit 2
docker push harbor.imle.io/library/stevenimle/wedding-backend:0.1.0 || exit 2

docker build -t harbor.imle.io/library/stevenimle/wedding-migrations:latest -t harbor.imle.io/library/stevenimle/wedding-migrations:0.1.0 -f migrations.dockerfile . || exit 1
docker push harbor.imle.io/library/stevenimle/wedding-migrations:latest || exit 2
docker push harbor.imle.io/library/stevenimle/wedding-migrations:0.1.0 || exit 2
