package handlers

// /help: Muestra ayuda sobre los comandos disponibles.
func Help() string {
	return `
			/fibonacci?num=N
			/createfile?name=filename&content=text&repeat=x
			/deletefile?name=filename
			/reverse?text=abc
			/toupper?text=abc
			/random?count=n&min=a&max=b
			/timestamp
			/hash?text=abc
			/simulate?seconds=s&task=name
			/sleep?seconds=s
			/loadtest?tasks=n&sleep=x
			/status
			/help
`
}

