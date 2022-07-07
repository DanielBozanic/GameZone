const api = "http://localhost:7000/api/products/processors";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_SOCKETS = api + "/getSockets";
export const GET_TYPES = api + "/getTypes";
export const GET_THREADS = api + "/getThreads";
export const GET_NUMBER_OF_CORES = api + "/getNumberOfCores";

const employeeProtectedProcessors = "/employeeProtectedProcessors";
export const CREATE = api + employeeProtectedProcessors;
export const UPDATE = api + employeeProtectedProcessors;
export const DELETE = api + employeeProtectedProcessors;
