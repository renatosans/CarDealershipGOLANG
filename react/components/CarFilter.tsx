import { useState } from 'react'
import styles from '../styles/CarFilter.module.css'


const CarFilter = () => {
  return (
    <div className={styles.container} >
      <form className={styles.form} >
        <label htmlFor="fromYear" className={styles.label} >Ano (De)</label>
        <input type="text" className={styles.input} name="fromYear" />
        <label htmlFor="toYear" className={styles.label} >Ano (Até)</label>
        <input type="text" className={styles.input} name="toYear" />
        <br/>
        <label htmlFor="fromMileage" className={styles.label} >Quilometragem (De)</label>
        <input type="text" className={styles.input} name="fromMileage" />
        <label htmlFor="toMileage" className={styles.label} >Quilometragem (Até)</label>
        <input type="text" className={styles.input} name="toMileage" />
        <br/>
        <label htmlFor="fromPrice" className={styles.label} >Preço (De)</label>
        <input type="text" className={styles.input} name="fromPrice" />
        <label htmlFor="toPrice" className={styles.label} >Preço (Até)</label>
        <input type="text" className={styles.input} name="toPrice" />
      </form>
    </div>
  )
}

export default CarFilter
