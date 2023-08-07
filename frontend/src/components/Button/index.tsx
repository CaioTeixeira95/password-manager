import { Button as ButtonStyled } from "./styled"

interface IProps {
    label: string;
    onClick: () => void;
    type?: "button" | "submit" | "reset";
}

export const Button = ({ label, onClick, type = "button" }: IProps): JSX.Element => {
    return (
        <ButtonStyled
            type={type}
            onClick={onClick} >
            {label}
        </ButtonStyled>
    )
}
