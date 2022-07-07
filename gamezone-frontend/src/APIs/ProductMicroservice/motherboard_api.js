const api = "http://localhost:7000/api/products/motherboards";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_SOCKETS = api + "/getSockets";
export const GET_PROCESSOR_TYPES = api + "/getProcessorTypes";
export const GET_FORM_FACTORS = api + "/getFormFactors";

const employeeProtectedMotherboards = "/employeeProtectedMotherboards";
export const CREATE = api + employeeProtectedMotherboards;
export const UPDATE = api + employeeProtectedMotherboards;
export const DELETE = api + employeeProtectedMotherboards;
