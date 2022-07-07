const api = "http://localhost:7000/api/products/ssds";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_CAPACITIES = api + "/getCapacities";
export const GET_FORMS = api + "/getForms";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_MAX_SEQUENTIAL_READS = api + "/getMaxSequentialReads";
export const GET_MAX_SEQUENTIAL_WRITES = api + "/getMaxSequentialWrites";

const employeeProtectedSsds = "/employeeProtectedSsds";
export const CREATE = api + employeeProtectedSsds;
export const UPDATE = api + employeeProtectedSsds;
export const DELETE = api + employeeProtectedSsds;
