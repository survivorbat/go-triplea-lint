package validmethods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Shortened name

func TestSayHello(t *testing.T) {
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayHello()

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_ReturnsExpectedMessage(t *testing.T) {
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayHello()

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

// No parallel

func TestMyService_SayHello(t *testing.T) {
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayHello()

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestMyService_SayHello_ReturnsExpectedMessage(t *testing.T) {
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayHello()

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

// Parallel

func TestMyService_SayGoodbye(t *testing.T) {
	t.Parallel()
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayGoodbye()

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

func TestMyService_SayGoodbye_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	// Arrange
	svc := &MyService{Name: "Josh"}

	// Act
	actual := svc.SayGoodbye()

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

// No arrange, no parallel

func TestMyService_SayGoodMorning(t *testing.T) {
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodMorning()

	// Assert
	assert.Equal(t, "Good morning Josh", actual)
}

func TestMyService_SayGoodMorning_ReturnsExpectedMessage(t *testing.T) {
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodMorning()

	// Assert
	assert.Equal(t, "Good morning Josh", actual)
}

// No arrange, parallel

func TestMyService_SayGoodAfternoon(t *testing.T) {
	t.Parallel()
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodAfternoon()

	// Assert
	assert.Equal(t, "Good afternoon Josh", actual)
}

func TestMyService_SayGoodAfternoon_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	// Act
	actual := (&MyService{Name: "Josh"}).SayGoodAfternoon()

	// Assert
	assert.Equal(t, "Good afternoon Josh", actual)
}
