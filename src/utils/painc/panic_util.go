package paincUtil

func CustomPanic(text string, err error) {
	panic(text + ": {" + err.Error() + "}")
}
