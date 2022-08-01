const api = "http://localhost:7003/api/contactAndReport/reports";

const userAndEmployeeProtected = api + "/userAndEmployeeProtected";
const adminProtected = api + "/adminProtected";

export const ADD_REPORT = userAndEmployeeProtected + "/addReport";
export const GET_REPORTS_BY_USER_ID = adminProtected + "/getReportsByUserId";
