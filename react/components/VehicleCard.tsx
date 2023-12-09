import { carType } from '../utils/types'
import styles from '../styles/VehicleCard.module.css'


type props = {
    car: carType;
    currency: string;  
}

const VehicleCard = ({car, currency}: props) => {
    const getInfo = (car: carType) => {
        return `${car.brand} ${car.model} ${car.year}`;
    }

    return (
      <div className={styles.container}>
        <img src={car.img} alt={getInfo(car)} width={200} height={150} />
        <h1 className={styles.title}>{getInfo(car)}</h1>
        <span className={styles.price}>{currency} {car.price} </span>
        <p className={styles.desc}>{car.color}</p>
        <p className={styles.desc}>{car.category}</p>
      </div>
    )
}

export default VehicleCard
