import { createGlobalStyle } from "styled-components";

const GlobalStyles = createGlobalStyle`
    html {
        /* A system font stack so things load nice and quick! */
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
        font-size: 16px;
    }
    body {
        padding: 0px;
        margin: 0px;
        height: 100%;
    }
    h1, h2, h3, h4, h5, p {
        margin: 0;
        padding: 0;
    }
`;

export default GlobalStyles;
