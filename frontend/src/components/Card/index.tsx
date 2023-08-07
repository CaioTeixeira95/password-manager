import { IPassword } from "../../types/password";
import { Content, Name, ContentTitle, Username } from "./styled";
import DeleteIcon from '@mui/icons-material/Delete';

interface IProps {
    password: IPassword,
    onClick: () => void;
    onDelete: () => void;
}

export const Card = ({ password, onClick, onDelete }: IProps): JSX.Element => {
    return (
        <Content onClick={onClick}>
            <ContentTitle>
                <Name>{password.name}</Name>
                <button onClick={(e) => {
                    e.stopPropagation();
                    onDelete();
                }}>
                    <DeleteIcon color="error" />
                </button>
            </ContentTitle>
            <Username>{password.username}</Username>
        </Content>
    )
}