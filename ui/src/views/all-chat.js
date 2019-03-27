import React from "react";

import { Flex, ChatList } from "../components";

const AllChat = props => {
  return (
    <Flex alignItems="center" py="20px">
      <ChatList />
    </Flex>
  );
};

export default AllChat;
