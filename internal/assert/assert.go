package assert

func Assert(err error) {
	if err != nil {
		panic(err.Error())
	}
}
