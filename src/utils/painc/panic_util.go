package paincUtil

func CustomPanic(text string, err error) {
	panic(text + ": {" + err.Error() + "}")
}

func Test() {

	err := recover()
	if err != nil {

	}
}
