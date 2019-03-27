import styled from "styled-components";
import { height, width, minHeight } from "styled-system";

export const Form = styled.form`
  ${height};
  ${width};
  display: flex;
  flex-direction: column;
  z-index: 1000;
`;

export const Input = styled.input`
  width: 100%;
  height: 30px;
  border-radius: 3px;
  border: 1px solid lightgrey;
  ${height};
  ${width};
`;

export const TextArea = styled.textarea`
  width: 100%;
  max-width: 100%;
  min-width: 100%;
  border-radius: 3px;
  border: 1px solid lightgrey;
  ${height};
  ${width};
  ${minHeight};
`;

export const Button = styled.button`
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
