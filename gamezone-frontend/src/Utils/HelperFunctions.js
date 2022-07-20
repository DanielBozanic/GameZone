import * as productType from "./ProductType";

export const str2Bool = (value) => {
	if (value && typeof value === "string") {
		if (value.toLowerCase() === "true") return true;
		if (value.toLowerCase() === "false") return false;
	}
	return value;
};

export const merge = (...schemas) => {
	const [first, ...rest] = schemas;
	const merged = rest.reduce(
		(mergedSchemas, schema) => mergedSchemas.concat(schema),
		first
	);
	return merged;
};

export const getUpdateProductRoute = (product) => {
	switch (product.Product.Type) {
		case productType.CONSOLE:
			return "/updateConsole/" + product.Product.Id;
		case productType.GRAPHICS_CARD:
			return "/updateGraphicsCard/" + product.Product.Id;
		case productType.HARD_DISK_DRIVE:
			return "/updateHdd/" + product.Product.Id;
		case productType.HEADPHONES:
			return "/updateHeadphones/" + product.Product.Id;
		case productType.KEYBOARD:
			return "/updateKeyboard/" + product.Product.Id;
		case productType.MONITOR:
			return "/updateMonitor/" + product.Product.Id;
		case productType.MOTHERBOARD:
			return "/updateMotherboard/" + product.Product.Id;
		case productType.MOUSE:
			return "/updateMouse/" + product.Product.Id;
		case productType.POWER_SUPPLY_UNIT:
			return "/updatePsu/" + product.Product.Id;
		case productType.PROCESSOR:
			return "/updateProcessor/" + product.Product.Id;
		case productType.RAM:
			return "/updateRam/" + product.Product.Id;
		case productType.SOLID_STATE_DRIVE:
			return "/updateSsd/" + product.Product.Id;
		case productType.VIDEO_GAME:
			return "/updateVideoGame/" + product.Product.Id;
	}
};

export const getProductDetailRoute = (product) => {
	switch (product.Type) {
		case productType.CONSOLE:
			return "/consoles/" + product.Id;
		case productType.GRAPHICS_CARD:
			return "/graphicsCards/" + product.Id;
		case productType.HARD_DISK_DRIVE:
			return "/hdds/" + product.Id;
		case productType.HEADPHONES:
			return "/headphones/" + product.Id;
		case productType.KEYBOARD:
			return "/keyboards/" + product.Id;
		case productType.MONITOR:
			return "/monitors/" + product.Id;
		case productType.MOTHERBOARD:
			return "/motherboards/" + product.Id;
		case productType.MOUSE:
			return "/mice/" + product.Id;
		case productType.POWER_SUPPLY_UNIT:
			return "/psus/" + product.Id;
		case productType.PROCESSOR:
			return "/processors/" + product.Id;
		case productType.RAM:
			return "/rams/" + product.Id;
		case productType.SOLID_STATE_DRIVE:
			return "/ssds/" + product.Id;
		case productType.VIDEO_GAME:
			return "/videoGames/" + product.Id;
	}
};

export const getProductListRoute = (product) => {
	switch (product.Product.Type) {
		case productType.CONSOLE:
			return "/consoles";
		case productType.GRAPHICS_CARD:
			return "/graphicsCards";
		case productType.HARD_DISK_DRIVE:
			return "/hdds";
		case productType.HEADPHONES:
			return "/headphones";
		case productType.KEYBOARD:
			return "/keyboards";
		case productType.MONITOR:
			return "/monitors";
		case productType.MOTHERBOARD:
			return "/motherboards";
		case productType.MOUSE:
			return "/mice";
		case productType.POWER_SUPPLY_UNIT:
			return "/psus";
		case productType.PROCESSOR:
			return "/processors";
		case productType.RAM:
			return "/rams";
		case productType.SOLID_STATE_DRIVE:
			return "/ssds";
		case productType.VIDEO_GAME:
			return "/videoGames";
	}
};
