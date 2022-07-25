const api = "http://localhost:7002/api/news/newsSubscriptions";

const userProtected = api + "/userProtected";

export const SUBSCRIBE = userProtected + "/subscribe";
export const UNSUBSCRIBE = userProtected + "/unsubscribe";
export const IS_USER_SUBSCRIBED = userProtected + "/isUserSubscribed";
