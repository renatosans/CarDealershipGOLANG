
export type carType = {
    id?: number;
    brand: string;
    model: string;
    year: number;
    img: string;
    color: string;
    mileage: number;
    category: string;
    price: number;
}

export type customerType = {
    id?: number;
    first_name: string;
    last_name: string;
    birth_date: string;
    email: string;
    phone: string;
}

export type salespersonType = {
    id?: number;
    first_name: string;
    last_name: string;
    commission: number;
}
