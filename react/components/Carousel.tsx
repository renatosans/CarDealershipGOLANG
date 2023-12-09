import { useState } from 'react'
import styles from '../styles/Carousel.module.css'


const Carousel = () => {
    const [index,setIndex]= useState(0)

    const images = [
        "/img/banner-1.png",
        "/img/banner-2.png",
        "/img/banner-3.png"
    ]

    const handleArrow = (direction: string) => { // direction as parameter  < left | or | right >
        if(direction==="l"){
            setIndex(index !== 0 ? index-1 : 2) // decrease until index is 2
        }
        if(direction==="r"){
            setIndex(index !== 2 ? index+1 : 0) // increase until index is 0
        }
    };

    return (
      <div className={styles.container}>
        <div className={styles.wrapper} style={{transform:`translateX(${-100*index}vw)`}}>{
          images.map((img, i)=>(
            <div className={styles.imgContainer}  key={i} >
              <img src={img} alt="" height="100%" />
            </div>))
        }
        </div>

        <div className={styles.arrowContainer}  style={{left:0}} onClick={() => handleArrow("l")}>
           <img src={"/img/arrowl.png"} alt="" width="64" height="64" />
        </div>
        <div className={styles.arrowContainer}  style={{right:0}} onClick={() => handleArrow("r")}>
           <img src={"/img/arrowr.png"} alt="" width="64" height="64" />
        </div>
      </div>
    )
}

export default Carousel
