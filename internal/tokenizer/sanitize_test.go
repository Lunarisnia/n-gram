package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sanitize(t *testing.T) {
	t.Run("Expect sanitized input properly", func(t *testing.T) {
		text := "Have you ever wondered why some symbols—like @, #, $, and %—appear so often in modern text? Well, in today's world of emails, usernames, and passwords, it's no surprise! Imagine typing an email: john.doe@example.com (yep, that's the @ symbol right there). Or, setting a new password: P@ssw0rd! (combining letters, numbers, and symbols like !, @, #, and $ for extra security). In fact, punctuation marks—such as parentheses ( ), curly braces { }, square brackets [ ], and even angle brackets < >—play critical roles in coding, making communication between machines and humans precise. But wait, there's more! Think about how we use commas, periods, colons, and semicolons: 'Where does it all end?' you might ask.\n\nWell, if you're coding, the answer is—sometimes—it doesn't end there. Take a deep dive into a complex script and you’ll encounter plenty of special characters: +-=/*~ might seem ordinary, but throw in some &^%$#@! symbols, and now it's a different game. These characters, when combined, help developers build logic and structure, ensuring programs work as expected. So next time you see a piece of code or a strong password like C0mpl3x&Str@nge! or even something seemingly random such as {x=[10*(2^3)]}, appreciate the intricate dance of letters, numbers, and punctuations coming together to make magic happen!"
		sanitized := Sanitize(text)
		assert.Equal(
			t,
			"Have you ever wondered why some symbols like and appear so often in modern text Well in today's world of emails usernames and passwords it's no surprise Imagine typing an email john doe example com yep that's the symbol right there Or setting a new password P ssw rd combining letters numbers and symbols like and for extra security In fact punctuation marks such as parentheses curly braces square brackets and even angle brackets play critical roles in coding making communication between machines and humans precise But wait there's more Think about how we use commas periods colons and semicolons 'Where does it all end ' you might ask Well if you're coding the answer is sometimes it doesn't end there Take a deep dive into a complex script and you ll encounter plenty of special characters might seem ordinary but throw in some symbols and now it's a different game These characters when combined help developers build logic and structure ensuring programs work as expected So next time you see a piece of code or a strong password like C mpl x Str nge or even something seemingly random such as x appreciate the intricate dance of letters numbers and punctuations coming together to make magic happen",
			sanitized,
		)
	})
}
