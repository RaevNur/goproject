go run . "First\nTest" shadow --output=test00.txt
go run . "hello" standard --output=test01.txt
go run . "123 -> #$%" standard --output=test02.txt
go run . "432 -> #$%&@" shadow --output=test03.txt
go run . "There" shadow --output=test04.txt
go run . "123 -> \"#$%@" thinkertoy --output=test05.txt
go run . "2 you" thinkertoy --output=test06.txt
go run . "Testing long output!" standard --output=test07.txt

