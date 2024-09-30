package cache

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i RedisClient -o ./mocks -p redis_client_mock -s "_minimock.go"
