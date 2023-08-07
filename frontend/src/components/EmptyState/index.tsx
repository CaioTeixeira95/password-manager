import { Content } from "./styled";

interface IProps {
    message: string;
}

export const EmptyState = ({ message }: IProps): JSX.Element => {
    return (
        <Content>{message}</Content>
    );
}