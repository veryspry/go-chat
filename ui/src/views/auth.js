import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import { connect } from "react-redux";
import axios from "axios";
import styled from "styled-components";
import {
  Flex,
  FooterText,
  Form,
  Input,
  TextArea,
  Button,
  StyledLink,
  TimelineDate
} from "../components";

import { setUser } from "../redux/actions";

const InputWrap = styled(Flex)`
  margin-top: 10px;
  margin-bottom: 10px;
`;

class Auth extends Component {
  constructor(props) {
    super(props);
    this.state = {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      errorMsg: ""
    };
  }

  _handleChange = event => {
    const { name, value } = event.target;
    this.setState({
      [name]: value
    });
  };

  _handleSubmit = async event => {
    event.preventDefault();

    let { config } = this.props;
    let { apiPath, fields } = config;

    const fieldNames = fields.map(field => (field = field.name));

    let requestBody = {};

    fieldNames.forEach(name => {
      return (requestBody[name] = this.state[name]);
    });

    axios
      .post(`${process.env.REACT_APP_API_URL}${apiPath}`, requestBody)
      .then(res => {
        const { history, setUser } = this.props;
        const {
          data: { isAuthenticated, message, user }
        } = res;

        if (message === "Email address not found") {
          this.setState({
            errorMsg: "Email or password does not exist"
          });
        }

        if (isAuthenticated) {
          setUser(user);
          history.push("/chat");
        }
      })
      .catch(err => console.log(err));
  };

  render() {
    let { config } = this.props;
    let { buttonText, action, fields } = config;

    return (
      <Flex flexDirection="column" alignItems="center" my="70px">
        <FooterText color="red">{this.state.errorMsg}</FooterText>
        <Form width={["95vw", "80vw", "600px"]}>
          {fields.map(({ title, name }) => {
            return (
              <InputWrap key={name}>
                <FooterText>{title}</FooterText>
                <Input
                  type="text"
                  name={name}
                  value={this.state[name]}
                  onChange={this._handleChange}
                />
              </InputWrap>
            );
          })}
          <InputWrap flexDirection="row" justifyContent="space-between">
            <StyledLink to={action.path}>
              {" "}
              <TimelineDate>{action.text}</TimelineDate>{" "}
            </StyledLink>
            <Button type="submit" value="submit" onClick={this._handleSubmit}>
              <FooterText>{buttonText}</FooterText>
            </Button>
          </InputWrap>
        </Form>
      </Flex>
    );
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    ...state
  };
};

const mapDispatchToProps = (dispatch, ownProps) => ({
  setUser: user => dispatch(setUser(user))
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withRouter(Auth));
