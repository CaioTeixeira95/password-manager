import { CloseButton, Container, Content } from "./styled";
import CloseIcon from '@mui/icons-material/Close';

interface IProps {
    isOpen: boolean;
    onClose: () => void;
    children: JSX.Element;
}

export const Modal = ({ isOpen, onClose, children }: IProps): JSX.Element => {
    if (!isOpen) return <></>
    return (
        <Container>
            <Content>
                <CloseButton onClick={onClose}>
                    <CloseIcon />
                </CloseButton>
                {children}
            </Content>
        </Container>
    )
}