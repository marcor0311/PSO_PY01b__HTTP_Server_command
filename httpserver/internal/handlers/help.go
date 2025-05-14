package handlers

// /help: Shows available commands
func HelpText() string {
	return `
	Available commands:
	/fibonacci?num=N           					- Returns the Nth Fibonacci number
	/random?count=n&min=a&max=b 				- Generates a list of random numbers
	/reverse?text=abc          					- Reverses the input text
	/toupper?text=abc          					- Converts text to uppercase
	/hash?text=abc             					- Returns SHA-256 hash
	/timestamp                 					- Returns current server time
	/help                      					- Lists available commands
	/status                    					- Shows server status
	/createfile?name=x&content=y&repeat=n 		- Creates a file
	/deletefile?name=x         					- Deletes a file
	/simulate?seconds=n       					- Simulates a blocking task
	/sleep?seconds=n           					- Sleeps for n seconds
	/loadtest?tasks=n&sleep=s  					- Starts concurrent load test
`
}
