import React from "react";
import { Flex, BodyHome } from "./";

const Loading = props => {
  return (
    <Flex justifyContent="center" alignItems="center" height="60vh">
      <BodyHome> Loading...</BodyHome>
    </Flex>
  );
};

export default Loading;
