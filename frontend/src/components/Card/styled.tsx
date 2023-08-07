import styled from "styled-components";

export const Content = styled.div`
    padding: 10px;
    border-radius: 8px;
    background-color: #FFF;
    -webkit-box-shadow: 2px 4px 8px 0px rgba(0,0,0,0.50);
    -moz-box-shadow: 2px 4px 8px 0px rgba(0,0,0,0.50);
    box-shadow: 2px 4px 8px 0px rgba(0,0,0,0.50);
    width: 200px;
    height: 150px;
    display: flex;
    flex-direction: column;
    gap: 10px;
`

export const Name = styled.div`
    max-width: 140px;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;  
    overflow: hidden;
    font-weight: bold;
`

export const Username = styled.span`
    max-width: 140px;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    font-weight: 400;
    opacity: 0.6;
`

export const ContentTitle = styled.div`
    display: flex;
    justify-content: space-between;
`
