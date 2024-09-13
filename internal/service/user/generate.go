package user_service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i UserService -o ./mocks -p user_service_mock -s "_minimock.go"
