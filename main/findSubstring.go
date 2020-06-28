package main

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return nil
	}
	res := make([]int, 0)
	if len(s) < len(words[0])*len(words) {
		return res
	}
	m := make(map[string]int)
	for _, value := range words {
		_, ok := m[value]
		if ok {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	help := func(s string) bool {
		start := 0
		end := start + len(words[0])
		item := make([]string, len(words), len(s))
		for end <= len(s) {
			item[start/len(words[0])] = s[start:end]
			start += len(words[0])
			end += len(words[0])
		}
		mtstring, _ := json.Marshal(m)
		var mt map[string]interface{}
		_ = json.Unmarshal([]byte(mtstring), &mt)
		for _, value := range item {
			_, ok := mt[value]
			if ok {
				mt[value] = mt[value].(float64) - 1
				if mt[value].(float64) == 0 {
					delete(mt, value)
				}
			} else {
				return false
			}
		}
		return len(mt) == 0
	}
	sum := len(words) * len(words[0])
	i := 0
	j := i + sum
	for j <= len(s) {
		item := s[i:j]
		if help(item) {
			res = append(res, i)
		}
		i++
		j++
	}
	return res
}
