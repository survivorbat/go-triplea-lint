package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	service *UserService
}

func (s *UserServiceSuite) SetupTest() {
	s.service = NewUserService()
}

func (s *UserServiceSuite) TestAddUser() {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	// Act
	s.service.AddUser(person)

	// Assert
	user, ok := s.service.GetUser("Josh")
	assert.True(s.T(), ok)
	assert.Equal(s.T(), "Josh", user.Name)
}

func (s *UserServiceSuite) TestGetUser() {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}
	s.service.AddUser(person)

	// Act
	user, ok := s.service.GetUser("Josh")

	// Assert
	assert.True(s.T(), ok)
	assert.Equal(s.T(), "Josh", user.Name)
}

func (s *UserServiceSuite) TestDeleteUser() {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}
	s.service.AddUser(person)

	// Act
	s.service.DeleteUser("Josh")

	// Assert
	_, ok := s.service.GetUser("Josh")
	assert.False(s.T(), ok)
}

func (s *UserServiceSuite) TestGetUserNotFound() {
	// Act
	user, ok := s.service.GetUser("NonExistent")

	// Assert
	assert.False(s.T(), ok)
	assert.Nil(s.T(), user)
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

type InvalidUserServiceSuite struct {
	suite.Suite
	service *UserService
}

func (s *InvalidUserServiceSuite) SetupTest() {
	s.service = NewUserService()
}

func (s *InvalidUserServiceSuite) TestAddUser() {
	person := &Person{Name: "Josh", Age: 30} // want `// Arrange statement expected`
	s.service.AddUser(person)                // want `// Act statement expected`
	user, ok := s.service.GetUser("Josh")    // want `// Assert statement expected`
	assert.True(s.T(), ok)
	assert.Equal(s.T(), "Josh", user.Name)
}

func (s *InvalidUserServiceSuite) TestGetUser() {
	person := &Person{Name: "Josh", Age: 30} // want `// Arrange statement expected`
	s.service.AddUser(person)
	user, ok := s.service.GetUser("Josh") // want `// Act statement expected`
	assert.True(s.T(), ok)                // want `// Assert statement expected`
	assert.Equal(s.T(), "Josh", user.Name)
}

func (s *InvalidUserServiceSuite) TestDeleteUser() {
	person := &Person{Name: "Josh", Age: 30} // want `// Arrange statement expected`
	s.service.AddUser(person)
	s.service.DeleteUser("Josh")       // want `// Act statement expected`
	_, ok := s.service.GetUser("Josh") // want `// Assert statement expected`
	assert.False(s.T(), ok)
}

func (s *InvalidUserServiceSuite) TestGetUserNotFound() {
	user, ok := s.service.GetUser("NonExistent") // want `// Act statement expected`
	assert.False(s.T(), ok)                      // want `// Assert statement expected`
	assert.Nil(s.T(), user)
}

func TestInvalidUserServiceSuite(t *testing.T) {
	suite.Run(t, new(InvalidUserServiceSuite))
}

type NestedSuite struct {
	suite.Suite
	service *UserService
}

func (s *NestedSuite) SetupTest() {
	s.service = NewUserService()
}

func (s *NestedSuite) TestAddUserWithSubtests() {
	tests := map[string]struct {
		person *Person
	}{
		"Josh": {person: &Person{Name: "Josh", Age: 30}},
		"Anne": {person: &Person{Name: "Anne", Age: 25}},
	}

	for name, testData := range tests {
		s.Run(name, func() {
			// Act
			s.service.AddUser(testData.person)

			// Assert
			user, ok := s.service.GetUser(testData.person.Name)
			assert.True(s.T(), ok)
			assert.Equal(s.T(), testData.person.Name, user.Name)
		})
	}
}

func (s *NestedSuite) TestGetUserWithSubtests() {
	tests := map[string]struct {
		person *Person
	}{
		"Josh": {person: &Person{Name: "Josh", Age: 30}},
		"Anne": {person: &Person{Name: "Anne", Age: 25}},
	}

	for name, testData := range tests {
		s.Run(name, func() {
			// Arrange
			s.service.AddUser(testData.person)

			// Act
			user, ok := s.service.GetUser(testData.person.Name)

			// Assert
			assert.True(s.T(), ok)
			assert.Equal(s.T(), testData.person.Name, user.Name)
		})
	}
}

func TestNestedSuite(t *testing.T) {
	suite.Run(t, new(NestedSuite))
}

type InvalidNestedSuite struct {
	suite.Suite
	service *UserService
}

func (s *InvalidNestedSuite) SetupTest() {
	s.service = NewUserService()
}

func (s *InvalidNestedSuite) TestAddUserWithSubtests() {
	tests := map[string]struct {
		person *Person
	}{
		"Josh": {person: &Person{Name: "Josh", Age: 30}},
		"Anne": {person: &Person{Name: "Anne", Age: 25}},
	}

	for name, testData := range tests {
		s.Run(name, func() {
			s.service.AddUser(testData.person)                  // want `// Act statement expected`
			user, ok := s.service.GetUser(testData.person.Name) // want `// Assert statement expected`
			assert.True(s.T(), ok)
			assert.Equal(s.T(), testData.person.Name, user.Name)
		})
	}
}

func (s *InvalidNestedSuite) TestGetUserWithSubtests() {
	tests := map[string]struct {
		person *Person
	}{
		"Josh": {person: &Person{Name: "Josh", Age: 30}},
		"Anne": {person: &Person{Name: "Anne", Age: 25}},
	}

	for name, testData := range tests {
		s.Run(name, func() {
			s.service.AddUser(testData.person)                  // want `// Arrange statement expected`
			user, ok := s.service.GetUser(testData.person.Name) // want `// Act statement expected`
			assert.True(s.T(), ok)                              // want `// Assert statement expected`
			assert.Equal(s.T(), testData.person.Name, user.Name)
		})
	}
}

func TestInvalidNestedSuite(t *testing.T) {
	suite.Run(t, new(InvalidNestedSuite))
}

type NilSubtestSuite struct {
	suite.Suite
	service *UserService
}

func (s *NilSubtestSuite) SetupTest() {
	s.service = NewUserService()
}

func (s *NilSubtestSuite) TestWithNilRun() {
	// Linter should not panic on nil function
	s.Run("nil subtest", nil)
}

func (s *NilSubtestSuite) TestWithVariableRun() {
	var fn func()
	// Linter should not panic on variable function
	s.Run("variable subtest", fn)
}

func (s *NilSubtestSuite) TestValidWithNilRun() {
	// Act
	result := SayHello("Josh")

	// Assert
	assert.Equal(s.T(), "Hello Josh", result)

	// Linter should not panic on nil function
	s.Run("nil subtest", nil)
}

func TestNilSubtestSuite(t *testing.T) {
	suite.Run(t, new(NilSubtestSuite))
}
