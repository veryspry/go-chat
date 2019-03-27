import React from "react";
import {
  faTwitter,
  faGithub,
  faInstagram
} from "@fortawesome/free-brands-svg-icons";
import {
  Flex,
  Box,
  StyledLink,
  StyledAnchor,
  HeaderText,
  Img,
  Icon,
  CreateChat
} from "../components";

import { logout, requestConstructor, getCurrentUser } from "../utils";

const Header = props => {
  const navItems = [
    {
      title: "All Chats",
      to: "/chat"
    },
    {
      title: "New Chat",
      to: "/chat/new"
    }
  ];

  const currentUser = getCurrentUser();

  const authAction = {
    text: "Logout",
    to: "/login",
    onClick: logout
  };

  if (!currentUser) {
    authAction.text = "Login";
    authAction.onClick = null;
  }

  return (
    <Flex
      flexDirection={["column-reverse", "column-reverse", "row"]}
      justifyContent={["center", "center", "space-around"]}
      alignItems="center"
      bg="lightpink"
    >
      <Flex
        flexDirection={["column", "column", "row"]}
        alignItems="center"
        textAlign="center"
      >
        <Flex
          flexDirection="row"
          my={["20px", "20px", "0px"]}
          zIndex="4000"
          justifyContent="space-between"
          py="40px"
        >
          {currentUser && currentUser.firstName && (
            <HeaderText>Hey, {currentUser.firstName}! 💬 </HeaderText>
          )}

          <Flex flexDirection="row">
            {navItems.map(({ title, to, onClick }) => {
              return (
                <Box mx="10px" key={title}>
                  <StyledLink
                    to={to}
                    fontWeight="100"
                    color="black"
                    hovercolor="#2096c7"
                    textDecoration="underline"
                    onClick={onClick}
                  >
                    <HeaderText>{title}</HeaderText>
                  </StyledLink>
                </Box>
              );
            })}

            <Box mx="10px">
              <StyledLink
                to={authAction.to}
                fontWeight="100"
                color="black"
                hovercolor="#2096c7"
                textDecoration="underline"
                onClick={authAction.onClick}
              >
                <HeaderText>{authAction.text}</HeaderText>
              </StyledLink>
            </Box>
          </Flex>
        </Flex>
      </Flex>
    </Flex>
  );
};

export default Header;
