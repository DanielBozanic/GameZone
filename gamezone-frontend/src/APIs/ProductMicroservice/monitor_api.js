const api = "http://localhost:7000/api/products/monitors";

export const GET_ALL = api;
export const GET_NUMBER_OF_RECORDS = api + "/getNumberOfRecords";
export const GET_BY_ID = api;
export const SEARCH_BY_NAME = api + "/searchByName";
export const GET_NUMBER_OF_RECORDS_SEARCH = api + "/getNumberOfRecordsSearch";
export const FILTER = api + "/filter";
export const GET_NUMBER_OF_RECORDS_FILTER = api + "/getNumberOfRecordsFilter";
export const GET_MANUFACTURERS = api + "/getManufacturers";
export const GET_ASPECT_RATIOS = api + "/getAspectRatios";
export const GET_RESOLUTIONS = api + "/getResolutions";
export const GET_REFRESH_RATES = api + "/getRefreshRates";

const employeeProtectedMonitors = "/employeeProtected";
export const CREATE = api + employeeProtectedMonitors;
export const UPDATE = api + employeeProtectedMonitors;
export const DELETE = api + employeeProtectedMonitors;
