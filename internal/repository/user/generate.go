package user_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i UserRepository -o ./mocks -p user_repository_mock -s "_minimock.go"
