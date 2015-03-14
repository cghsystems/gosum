package retry

func Execute(count int, retryFunc func()) error {
	for i := 0; i < count; i++ {
		retryFunc()
	}
	return nil
}
