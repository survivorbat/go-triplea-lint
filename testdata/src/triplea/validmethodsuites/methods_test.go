package validmethods

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ValidSuite struct {
	suite.Suite
}

// No T

func (v *ValidSuite) TestSayHello() {
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayHello()

	// Assert
	assert.Equal(v.T(), "Hello Josh", actual)
}

func (v *ValidSuite) TestSayHello_ReturnsExpectedMessage() {
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayHello()

	// Assert
	assert.Equal(v.T(), "Hello Josh", actual)
}

// T

func (v *ValidSuite) TestSayGoodBye() {
	t := v.T()
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayGoodbye()

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

func (v *ValidSuite) TestSayGoodBye_ReturnsExpectedMessage() {
	t := v.T()
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayGoodbye()

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

// No arrange, no T

func (v *ValidSuite) TestSayGoodMorning() {
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodMorning()

	// Assert
	assert.Equal(v.T(), "Good morning Josh", actual)
}

func (v *ValidSuite) TestSayGoodMorning_ReturnsExpectedMessage() {
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodMorning()

	// Assert
	assert.Equal(v.T(), "Good morning Josh", actual)
}

// No arrange, T

func (v *ValidSuite) TestSayGoodAfternoon() {
	t := v.T()
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodAfternoon()

	// Assert
	assert.Equal(t, "Good afternoon Josh", actual)
}

func (v *ValidSuite) TestSayGoodAfternoon_ReturnsExpectedMessage() {
	t := v.T()
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodAfternoon()

	// Assert
	assert.Equal(t, "Good afternoon Josh", actual)
}

func TestValidSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ValidSuite))
}
