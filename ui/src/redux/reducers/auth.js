import { SET_AUTH_TOKEN, SET_USER, LOGOUT } from "../actions";

const auth = (state = {}, action) => {
  switch (action.type) {
    case SET_AUTH_TOKEN:
      return Object.assign({}, state, { authToken: action.authToken });
    case SET_USER:
      return Object.assign({}, state, { user: action.user });
    case LOGOUT:
      return Object.assign({}, state, { user: null });
    default:
      return state;
  }
};

export default auth;
