package apitools

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MemorablePasswordGenerator struct {
}

func (m *MemorablePasswordGenerator) Name() string {
    return "memorable-password-generator"
}

func (m *MemorablePasswordGenerator) Handle(c echo.Context) error {
	words := []string{
    "Mountain", "Ocean", "Sunshine", "Rainbow", "Galaxy", "Thunder", "Whisper", "Breeze", "Cascade", "Horizon",
    "Velvet", "Crimson", "Emerald", "Sapphire", "Golden", "Silver", "Crystal", "Marble", "Silk", "Amber",
    "Leopard", "Penguin", "Dolphin", "Phoenix", "Dragon", "Unicorn", "Panther", "Falcon", "Jaguar", "Octopus",
    "Violin", "Piano", "Guitar", "Trumpet", "Harmony", "Melody", "Rhythm", "Symphony", "Jazz", "Opera",
    "Quantum", "Nebula", "Comet", "Quasar", "Neutron", "Photon", "Cosmic", "Lunar", "Solar", "Stellar",
    "Waterfall", "Volcano", "Glacier", "Canyon", "Forest", "Desert", "Island", "Meadow", "Tundra", "Oasis",
    "Velocity", "Serenity", "Enigma", "Zenith", "Ethereal", "Cascade", "Radiant", "Luminous", "Tranquil", "Euphoria",
    "Quantum", "Cypher", "Matrix", "Vector", "Nexus", "Synergy", "Quantum", "Cipher", "Crypto", "Binary",
    "Saffron", "Cinnamon", "Ginger", "Paprika", "Wasabi", "Sesame", "Pepper", "Vanilla", "Cocoa", "Jasmine",
    "Dazzle", "Whisper", "Glimmer", "Ripple", "Sizzle", "Rustle", "Twinkle", "Murmur", "Glisten", "Flutter",
	}

	password := ""
	wordCount := 3
	if i, _ := strconv.Atoi(c.QueryParam("word-count")); i > 0 {
		count, _ := strconv.Atoi(c.QueryParam("word-count"))
		wordCount = count
	}

	for i := 0; i < wordCount; i++ {
        password += words[rand.Intn(len(words))]
    }
	
	password += strconv.Itoa(rand.Intn(99-20+1) + 20) // Add a random number between 20 and 99

    specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?";
	if c.QueryParam("special-character") == "on" {
		password += string(specialChars[rand.Intn(len(specialChars))])
	}
    return c.String(http.StatusOK, password)
}
