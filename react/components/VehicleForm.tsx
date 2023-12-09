import { carType } from '../utils/types'
import { baseUrl } from '../utils/defines'
import { notification } from '../utils/notification'
import { useState, FormEvent } from 'react'
import toast, { Toaster, ToastOptions } from 'react-hot-toast'
import styles from '../styles/VehicleForm.module.css'


const emptyCar: carType = {
    id: 0,
    brand: "",
    model: "",
    year:  2000,
    img: "",
    color: "Prata",
    mileage: 0,
    category: "",
    price: 0,
}

const emptyImage = {
    image_format: "",
    image_data: "",
}

export const VehicleForm = ({parentRef, opened}: any) => {
	const [car, setCar] = useState<carType>(emptyCar);
    const [image, setImage] = useState(emptyImage);

    const handleSubmit = async (e: FormEvent) => {
		e.preventDefault();

		if (car.brand === "" || car.model === "" || car.price === 0) {
			toast.error('Favor preencher todos os campos!', notification.options as ToastOptions);
			return;
		}

        const payload = {...car, ...image};

        fetch(baseUrl + `/api/cars`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json', },
            body: JSON.stringify(payload), })
        .then((response) => {
            if(response.ok)
            {
                toast.success('Registro salvo com sucesso', notification.options as ToastOptions);
                setCar(emptyCar);
                setImage(emptyImage);
                parentRef.getCars();  // faz o refresh da lista de veículos
                parentRef.setForm1Open(false);   // Fecha o formulario
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
        if (e.target.type === 'number') {
            let value: number = parseFloat(e.target.value);
            if (!isNaN(value)) {
                value = parseFloat((Math.round(value * 100) / 100).toFixed(2)); // fixing 2 decimal places
                setCar({...car, [e.target.name]: value, });
                return;
            }
        }

		if (e.target.type === 'file') {
			const file = e.target.files[0];
			// Reads the file using the FileReader API
			const reader = new FileReader();
			reader.onloadend = () => {
                const result = reader.result as string;
				const fileData = result.split(';base64,');
				let formato = fileData[0].replace('data:', '') + ';base64';
                setImage({image_format: formato, image_data: fileData[1]});
			}
			reader.readAsDataURL(file);
		}

        setCar({...car, [e.target.name]: e.target.value, });
	}

    const getVisibility = (open: boolean) => {
        const visibility = open ? `visible` : `hidden`;
        return visibility;
    }

    return (
        <div className={styles.container} style={ {visibility: getVisibility(opened)} }>
            <Toaster />
            <form onSubmit={handleSubmit} className={styles.form} >
                <label htmlFor="brand" className={styles.label} >Marca</label>
                <input type="text" className={styles.input} name="brand" value={car.brand} onChange={onChange} />
                <label htmlFor="model" className={styles.label} >Modelo</label>
                <input type="text" className={styles.input} name="model" value={car.model} onChange={onChange} />
                <label htmlFor="category" className={styles.label} >Categoria</label>
                <select className={styles.input} name="category" onChange={onChange}>
                    <option value={"-"}> -Selecione- </option>
                    <option value={"SUV"}>SUV</option>
                    <option value={"hatchback"}>hatchback</option>
                    <option value={"sedan"}>sedan</option>
                </select>
                <label htmlFor="price" className={styles.label} >Preço</label>
                <input type="number" className={styles.input} name="price" value={car.price} onChange={onChange} />

                <label htmlFor="img" className={styles.label} >Foto</label>
                <div className={styles.fileDialog}>
                    <input type="file" name="img" accept=".gif,.jpg,.jpeg,.png" onChange={onChange} />
                    <img style={{'width':'100%'}} src={"data:" + image.image_format + ", " + image.image_data} alt={"car photo"}></img>
                </div>
                <button type="submit" className={styles.button}>Salvar</button>
            </form>
        </div>
    )
}

export default VehicleForm
