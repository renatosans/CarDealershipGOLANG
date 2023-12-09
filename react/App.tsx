import "./App.css";
import Draggable from 'react-draggable'
import { carType } from './utils/types'
import { baseUrl } from './utils/defines'   // RUST API
import { useState, useEffect } from 'react'
import Carousel from './components/Carousel'
import VehicleList from './components/VehicleList'
import VehicleForm from './components/VehicleForm'
import CustomerForm from './components/CustomerForm'
import toast, { Toaster } from "react-hot-toast"


export default function Home() {
  const [cars, setCars] = useState<carType[]>();

  const [form1Open, setForm1Open] = useState<boolean>(false);
  const [form2Open, setForm2Open] = useState<boolean>(false);

  useEffect(() => {
    getCars();
  }, []);

  const getCars = () => {
    fetch(baseUrl + `/api/cars`)
    .then(resp => resp.json())
    .then(resultSet => setCars(resultSet))
    .catch(error => console.error(error))
  }

  const addCar = () => {
    setForm1Open(true);
  }

  const addCustomer = () => {
    setForm2Open(true);
  }

  return (
    <div className="container">
      <header>
        <title>Car Dealership</title>
        <meta name="description" content="Ride your own car" />
        <link rel="icon" href="/favicon.ico" />
      </header>
      <Toaster/>
      <Carousel/>
      <div className="actions">
        <button className="button" onClick={addCar}>Cadastrar veículo</button>        
        <button className="button" onClick={addCustomer}>Cadastrar Cliente</button>
      </div>
      <VehicleList items={cars} desc={'Car for sale. Available'} />
      <Draggable>
        <div style={{height: 0}}>
          <VehicleForm parentRef={{ setForm1Open, getCars }} opened={form1Open} />
        </div>
      </Draggable>
      <Draggable>
        <div style={{height: 0}}>
          <CustomerForm parentRef={{ setForm2Open }} opened={form2Open} />
        </div>
      </Draggable>
    </div>
  )
}
