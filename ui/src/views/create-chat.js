import React from "react";
import { Flex, CreateChat } from "../components";

const CreateChatView = props => {
  return (
    <Flex
      width="100%"
      height="100%"
      alignItems="center"
      justifyContent="center"
      mt="30px"
    >
      <CreateChat />
    </Flex>
  );
};

export default CreateChatView;
