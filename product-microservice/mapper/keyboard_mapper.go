package mapper

import (
	"product/dto"
	"product/model"
)


func ToKeyboard(keyboardDTO dto.KeyboardDTO) (model.Keyboard) {
	return model.Keyboard {
		Product: model.Product(keyboardDTO.Product),
		Wireless: keyboardDTO.Wireless,
		KeyboardConnector: keyboardDTO.KeyboardConnector,
		KeyType: keyboardDTO.KeyType,
		LetterLayout: keyboardDTO.LetterLayout,
		KeyboardColor: keyboardDTO.KeyboardColor,
	}
}

func ToKeyboardDTO(keyboard model.Keyboard) dto.KeyboardDTO {
	return dto.KeyboardDTO {
		Product: dto.ProductDTO(keyboard.Product),
		Wireless: keyboard.Wireless,
		KeyboardConnector: keyboard.KeyboardConnector,
		KeyType: keyboard.KeyType,
		LetterLayout: keyboard.LetterLayout,
		KeyboardColor: keyboard.KeyboardColor,
	}
}

func ToKeyboardDTOs(keyboards []model.Keyboard) []dto.KeyboardDTO {
	keyboardDTOs := make([]dto.KeyboardDTO, len(keyboards))

	for i, itm := range keyboards {
		keyboardDTOs[i] = ToKeyboardDTO(itm)
	}

	return keyboardDTOs
}