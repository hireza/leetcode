package main

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	hash := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	res := []string{}

	for _, v := range s {
		if val, ok := hash[string(v)]; ok {
			if len(res) > 0 {
				if res[len((res))-1] == val {
					res = res[:len(res)-1]
				} else {
					return false
				}

				continue
			}

			return false
		}

		res = append(res, string(v))
	}

	return len(res) == 0
}
