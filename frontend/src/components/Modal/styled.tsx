import { styled } from "styled-components";

export const Container = styled.div`
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    background-color: #00000099;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 888;
    backdrop-filter: blur(6px);
`

export const Content = styled.div`
    background-color: #FFF;
    min-width: 600px;
    min-height: 400px;
    position: relative;
    border-radius: 8px;
`

export const CloseButton = styled.button`
    border: none;
    background-color: transparent;
    position: absolute;
    right: 10px;
    top: 10px;
    cursor: pointer;
`