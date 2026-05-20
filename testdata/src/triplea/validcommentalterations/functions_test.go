package validcommentalterations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello_ExtraPunctuation(t *testing.T) {
	// Arrange!
	name := "Josh"

	// Act?
	actual := SayHello(name)

	// Assert<3
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_ExtraText(t *testing.T) {
	// Arrange is very cool
	name := "Josh"

	// Act is cooler
	actual := SayHello(name)

	// Assert is the coolest
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_ExtraWhitespace(t *testing.T) {
	// 								Arrange
	name := "Josh"

	//								Act
	actual := SayHello(name)

	//                Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_BlockComments(t *testing.T) {
	/* Arrange */
	name := "Josh"

	/* Act */
	actual := SayHello(name)

	/* Assert */
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_MultipleCommentSlashes(t *testing.T) {
	//// Arrange
	name := "Josh"

	/// Act
	actual := SayHello(name)

	/// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_RandomStars(t *testing.T) {
	//**** Arrange
	name := "Josh"

	///**. Act
	actual := SayHello(name)

	///*/**/* Assert
	assert.Equal(t, "Hello Josh", actual)
}
