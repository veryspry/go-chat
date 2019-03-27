import axios from "axios";
import { store } from "../App";

export const requestConstructor = () => {
  const {
    auth: { user }
  } = store.getState();

  if (!user) {
    console.log("Error making request: no current autheticated user");
    return null;
  }

  const { ID, token } = user;

  return axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    headers: { UserID: ID, Authorization: `Bearer ${token}` }
  });
};

export default requestConstructor;
