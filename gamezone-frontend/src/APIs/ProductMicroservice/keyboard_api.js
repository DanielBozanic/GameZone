const api = "http://localhost:7000/api/products/keyboards";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_KEYBOARD_CONNECTORS = api + "/getKeyboardConnectors";
export const GET_KEY_TYPES = api + "/getKeyTypes";

const employeeProtectedKeyboards = "/employeeProtectedKeyboards";
export const CREATE = api + employeeProtectedKeyboards;
export const UPDATE = api + employeeProtectedKeyboards;
export const DELETE = api + employeeProtectedKeyboards;
