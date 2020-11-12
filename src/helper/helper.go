package helper

func DeleteIntersectionElements(baseElements []string, searchElements []string) []string {
	intermediateResult := map[string]bool{}
	for _, v := range baseElements {
		intermediateResult[v] = true
	}
	for _, v := range searchElements {
		delete(intermediateResult, v)
	}
	var result []string
	for k, _ := range intermediateResult {
		result = append(result, k)
	}
	return result
}
