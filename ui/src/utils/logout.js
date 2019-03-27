import { store } from "../App";
import { logout as logoutAction } from "../redux/actions";

export const logout = () => {
  const { auth } = store.getState();
  store.dispatch(logoutAction());
};

export default logout;
