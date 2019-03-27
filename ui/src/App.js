import React, { Component, Fragment } from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import { ThemeProvider } from "styled-components";
import { createStore } from "redux";
import { Provider } from "react-redux";
import { persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage"; // defaults to localStorage for web
import { PersistGate } from "redux-persist/integration/react";

import configureStore from "./redux/configureStore";

import { Auth, Home, Chat, AllChat, NotFound, CreateChatView } from "./views";
import rootReducer from "./redux/reducers";
import DefaultRoute from "./default-route";

import GlobalStyles from "./styles/global";

import theme from "./theme";

let { store, persistor } = configureStore();
// Export the store object for use elsewhere
export { store };

class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <PersistGate loading={<div>loading...</div>} persistor={persistor}>
          {/* Global Style reset */}
          <GlobalStyles />
          <ThemeProvider theme={theme}>
            <BrowserRouter>
              <Switch>
                <DefaultRoute
                  exact
                  path="/login"
                  component={Auth}
                  config={{
                    apiPath: "/login",
                    buttonText: "Login",
                    action: {
                      text: "Don't have an account? Create one here!",
                      path: "/user/new"
                    },
                    fields: [
                      {
                        title: "Email:",
                        name: "email"
                      },
                      {
                        title: "Password",
                        name: "password"
                      }
                    ]
                  }}
                />
                <DefaultRoute
                  exact
                  path="/user/new"
                  component={Auth}
                  config={{
                    apiPath: "/user/new",
                    buttonText: "Create User",
                    action: {
                      text: "Already have an account? Login here!",
                      path: "/login"
                    },
                    fields: [
                      {
                        title: "First Name:",
                        name: "firstName"
                      },
                      {
                        title: "Last Name:",
                        name: "lastName"
                      },
                      {
                        title: "Email:",
                        name: "email"
                      },
                      {
                        title: "Password",
                        name: "password"
                      }
                    ]
                  }}
                />
                <DefaultRoute
                  exact
                  path="/chat"
                  component={AllChat}
                  isAuthenticated
                />
                <DefaultRoute
                  exact
                  path="/chat/new"
                  component={CreateChatView}
                  isAuthenticated
                />
                <DefaultRoute
                  exact
                  path="/chat/:roomID([0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12})"
                  component={Chat}
                  isAuthenticated
                />
                <DefaultRoute path="/" component={NotFound} />
              </Switch>
            </BrowserRouter>
          </ThemeProvider>
        </PersistGate>
      </Provider>
    );
  }
}

export default App;
