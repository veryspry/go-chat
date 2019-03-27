import styled from "styled-components";
import {
  space,
  width,
  height,
  color,
  position,
  alignContent,
  alignItems,
  justifyContent,
  borders,
  borderRadius,
  textAlign,
  flexDirection,
  zIndex,
  top,
  bottom,
  right,
  left,
  overflow,
  minHeight
} from "styled-system";

export const Flex = styled.div`
  display: flex;
  flex-direction: ${({ flexDirection }) =>
    flexDirection ? flexDirection : "column"};
  ${flexDirection}
  ${space};
  ${width};
  ${color};
  ${height};
  ${minHeight};
  ${borders};
  ${position};
  ${top};
  ${bottom};
  ${right};
  ${left};
  ${alignContent};
  ${alignItems};
  ${justifyContent};
  ${borderRadius};
  ${textAlign};
  ${zIndex};
  ${overflow};
  transform: ${({ rotate }) => (rotate ? `rotate(${rotate})` : null)};
`;

export const Box = styled.div`
    ${space} ${width} ${color};
`;

export const Img = styled.img`
  ${width};
  ${height};
  ${borderRadius};
  ${position};
  ${space};
`;
