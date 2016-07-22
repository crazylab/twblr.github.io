package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	result := []string{}
	for people, interests:= range data{
		if contains(interests, interest){
			result = append(result, people)
		}
	}
	return result
}

func contains(src []string, elem string) bool {
	for _,value := range src{
		if value == elem{
			return true
		}
	}
	return false
}
