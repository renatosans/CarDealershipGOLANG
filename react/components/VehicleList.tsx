import CarFilter from './CarFilter';
import VehicleCard from './VehicleCard'
import { carType } from '../utils/types'
import styles from '../styles/VehicleList.module.css'


type props = {
  items: carType[] | undefined;
  desc: string;
}

const VehicleList = ({items, desc}: props) => {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Best Deal Guarantee</h1>
      <p className={styles.desc}>{desc}</p>
      <div className={styles.mainWrapper}>
        <CarFilter/>
        <div className={styles.listWrapper}>{
          (items)
          ? items.map((car: carType) => <VehicleCard car={car} currency={"R$"} />)
          : <p>No item to display</p>
        }
        </div>
      </div>
    </div>
  )
}

export default VehicleList
