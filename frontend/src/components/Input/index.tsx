import React from "react";
import { Content, Float, InputError, Input as InputStyled } from "./styled";

interface IProps {
    value: string;
    onChange: (value: string) => void;
    renderRight?: JSX.Element; 
    type?: string;
    placeholder?: string;
    border?: React.CSSProperties["border"];
    color?: React.CSSProperties["color"];
    width?: React.CSSProperties["width"];
    error?: string;
}


export const Input = ({
    type = "text",
    placeholder,
    border,
    color,
    width,
    value,
    onChange,
    error,
    renderRight,
}: IProps): JSX.Element => {
    return (
        <Content width={width}>
            <InputStyled
                type={type}
                placeholder={placeholder}
                value={value}
                onChange={(e) => onChange(e.target.value)}
                border={border}
                color={color}
                width={width}
                error={error} />
            {error && <InputError>{error}</InputError>}
            {renderRight && <Float>{renderRight}</Float>}
        </Content>
    )
}
