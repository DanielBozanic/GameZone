const api = "http://localhost:7000/api/products/mouses";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_DPIS = api + "/getDpis";
export const GET_CONNECTIONS = api + "/getConnections";

const employeeProtectedMouses = "/employeeProtectedMouses";
export const CREATE = api + employeeProtectedMouses;
export const UPDATE = api + employeeProtectedMouses;
export const DELETE = api + employeeProtectedMouses;
