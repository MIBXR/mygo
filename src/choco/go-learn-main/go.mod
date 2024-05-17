module choco/go-learn-main

go 1.22.2

replace choco/greetings => ../greetings

replace choco/utils => ../utils

require (
	choco/greetings v0.0.0-00010101000000-000000000000
	choco/utils v0.0.0-00010101000000-000000000000
)
