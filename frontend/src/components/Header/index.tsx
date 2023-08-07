import { Header as HeaderStyled } from "./styled"
import { Input } from "../Input"
import SearchIcon from '@mui/icons-material/Search';

interface IProps {
    searchedValue: string;
    changeSearchedValue: (value: string) => void;
}

export const Header = ({ searchedValue, changeSearchedValue }: IProps): JSX.Element => {
    return (
        <HeaderStyled>
            <span>Password Manager</span>
            <Input
                type="text"
                placeholder="Search"
                value={searchedValue}
                onChange={changeSearchedValue}
                border="none"
                color="#FFF"
                width="200px"
                renderRight={<SearchIcon />} />
        </HeaderStyled>
    )
}
