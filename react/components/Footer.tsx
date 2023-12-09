import { useState } from 'react'
import styles from "../styles/Footer.module.css";


const Footer = () => {
    return (
        <div className={styles.container}>

          <div className={styles.item}>
              <img src="/img/trip.jpg" alt='' width="320px" height="300px" />
          </div>

          <div className={styles.item}>
          <div className={styles.card}>
          <h2 className={styles.motto}>
              Ride your own car
          </h2>
          </div>
          
          <div className={styles.card}>
              <h1 className={styles.title}>Our retailers</h1>
              <p className={styles.text}>1654 R. Don Road Avenue #304.<br/>Town City, 8524<br/>(602) 867-1010</p>
              <p className={styles.text}>1023 R. W. Caroll st Jones #125.<br/>Town City, 8524<br/>(602) 867-1018</p>
          </div>
          <div className={styles.card}>
          <h1 className={styles.title}>Working Hours</h1>
          <p className={styles.text}> Saturday - Sunday <br/>12:00 - 24:00</p>
          </div>
          </div> 
        </div>
    )
}

export default Footer
