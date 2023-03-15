package panicUtil

func CustomPanic(text string, err error) {
	panic(text + ": {" + err.Error() + "}")
}
