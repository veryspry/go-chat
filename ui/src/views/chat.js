import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import { connect } from "react-redux";
import styled from "styled-components";
import { height, width, minHeight } from "styled-system";
import { Flex, FooterText } from "../components";
import requestConstructor from "../utils/request";

import { setAuthToken } from "../redux/actions";

import ChatList from "../components";

const Form = styled.form`
  ${height};
  ${width};
  display: flex;
  flex-direction: column;
  z-index: 1000;
`;

const Input = styled.input`
  width: 100%;
  height: 30px;
  border-radius: 3px;
  border: 1px solid lightgrey;
  ${height};
  ${width};
`;

const TextArea = styled.textarea`
  width: 100%;
  max-width: 100%;
  min-width: 100%;
  border-radius: 3px;
  border: 1px solid lightgrey;
  ${height};
  ${width};
  ${minHeight};
`;

const Button = styled.button`
  height: 60px;
  width: 200px;
  border-radius: 3px;
  border: 0px;
  background-color: #181e2f;
  color: white;
  &:hover {
    cursor: pointer;
  }
`;

const InputWrap = styled(Flex)`
  margin-top: 10px;
  margin-bottom: 10px;
`;

class Chat extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: "",
      messages: []
    };
  }

  async componentDidMount() {
    const {
      match: {
        params: { roomID }
      },
      user: { ID: userID }
    } = this.props;

    const requestor = requestConstructor();

    let messages = requestor
      .get(`/chat/conversations/${roomID}/messages`)
      .then(res => {
        if (res.data.messages) {
          this.setState({
            messages: res.data.messages.Value
          });
        }
      })
      .catch(err => console.log(err));

    this.ws = new WebSocket(
      `${process.env.REACT_WEBSOCKET_BASE_URL}/${roomID}?userID=${userID}`
    );
    this.ws.addEventListener("message", this._handleNewMessage);
  }

  componentWillUnmount() {
    this.ws.removeEventListener("message", this._handleNewMessage);
  }

  _handleNewMessage = async event => {
    let msg = JSON.parse(event.data);
    this._addNewMessageToState(msg);
  };

  _addNewMessageToState = async msg => {
    let msgCopy = this.state.messages.slice();
    msgCopy.push(msg);
    await this.setState({
      messages: msgCopy
    });
  };

  _handleChange = event => {
    const { name, value } = event.target;
    this.setState({
      [name]: value
    });
  };

  _handleSubmit = async event => {
    event.preventDefault();
    const { message } = this.state;
    const {
      user: { ID: userID }
    } = this.props;

    const formattedMsg = {
      message,
      userID
    };

    this.ws.send(JSON.stringify(formattedMsg));
    this._addNewMessageToState(formattedMsg);
    await this.setState({
      message: ""
    });
  };

  render() {
    const { ID: currUsrID } = this.props.user;
    return (
      <Flex flexDirection="column" alignItems="center" my="70px">
        <Form width={["95vw", "80vw", "600px"]}>
          <Flex>
            {this.state.messages.map(({ message, userID }, i) => {
              let bgColor = "lightblue";
              let align = "flex-start";
              if (userID === currUsrID) {
                bgColor = "lightgrey";
                align = "flex-end";
              }
              return (
                <Flex key={i} width="100%" alignItems={align}>
                  <Flex bg={bgColor} width="50%">
                    <FooterText>{message}</FooterText>
                  </Flex>
                </Flex>
              );
            })}
          </Flex>
          <InputWrap>
            <FooterText>Message:</FooterText>
            <Input
              type="text"
              name="message"
              value={this.state.message}
              onChange={this._handleChange}
            />
          </InputWrap>
          <InputWrap alignItems="flex-end">
            <Button type="submit" value="submit" onClick={this._handleSubmit}>
              <FooterText>Send</FooterText>
            </Button>
          </InputWrap>
        </Form>
      </Flex>
    );
  }
}

const mapStateToProps = ({ auth: { user } }, ownProps) => {
  return {
    user: user
  };
};

export default connect(
  mapStateToProps,
  null
)(withRouter(Chat));
