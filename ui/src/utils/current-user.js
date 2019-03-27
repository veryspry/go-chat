import { store } from "../App";

// getCurrentUser returns the current user, or null if it doesn't exist
export const getCurrentUser = () => {
  const { auth } = store.getState();
  if (!auth) return null;
  const { user } = auth;
  if (!user) return null;
  return user;
};

export default getCurrentUser;
