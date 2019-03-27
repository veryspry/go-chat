export const SET_AUTH_TOKEN = "SET_AUTH_TOKEN";
export const SET_USER = "SET_USER";
export const LOGOUT = "LOGOUT";

/* setAuthToken 
  @param authToken: string
*/
export const setAuthToken = authToken => ({
  type: SET_AUTH_TOKEN,
  authToken
});

/* setUser
  @param user: object
*/
export const setUser = user => ({
  type: SET_USER,
  user
});

/* logout
  @param none
*/
export const logout = () => ({
  type: LOGOUT
});
