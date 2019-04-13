package retry

func With(fn func(int) error, retries int) error {
	var err error
	attempts := 0
	for {
		err = fn(attempts)
		attempts++
		if err == nil || attempts > retries {
			break
		}
	}
	return err
}
