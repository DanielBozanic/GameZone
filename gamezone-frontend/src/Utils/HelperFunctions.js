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
