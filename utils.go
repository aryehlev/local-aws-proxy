package local_aws_proxy

func getFirstFromMulti(multiMap map[string][]string) map[string]string {
	singularMap := make(map[string]string)
	for key, val := range multiMap {
		if len(val) > 0 {
			singularMap[key] = val[0]
		}
	}

	return singularMap
}
