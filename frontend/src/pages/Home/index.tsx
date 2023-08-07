import { useEffect, useMemo, useState } from "react";
import { IPassword } from "../../types/password";
import api from "../../api";
import { Header } from "../../components/Header";
import { CardsWrapper, FABButton } from "./styled";
import { EmptyState } from "../../components/EmptyState";
import { Card } from "../../components/Card";
import { PasswordFormModal } from "../../components/PasswordFormModal";
import AddIcon from '@mui/icons-material/Add';

export const Home = (): JSX.Element => {
    const [openModal, setOpenModal] = useState(false);
  const [searchedValue, setSearchedValue] = useState("");
  const [passwords, setPasswords] = useState<IPassword[]>([]);
  const [passwordSelected, setPasswordSelected] = useState<IPassword | null>(null);

  useEffect(() => {
    api.get("/password-cards").then(resp => {
      setPasswords(resp.data ?? []);
    });
  }, [])

  function handleUpdatePasswordCards(value: IPassword) {
    const index = passwords.findIndex(pass => pass.id === value.id);

    if (index > -1) {
      passwords[index] = value;
      setPasswords([...passwords]);
    } else {
      setPasswords([...passwords, value]);
    }
  }

  function handleDeletePassword(id: string) {
    api.delete(`/password-cards/${id}`).then(() => {
      const index = passwords.findIndex(pass => pass.id === id);
      passwords.splice(index, 1);
      setPasswords([...passwords]);
    }).catch(err => {
      console.error(err)
      alert("An error has occurred");
    });
  }

  const filteredPassword = useMemo(() => {
    return passwords?.filter(item => item?.name?.toUpperCase()?.includes(searchedValue?.toLocaleUpperCase()))
  }, [passwords, searchedValue])

  return (
    <>
      <Header
        searchedValue={searchedValue}
        changeSearchedValue={setSearchedValue} />

      <CardsWrapper>
        {filteredPassword.length === 0 ? 
          <EmptyState message={
            passwords.length === 0 ? "No Passwords registered" : "No Passwords found with this name"
          } /> :
          filteredPassword.map((password) => (
            <Card
              key={password.id}
              password={password}
              onClick={() => {
                setPasswordSelected(password);
                setOpenModal(true);
              }}
              onDelete={() => handleDeletePassword(password.id)} />
          ))
        }
      </CardsWrapper>
      
      <FABButton onClick={() => {
        setPasswordSelected(null);
        setOpenModal(true);
      }}>
          <AddIcon />
      </FABButton>
      
      {openModal &&
        <PasswordFormModal
          isOpen={openModal}
          onClose={() => setOpenModal(false)}
          handleUpdatePasswordCards={handleUpdatePasswordCards}
          passwordSelected={passwordSelected}
          flow={passwordSelected ? "edit" : "create"} />}
    </>
  )
}
