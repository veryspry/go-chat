import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import styled from "styled-components";

import {
  Flex,
  Form,
  Input,
  TextArea,
  Button,
  FooterText,
  Footer
} from "../../components";

import { requestConstructor } from "../../utils";

class CreateChat extends Component {
  constructor(props) {
    super(props);
    this.state = {
      currentFormInput: "",
      errorMessage: "",
      users: []
    };
    this.requestor = requestConstructor();
  }

  _handleChange = event => {
    const { name, value } = event.target;
    this.setState({
      [name]: value
    });
  };

  _lookupUser = event => {
    let { users, currentFormInput } = this.state;
    event.preventDefault();
    this.requestor
      .get("/user", {
        params: {
          email: currentFormInput
        }
      })
      .then(res => {
        let { data } = res;
        if (data.message) {
          if (data.message === "User not found") {
            this.setState({
              errorMessage: "User not found"
            });
            return;
          }
        }
        this.setState({
          users: [...users, res.data.user],
          currentFormInput: ""
        });
      })
      .catch(err => console.log(err));
  };

  _createConversation = event => {
    event.preventDefault();
    let userIDs = this.state.users.map(user => user.ID);
    this.requestor
      .post("/chat/conversations/new", {
        userIDs
      })
      .then(res => {
        let { conversation } = res.data;
        if (conversation) {
          this.props.history.push(`/chat/${conversation.ID}`);
        }
      })
      .catch(err => console.log(err));
  };

  _removeUserFromList = event => {
    let id = event.target.getAttribute("id");
    this.setState({
      users: this.state.users.filter(user => user.ID != id)
    });
  };

  render() {
    const { users, currentFormInput, errorMessage } = this.state;
    return (
      <Flex>
        {users.length > 0 && (
          <FooterText>Current selected users to start chat with:</FooterText>
        )}
        {users.map(user => {
          return (
            <Flex key={user.ID} flexDirection="row">
              <Flex mr="10px">
                <FooterText>{user.email}</FooterText>
              </Flex>
              <Flex onClick={this._removeUserFromList}>
                <span id={user.ID}> Remove (x)</span>
              </Flex>
            </Flex>
          );
        })}
        <Form>
          <Flex my="10px">
            <FooterText>Type in a users email to find them</FooterText>
          </Flex>
          <Flex my="10px">
            <Input
              type="text"
              name="currentFormInput"
              value={this.state.currentFormInput}
              onChange={this._handleChange}
            />
          </Flex>
          {errorMessage.length > 0 && (
            <Flex my="10px">
              {" "}
              <FooterText>{errorMessage}</FooterText>
            </Flex>
          )}
          <Flex my="10px">
            <Button onClick={this._lookupUser}>
              <FooterText>Find user</FooterText>
            </Button>
          </Flex>
          <Flex my="10px">
            <Button>
              <FooterText onClick={this._createConversation}>
                Create Conversation
              </FooterText>
            </Button>
          </Flex>
        </Form>
      </Flex>
    );
  }
}

export default withRouter(CreateChat);
