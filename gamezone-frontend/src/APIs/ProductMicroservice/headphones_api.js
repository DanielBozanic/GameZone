const api = "http://localhost:7000/api/products/headphones";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_CONNECTION_TYPES = api + "/getConnectionTypes";

const employeeProtectedHeadphones = "/employeeProtected";
export const CREATE = api + employeeProtectedHeadphones;
export const UPDATE = api + employeeProtectedHeadphones;
export const DELETE = api + employeeProtectedHeadphones;
