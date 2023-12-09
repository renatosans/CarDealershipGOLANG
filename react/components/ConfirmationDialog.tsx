import { useState } from 'react';
import { Dialog, Group, Text, Button } from '@mantine/core';


type props = {
    message: string;
    handleResult: (result: boolean) => void;
}

export default function ConfirmationDialog({message, handleResult }: React.PropsWithChildren<props>) {
    const [open, setOpen] = useState(true);

    const confirm = () => {
        setOpen(false);
        handleResult(true);
    }

    const abort = () => {
        setOpen(false);
        handleResult(false);
    }

    return <>
        <Dialog opened={open} >
            <Text fz="xl">Confirmação</Text>
            <Text fz="md">{message}</Text>
            <br/>
            <Group align="flex-end">
                <Button variant="outline" onClick={confirm} >Sim</Button>
                <Button variant="outline" onClick={abort}   >Não</Button>
            </Group>
        </Dialog>
	</>
}
