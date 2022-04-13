package mapper

import (
	"product/dto"
	"product/model"
)


func ToKeyboard(keyboardDTO dto.KeyboardDTO) (model.Keyboard) {
	return model.Keyboard {
		Name: keyboardDTO.Name, 
		Price: keyboardDTO.Price,
		Wireless: keyboardDTO.Wireless,
		KeyboardConnector: keyboardDTO.KeyboardConnector,
		KeyType: keyboardDTO.KeyType,
		LetterLayout: keyboardDTO.LetterLayout,
		KeyboardColor: keyboardDTO.KeyboardColor,
		Amount: keyboardDTO.Amount,
		Manufacturer: keyboardDTO.Manufacturer,
	}
}

func ToKeyboardDTO(keyboard model.Keyboard) dto.KeyboardDTO {
	return dto.KeyboardDTO {
		Id: keyboard.Id, 
		Name: keyboard.Name, 
		Price: keyboard.Price,
		Wireless: keyboard.Wireless,
		KeyboardConnector: keyboard.KeyboardConnector,
		KeyType: keyboard.KeyType,
		LetterLayout: keyboard.LetterLayout,
		KeyboardColor: keyboard.KeyboardColor,
		Amount: keyboard.Amount,
		Manufacturer: keyboard.Manufacturer,
	}
}

func ToKeyboardDTOs(keyboards []model.Keyboard) []dto.KeyboardDTO {
	keyboardDTOs := make([]dto.KeyboardDTO, len(keyboards))

	for i, itm := range keyboards {
		keyboardDTOs[i] = ToKeyboardDTO(itm)
	}

	return keyboardDTOs
}