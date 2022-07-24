const api = "http://localhost:7000/api/products/graphicsCards";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_CHIP_MANUFACTURERS = api + "/getChipManufacturers";
export const GET_MEMORY_SIZES = api + "/getMemorySizes";
export const GET_MEMORY_TYPES = api + "/getMemoryTypes";
export const GET_MODEL_NAMES = api + "/getModelNames";

const employeeProtectedGraphicsCards = "/employeeProtected";
export const CREATE = api + employeeProtectedGraphicsCards;
export const UPDATE = api + employeeProtectedGraphicsCards;
export const DELETE = api + employeeProtectedGraphicsCards;
