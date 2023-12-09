import { customerType } from '../utils/types'
import { baseUrl } from '../utils/defines'
import { notification } from '../utils/notification'
import { useState, FormEvent } from 'react'
import toast, { Toaster, ToastOptions } from 'react-hot-toast'
import styles from '../styles/VehicleForm.module.css'


const emptyCustomer: customerType = {
    id: 0,
    first_name: "",
    last_name: "",
    birth_date: "1983-04-04",      // TODO: fix date
    email: "",
    phone: "",
}

export const CustomerForm = ({parentRef, opened}: any) => {
	const [customer, setCustomer] = useState<customerType>(emptyCustomer);

    const handleSubmit = async (e: FormEvent) => {
		e.preventDefault();

		if (customer.first_name === "" || customer.last_name === "" || customer.email === "") {
			toast.error('Favor preencher todos os campos!', notification.options as ToastOptions);
			return;
		}

        const payload = {...customer};

        fetch(baseUrl + `/api/customers`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json', },
            body: JSON.stringify(payload), })
        .then((response) => {
            if(response.ok)
            {
                toast.success('Registro salvo com sucesso', notification.options as ToastOptions);
                setCustomer(emptyCustomer);
                parentRef.setForm2Open(false);   // Fecha o formulario
                return;
            }

            throw new Error(response.statusText);
        })  
        .then((text) => {
            console.log(text);
        })  
        .catch((error: any) => {
            toast.error(error.message, notification.options as ToastOptions);
        })
    }

    const onClose = () => {
        // not implemented yet
    }

	const onChange = (e: any) => {
        setCustomer({...customer, [e.target.name]: e.target.value, });
	}

    const getVisibility = (open: boolean) => {
        const visibility = open ? `visible` : `hidden`;
        return visibility;
    }

    return (
        <div className={styles.container} style={ {visibility: getVisibility(opened)} }>
            <Toaster />
            <form onSubmit={handleSubmit} className={styles.form} >
                <label htmlFor="first_name" className={styles.label} >Nome</label>
                <input type="text" className={styles.input} name="first_name" value={customer.first_name} onChange={onChange} />
                <label htmlFor="last_name" className={styles.label} >Sobrenome</label>
                <input type="text" className={styles.input} name="last_name" value={customer.last_name} onChange={onChange} />
                <label htmlFor="birth_date" className={styles.label} >Data Nascimento</label>
                <input type="date" className={styles.input} name="birth_date" />
                <label htmlFor="email" className={styles.label} >Email</label>
                <input type="text" className={styles.input} name="email" value={customer.email} onChange={onChange} />
                <label htmlFor="phone" className={styles.label} >Número de telefone</label>
                <input type="text" className={styles.input} name="phone" value={customer.phone} onChange={onChange} />

                <button type="submit" className={styles.button}>Salvar</button>
            </form>
        </div>
    )
}

export default CustomerForm
