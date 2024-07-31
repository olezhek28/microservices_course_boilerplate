package usecases

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i UserService -o ./mocks/ -s "_mock.go"
