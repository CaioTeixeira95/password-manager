import { useEffect, useState } from "react";
import { Input } from "../Input";
import { Modal } from "../Modal";
import { Container, InputGroup, PasswordIconContainer } from "./styled";
import { Button } from "../Button";
import { handleValidate } from "../../helpers/utils";
import api from "../../api"
import uuid from "react-uuid";
import { IPassword } from "../../types/password";
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';

interface IProps {
    isOpen: boolean;
    onClose: () => void;
    handleUpdatePasswordCards: (password: IPassword) => void;
    flow: "create" | "edit";
    passwordSelected?: IPassword | null;
}

export const PasswordFormModal = ({
    isOpen,
    onClose,
    handleUpdatePasswordCards,
    flow,
    passwordSelected,
}: IProps): JSX.Element => {
    const [url, setURL] = useState("");
    const [name, setName] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showPassword, setShowPassword] = useState(false);
    const [errors, setErrors] = useState({});

    useEffect(() => {
        if (passwordSelected) {
            setURL(passwordSelected.url);
            setName(passwordSelected.name);
            setUsername(passwordSelected.username);
            setPassword(passwordSelected.password);
        }
    }, [passwordSelected])

    function handleSubmit(): void {
        const validation = handleValidate({
            url,
            name,
            username,
            password
        });

        setErrors(validation)
        if (Object.keys(validation).length > 0) {
            return;
        }

        const pass = {
            url,
            name,
            username,
            password
        }
        if (flow === "create") {
            api.post("/password-cards", {...pass, id: uuid()}).then(resp => {
                handleUpdatePasswordCards(resp.data)
                onClose()
            }).catch(err => {
                console.error(err)
                alert("An error has occurred");
            })
        } else {
            api.put(`/password-cards/${passwordSelected?.id}`, pass).then(resp => {
                handleUpdatePasswordCards(resp.data)
                onClose()
            }).catch(err => {
                console.error(err)
                alert("An error has occurred");
            })
        }
    }

    return (
        <Modal
            isOpen={isOpen}
            onClose={onClose} >
            <Container>
                <h3>{flow === "create" ? "New" : "Update"} Password</h3>

                <InputGroup>
                    <label>URL</label>
                    <Input
                        placeholder="https://example.com/login"
                        value={url}
                        onChange={setURL}
                        error={errors?.url} />
                </InputGroup>
                
                <InputGroup>
                    <label>Name</label>
                    <Input
                        placeholder="Example"
                        value={name}
                        onChange={setName}
                        error={errors?.name} />
                </InputGroup>

                <InputGroup>
                    <label>Username</label>
                    <Input
                        placeholder="person@mail.io"
                        value={username}
                        onChange={setUsername}
                        error={errors?.username} />
                </InputGroup>

                <InputGroup>
                    <label>Password</label>
                    <Input
                        type={showPassword ? "text" : "password"}
                        placeholder="**************"
                        value={password}
                        onChange={setPassword}
                        error={errors?.password}
                        renderRight={
                            <PasswordIconContainer>
                                <button onClick={() => {
                                    navigator.clipboard.writeText(password);
                                    alert("Copied the password!");
                                }}>
                                    <ContentCopyIcon />
                                </button>
                                <button onClick={() => setShowPassword(!showPassword)}>
                                    {showPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
                                </button>
                            </PasswordIconContainer>
                        } />
                </InputGroup>

                <Button
                    label="Submit"
                    onClick={handleSubmit} />
            </Container>
        </Modal>
    );
}