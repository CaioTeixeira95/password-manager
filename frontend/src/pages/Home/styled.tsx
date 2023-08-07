import { styled } from "styled-components";

export const CardsWrapper = styled.div`
    padding: 100px 30px 30px 30px;
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    overflow: auto;
    max-height: 100%;
`

export const FABButton = styled.button`
    position: absolute;
    right: 30px;
    bottom: 30px;
    width: 60px;
    height: 60px;
    border-radius: 50%;
    background-color: #2C2C2C;
    color: #FFF;
    border: none;
    -webkit-box-shadow: 2px 4px 8px 0px rgba(0,0,0,0.50);
    -moz-box-shadow: 2px 4px 8px 0px rgba(0,0,0,0.50);
    box-shadow: 2px 4px 8px 0px rgba(0,0,0,0.50);
    cursor: pointer;
`
