import React from "react";

import { FooterText, Flex, StyledAnchor } from "./index";

const Footer = props => {
  return (
    <Flex
      alignItems="center"
      py="40px"
      px="20px"
      textAlign="center"
      position="fixed"
      bottom={0}
      width="100%"
    >
      <FooterText zIndex="4000">
        This site is built with Go, React, Redux, Postgres &{" "}
        <span role="img" aria-label="heart-emoji">
          ❤️
        </span>{" "}
        <br />
        View the ui source{" "}
        <StyledAnchor color="link" href="https://github.com/veryspry/js-chat">
          here
        </StyledAnchor>
        and the the server code{" "}
        <StyledAnchor color="link" href="https://github.com/veryspry/go-chat">
          here
        </StyledAnchor>
      </FooterText>
    </Flex>
  );
};

export default Footer;
