package parsing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/doctorwoot420/d209-armory-api/internal/domain"
)

// Parser obtains json from d2s109 parser then we'll convert to our domain model
type Parser struct {
	d2spath string
}

// Parse will provide a name to the d2s109 parser and receive json to convert into a character in our domain model.
func (p Parser) Parse(name string) (*domain.Character, error) {
	//TASK Consider adding support for multiple characters.  Better to do a single query to java app, it already accepts a list of names.
	d2s109Json, err := GetJsonFromJavaParser(name)
	if err != nil {
		return nil, fmt.Errorf("binary parse error: %w", err)
	}

	/*
		var i domain.Character

		if err := json.Unmarshal([]byte(d2s109Json), &i); err != nil {
			fmt.Println("ugh: ", err)
		}

		fmt.Println("info: ", i)
		fmt.Println("characterName: ", i.characterName)
	*/

	d2schar, err := ConvertJsonToD2sCharacter(d2s109Json)
	if err != nil {
		return nil, fmt.Errorf("binary parse error: %w", err)
	}

	/*
		// Character path on disk.
		file, err := os.Open(fmt.Sprintf("%s/%s", p.d2spath, name))
		if err != nil {
			return nil, fmt.Errorf("character binary does not exist: %w", domain.ErrNotFound)
		}

		// Close the file when we're done.
		defer file.Close()

		// Parse the actual .d2s binary file.
		d2schar, err := d2s.Parse(file)
	*/

	character := domain.Character{
		ID:         name,
		D2s:        d2schar,
		LastParsed: time.Now(),
	}

	return &character, nil
}

//Call to Chmiel's d2s109 parser with a character name and return json
func GetJsonFromJavaParser(characterName string) ([]byte, error) {
	resp, err := http.Get("http://localhost:8181/armory/api/armory?character=" + characterName)
	if err != nil {
		return nil, fmt.Errorf("GetJsonFromJavaParser error retrieving URL: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	//fmt.Println("string(body): ", string(body))
	return body, nil
}

//Take d2s109 json and convert into d2sMode/.Character struct
func ConvertJsonToD2sCharacter(d2s109Json []byte) (*domain.D2sCharacter, error) {
	f := domain.D2s109Character{}
	if err := json.Unmarshal(d2s109Json, &f); err != nil {
		panic(err)
	}
	fmt.Println("f: ", f)
	fmt.Println("f.CharacterName: ", f.CharacterName)
	char := domain.D2sCharacter{}

	/*
		if err := json.Unmarshal(d2s109Json, &char); err != nil {
			panic(err)
		}*/

	name, err := char.Header.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("ConvertJsonToD2sCharacter error : %w", err)
	}
	fmt.Println("name: ", name)

	return char, nil
}

// NewParser constructs a new parser with dependencies.
func NewParser() *Parser {
	return &Parser{}
}
