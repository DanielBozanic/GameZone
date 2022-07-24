const api = "http://localhost:7000/api/products/consoles";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_PLATFORMS = api + "/getPlatforms";

const employeeProtectedConsoles = "/employeeProtected";
export const CREATE = api + employeeProtectedConsoles;
export const UPDATE = api + employeeProtectedConsoles;
export const DELETE = api + employeeProtectedConsoles;
