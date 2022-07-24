const api = "http://localhost:7000/api/products/psus";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_POWERS = api + "/getPowers";
export const GET_TYPES = api + "/getTypes";
export const GET_FORM_FACTORS = api + "/getFormFactors";

const employeeProtectedPsus = "/employeeProtected";
export const CREATE = api + employeeProtectedPsus;
export const UPDATE = api + employeeProtectedPsus;
export const DELETE = api + employeeProtectedPsus;
