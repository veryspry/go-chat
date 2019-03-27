import { Link } from "react-router-dom";
import styled from "styled-components";
import { color, fontSize, fontWeight, zIndex } from "styled-system";

export const HeaderText = styled.h1`
  ${color}
  font-weight: 200;
  font-size: 2rem;
  z-index: 4000;
  ${zIndex};
  ${fontSize};
`;

export const BodyHome = styled.h3`
  ${color}
  font-weight: 200;
  font-size: 2.5rem;
  ${fontSize};
  z-index: 4000;
  ${zIndex};
`;

export const FooterText = styled.h3`
  ${color};
  font-weight: 200;
  font-size: 1.4rem;
  z-index: 4000;
  ${zIndex};
  ${fontSize};
`;

export const TimelineText = styled.h1`
  ${color}
  font-weight: 300;
  font-size: 1.2rem;
  z-index: 4000;
  ${zIndex};
  ${fontSize};
`;

export const TimelineDate = styled.h1`
  ${color};
  font-weight: 700;
  font-size: 0.9rem;
  z-index: 4000;
  ${zIndex};
  ${fontSize};
`;

export const StyledLink = styled(Link)`
  ${color};
  text-decoration: none;
  ${fontWeight}
  ${fontSize};
  &:visited {
    color: ${({ color }) => (color ? color : "black")};
  }
  &:hover {
    text-decoration: ${({ textDecoration }) =>
      textDecoration ? textDecoration : "underline"};
    color: ${({ hovercolor }) => (hovercolor ? hovercolor : "black")};
  }
`;

export const StyledAnchor = styled.a`
  ${color}
  text-decoration: none;
  ${fontWeight}
  ${fontSize};
  &:hover {
    cursor: pointer;
    text-decoration: underline;
  }
`;
