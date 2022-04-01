go run . "hello" standard
go run . "HELLO" shadow
go run . "nice 2 meet you" thinkertoy
go run . "you & me" standard
go run . "123" shadow
go run . "/(\")" thinkertoy
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow
go run . "\"#$%&/()*+,-./" thinkertoy
go run . "It's Working" thinkertoy

echo -e "\e[30;1;45mTESTS WITH RANDOM STRINGS\033[0m"
go run . "RNDtext" zigzag
go run . "lower 90" zigzag
go run . "A~-=" zigzag
go run . "ab  1./,TRY" zigzag

