import { useState } from 'react'
import styles from '../styles/Navbar.module.css'


const Navbar = () => {
  return (
    <div className={styles.container}>
      <div className={styles.item}>
        <div className={styles.callButton}>
          <img src="/img/phone.png" alt="" width="32" height="32" />
        </div>
        <div className={styles.texts}>
          <div className={styles.text}>CONTACT US</div>
          <div className={styles.text}>012 345 678</div>
        </div>
      </div>
      <div className={styles.item}>
        <ul className={styles.list}>
          <li className={styles.listItem}>Home</li>
          <li className={styles.listItem}>Cars for sale</li>
          <img src="/img/logo.png" alt="" width={350} height={105} />
          <li className={styles.listItem}>Events</li>
          <li className={styles.listItem}>Media</li>
        </ul>
      </div>
      <div className={styles.item}>
        <div className={styles.cart}>
          <img src="/img/cart.png" alt="" width={30} height={30} />
          <div className={styles.counter}>2</div>
        </div>
      </div>
    </div>
  )
}


export default Navbar
