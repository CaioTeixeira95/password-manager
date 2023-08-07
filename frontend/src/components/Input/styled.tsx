import styled, { css } from "styled-components";

interface IInput {
    border?: React.CSSProperties["border"];
    color?: React.CSSProperties["color"];
    width?: React.CSSProperties["width"];
    error?: string;
}

interface IContent {
    width?: React.CSSProperties["width"];
}

export const Input = styled.input<IInput>`
    background-color: #FFFFFF55;
    border-radius: 8px;
    border: ${({ border }) => border ?? "1px solid #2C2C2C"};
    ${({ error }) => error && css`border-color: red;`};
    height: 40px;
    width: ${({ width }) => width ?? "100%"};
    color: ${({ color }) => color ?? "#2C2C2C"};
    padding: 8px;
    font-size: 16px;

    &:focus {
        outline: none;
    }

    &::placeholder {
        color: ${({ color }) => color ?? "#2C2C2C99"};
    }
`

export const InputError = styled.span`
    color: red;
`

export const Content = styled.div<IContent>`
    width: ${({ width }) => width ?? "100%"};
    position: relative;
`

export const Float = styled.div`
    position: absolute;
    right: 10px;
    top: 10px;
`
