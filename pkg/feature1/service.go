package feature1

// fake person
var fakePerson = Person{"Tian Zhi", "Male", 22}

// GetPerson returns the fake person
func GetPerson() Person {
	return fakePerson
}
